package http

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"rujukan/internal/infrastructure/auth"
	"rujukan/internal/infrastructure/cache"
	"rujukan/internal/modules/user/domain"
	"rujukan/internal/modules/user/repository"
	"rujukan/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo repository.UserRepository
	rdb  *redis.Client
}

func NewUserHandler(repo repository.UserRepository, rdb *redis.Client) *UserHandler {
	return &UserHandler{repo: repo, rdb: rdb}
}

func (h *UserHandler) DataAll(c *gin.Context) {
	respon, err := h.repo.QueryAll()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.OKList(c, respon, len(respon), len(respon), 0)
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, err)
		return
	}

	user, err := h.repo.GetByUsername(input.Username)
	if err != nil {
		response.Unauthorized(c, "Username atau password tidak valid")
		return
	}

	if user.STATUS != 1 {
		response.Unauthorized(c, "Akun tidak aktif")
		return
	}

	if user.MASA_AKTIF != "" {
		if expiryDate, err := time.Parse("2006-01-02", user.MASA_AKTIF); err == nil {
			if time.Now().After(expiryDate) {
				response.Unauthorized(c, "Akun telah kadaluarsa")
				return
			}
		}
	}

	passwordMatched := false
	err = bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(input.Password))
	if err == nil {
		passwordMatched = true
	} else if user.PASSWORD == input.Password {
		// Fallback plaintext — migrasi otomatis ke bcrypt
		passwordMatched = true
		if hashed, hashErr := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost); hashErr == nil {
			if updateErr := h.repo.UpdatePassword(user.ID, string(hashed)); updateErr != nil {
				log.Printf("⚠️ Gagal upgrade password user %s: %v", user.USERNAME, updateErr)
			} else {
				log.Printf("🔒 Password user %s berhasil diupgrade ke bcrypt", user.USERNAME)
			}
		}
	}

	if !passwordMatched {
		response.Unauthorized(c, "Username atau password tidak valid")
		return
	}

	tokenDuration := tokenExpiredDuration()
	token, err := auth.GenerateToken(user.ID, user.USERNAME, user.ROLE, user.ID_FASKES, user.NAMA, tokenDuration)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	if err := h.rdb.Set(context.Background(), cache.SessionKey(user.ID), token, tokenDuration).Err(); err != nil {
		response.ServerError(c, err)
		return
	}

	c.SetCookie("session_token", token, int(tokenDuration.Seconds()), "/", "", isProduction(), true)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Login berhasil",
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

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
	Nama     string `json:"nama" binding:"required"`
	Nik      string `json:"nik"`
	IDFaskes uint   `json:"id_faskes"`
	Role     string `json:"role" binding:"required,oneof=admin faskes dinkeskab dinkesprov"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, err)
		return
	}

	if _, err := h.repo.GetByUsername(input.Username); err == nil {
		response.Conflict(c, "Username sudah terdaftar")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	newUser := domain.User{
		USERNAME:               input.Username,
		PASSWORD:               string(hashedPassword),
		NAMA:                   input.Nama,
		NIK:                    input.Nik,
		ID_FASKES:              input.IDFaskes,
		ROLE:                   input.Role,
		STATUS:                 1,
		STATUS_TELEGRAM:        "0",
		TOKEN:                  "",
		TERAKHIR_UBAH_PASSWORD: time.Now().Format("2006-01-02 15:04:05"),
		MASA_AKTIF:             time.Now().AddDate(1, 0, 0).Format("2006-01-02"),
	}

	if err := h.repo.Create(&newUser); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Created(c, gin.H{
		"id":        newUser.ID,
		"username":  newUser.USERNAME,
		"nama":      newUser.NAMA,
		"role":      newUser.ROLE,
		"id_faskes": newUser.ID_FASKES,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "Unauthorized")
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

	if err := h.rdb.Del(context.Background(), cache.SessionKey(userID)).Err(); err != nil {
		log.Printf("⚠️ Gagal hapus sesi dari Redis saat logout: %v", err)
	}

	c.SetCookie("session_token", "", -1, "/", "", isProduction(), true)
	response.OKMessage(c, "Logout berhasil")
}

func (h *UserHandler) Me(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.Unauthorized(c, "Unauthorized")
		return
	}

	dbUser := user.(domain.User)

	response.OK(c, gin.H{
		"id":        dbUser.ID,
		"username":  dbUser.USERNAME,
		"nama":      dbUser.NAMA,
		"nik":       dbUser.NIK,
		"role":      dbUser.ROLE,
		"id_faskes": dbUser.ID_FASKES,
	})
}

func tokenExpiredDuration() time.Duration {
	hours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_HOURS"))
	if err != nil || hours <= 0 {
		hours = 24
	}
	return time.Duration(hours) * time.Hour
}

func isProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}
