package http

import (
	"net/http"
	"rujukan/internal/modules/user/repository"

	"github.com/gin-gonic/gin"
)

// UserHandler handles HTTP requests related to the user domain.
type UserHandler struct {
	repo repository.UserRepository
}

// NewUserHandler returns a new instance of UserHandler.
func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
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
