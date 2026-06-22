package http

import (
	"net/http"
	"strconv"
	"rujukan/internal/modules/referensi/repository"

	"github.com/gin-gonic/gin"
)

type ReferensiHandler struct {
	repo repository.ReferensiRepository
}

func NewReferensiHandler(repo repository.ReferensiRepository) *ReferensiHandler {
	return &ReferensiHandler{repo: repo}
}

func (h *ReferensiHandler) QueryReferensi(c *gin.Context) {
	jenisStr := c.Query("jenis")
	search := c.Query("search")

	var jenis int8 = 0
	if jenisStr != "" {
		j, err := strconv.Atoi(jenisStr)
		if err == nil {
			jenis = int8(j)
		}
	}

	data, err := h.repo.QueryReferensi(jenis, search)
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

func (h *ReferensiHandler) GetReferensiByKeys(c *gin.Context) {
	jenisStr := c.Param("jenis")
	idStr := c.Param("id")

	jenis, err := strconv.Atoi(jenisStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid jenis format",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format",
		})
		return
	}

	data, err := h.repo.GetReferensiByKeys(int8(jenis), int16(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Referensi not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func (h *ReferensiHandler) QueryJenisReferensi(c *gin.Context) {
	data, err := h.repo.QueryJenisReferensi()
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

func (h *ReferensiHandler) GetJenisReferensiByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format",
		})
		return
	}

	data, err := h.repo.GetJenisReferensiByID(int8(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Jenis referensi not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}
