package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"rujukan/internal/infrastructure/auth"
	"rujukan/internal/infrastructure/cache"
	"rujukan/internal/modules/user/domain"
	"rujukan/internal/modules/user/repository"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler handles HTTP requests related to the user domain.
type UserHandler struct {
	repo repository.UserRepository
	rdb  *redis.Client
}

// NewUserHandler returns a new instance of UserHandler.
func NewUserHandler(repo repository.UserRepository, rdb *redis.Client) *UserHandler {
	return &UserHandler{repo: repo, rdb: rdb}
}

// DataAll handles request to retrieve all user data.
func (h *UserHandler) DataAll(c *gin.Context) {
	respon, err := h.repo.QueryAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": len(respon) > 0,
		"total":   len(respon),
		"data":    respon,
	})
}

// LoginInput defines the request body format for authentication.
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login authenticates a user and returns a token & session cookie.
func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Username and password are required",
		})
		return
	}

	// 1. Fetch user by username
	user, err := h.repo.GetByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid username or password",
		})
		return
	}

	// 2. Verify status
	if user.STATUS != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Account is inactive",
		})
		return
	}

	// 3. Verify masa aktif (expiry)
	if user.MASA_AKTIF != "" {
		if expiryDate, err := time.Parse("2006-01-02", user.MASA_AKTIF); err == nil {
			if time.Now().After(expiryDate) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "Account has expired",
				})
				return
			}
		}
	}

	// 4. Verify password (bcrypt with plaintext fallback for legacy database credentials)
	passwordMatched := false
	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(input.Password))
	if err == nil {
		passwordMatched = true
	} else {
		// Fallback plaintext check (for migration)
		if user.PASSWORD == input.Password {
			passwordMatched = true
			// Proactively migrate/hash the plaintext password in DB to bcrypt!
			hashed, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
			if hashErr == nil {
				if updateErr := h.repo.UpdatePassword(user.ID, string(hashed)); updateErr != nil {
					log.Printf("⚠️ Failed to upgrade password for user %s: %v", user.USERNAME, updateErr)
				} else {
					log.Printf("🔒 Password successfully upgraded to bcrypt for user: %s", user.USERNAME)
				}
			}
		}
	}

	if !passwordMatched {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid username or password",
		})
		return
	}

	// 5. Generate JWT Token
	tokenDuration := tokenExpiredDuration()
	token, err := auth.GenerateToken(user.ID, user.USERNAME, user.ROLE, user.ID_FASKES, user.NAMA, tokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Could not generate authentication token",
		})
		return
	}

	// 6. Store token in Redis (overwrites previous session — single active session)
	if err := h.rdb.Set(context.Background(), cache.SessionKey(user.ID), token, tokenDuration).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Could not register session in cache",
		})
		return
	}

	// 7. Set HTTP-only Cookie for Session Authentication
	c.SetCookie("session_token", token, int(tokenDuration.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":        user.ID,
			"username":  user.USERNAME,
			"nama":      user.NAMA,
			"role":      user.ROLE,
			"id_faskes": user.ID_FASKES,
		},
	})
}

// RegisterInput defines the request body format for user registration.
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
	Nama     string `json:"nama" binding:"required"`
	Nik      string `json:"nik"`
	IDFaskes uint   `json:"id_faskes"`
	Role     string `json:"role" binding:"required,oneof=admin faskes dinkeskab dinkesprov"`
}

// Register creates a new user with hashed password.
func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Check if username already exists
	_, err := h.repo.GetByUsername(input.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Username is already registered",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error securing password",
		})
		return
	}

	newUser := domain.User{
		USERNAME:               input.Username,
		PASSWORD:               string(hashedPassword),
		NAMA:                   input.Nama,
		NIK:                    input.Nik,
		ID_FASKES:              input.IDFaskes,
		ROLE:                   input.Role,
		STATUS:                 1, // Active
		STATUS_TELEGRAM:        "0",
		TOKEN:                  "",
		TERAKHIR_UBAH_PASSWORD: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Set expiration date (e.g., active for 1 year by default)
	newUser.MASA_AKTIF = time.Now().AddDate(1, 0, 0).Format("2006-01-02")

	if err := h.repo.Create(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
		"data": gin.H{
			"id":        newUser.ID,
			"username":  newUser.USERNAME,
			"nama":      newUser.NAMA,
			"role":      newUser.ROLE,
			"id_faskes": newUser.ID_FASKES,
		},
	})
}

// Logout clears the user session token in DB and clears the session cookie.
func (h *UserHandler) Logout(c *gin.Context) {
	// Get user ID from middleware context
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	var userID uint
	switch v := userIDValue.(type) {
	case uint:
		userID = v
	case uint16:
		userID = uint(v)
	case int:
		userID = uint(v)
	}

	// Delete session from Redis
	if err := h.rdb.Del(context.Background(), cache.SessionKey(userID)).Err(); err != nil {
		log.Printf("⚠️ Failed to delete session from Redis on logout: %v", err)
	}

	// Clear the cookie (expire immediately)
	c.SetCookie("session_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logged out successfully",
	})
}

// Me retrieves current user profile from context.
func (h *UserHandler) Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	dbUser := user.(domain.User)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"id":        dbUser.ID,
			"username":  dbUser.USERNAME,
			"nama":      dbUser.NAMA,
			"nik":       dbUser.NIK,
			"role":      dbUser.ROLE,
			"id_faskes": dbUser.ID_FASKES,
		},
	})
}

// tokenExpiredDuration reads TOKEN_EXPIRED_HOURS from env, defaults to 24 hours.
func tokenExpiredDuration() time.Duration {
	hours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_HOURS"))
	if err != nil || hours <= 0 {
		hours = 24
	}
	return time.Duration(hours) * time.Hour
}
