package http

import (
	"net/http"
	"strconv"
	"rujukan/internal/modules/pasien/domain"
	"rujukan/internal/modules/pasien/repository"

	"github.com/gin-gonic/gin"
)

type PasienHandler struct {
	repo repository.PasienRepository
}

func NewPasienHandler(repo repository.PasienRepository) *PasienHandler {
	return &PasienHandler{repo: repo}
}

func (h *PasienHandler) QueryAll(c *gin.Context) {
	search := c.Query("search")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	data, total, err := h.repo.QueryAll(search, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"total":   total,
		"limit":   limit,
		"offset":  offset,
		"data":    data,
	})
}

func (h *PasienHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Pasien not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *PasienHandler) Create(c *gin.Context) {
	var pasien domain.Pasien
	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := h.repo.Create(&pasien); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    pasien,
	})
}

func (h *PasienHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var pasien domain.Pasien
	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	pasien.ID = id
	if err := h.repo.Update(&pasien); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pasien,
	})
}

func (h *PasienHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Pasien deleted successfully",
	})
}
