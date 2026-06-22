package http

import (
	"net/http"
	"strconv"
	"rujukan/internal/modules/wilayah/repository"

	"github.com/gin-gonic/gin"
)

type WilayahHandler struct {
	repo repository.WilayahRepository
}

func NewWilayahHandler(repo repository.WilayahRepository) *WilayahHandler {
	return &WilayahHandler{repo: repo}
}

func (h *WilayahHandler) QueryWilayah(c *gin.Context) {
	search := c.Query("search")
	parentID := c.Query("parent_id")
	jenisStr := c.Query("jenis")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	var jenis int8 = 0
	if jenisStr != "" {
		j, err := strconv.Atoi(jenisStr)
		if err == nil {
			jenis = int8(j)
		}
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	data, total, err := h.repo.QueryWilayah(search, jenis, parentID, limit, offset)
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

func (h *WilayahHandler) GetWilayahByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.repo.GetWilayahByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Wilayah not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *WilayahHandler) QueryJenisWilayah(c *gin.Context) {
	data, err := h.repo.QueryJenisWilayah()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *WilayahHandler) GetJenisWilayahByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format",
		})
		return
	}

	data, err := h.repo.GetJenisWilayahByID(int8(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Jenis wilayah not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
