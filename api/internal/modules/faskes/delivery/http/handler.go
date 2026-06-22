package http

import (
	"net/http"
	"strconv"
	"rujukan/internal/modules/faskes/domain"
	"rujukan/internal/modules/faskes/repository"

	"github.com/gin-gonic/gin"
)

type FaskesHandler struct {
	repo repository.FaskesRepository
}

func NewFaskesHandler(repo repository.FaskesRepository) *FaskesHandler {
	return &FaskesHandler{repo: repo}
}

func (h *FaskesHandler) QueryAll(c *gin.Context) {
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

func (h *FaskesHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format",
		})
		return
	}

	data, err := h.repo.GetByID(int16(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Faskes not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *FaskesHandler) Create(c *gin.Context) {
	var faskes domain.Faskes
	if err := c.ShouldBindJSON(&faskes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if err := h.repo.Create(&faskes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    faskes,
	})
}

func (h *FaskesHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format",
		})
		return
	}

	var faskes domain.Faskes
	if err := c.ShouldBindJSON(&faskes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	faskes.ID = int16(id)
	if err := h.repo.Update(&faskes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    faskes,
	})
}

func (h *FaskesHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format",
		})
		return
	}

	if err := h.repo.Delete(int16(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Faskes deleted successfully",
	})
}
