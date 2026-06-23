package http

import (
	"net/http"
	"strconv"

	"rujukan/internal/modules/rujukan/domain"
	"rujukan/internal/modules/rujukan/repository"

	"github.com/gin-gonic/gin"
)

type RujukanHandler struct {
	repo repository.RujukanRepository
}

func NewRujukanHandler(repo repository.RujukanRepository) *RujukanHandler {
	return &RujukanHandler{repo: repo}
}

func (h *RujukanHandler) QueryAll(c *gin.Context) {
	search := c.Query("search")
	status := c.Query("status")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	data, total, err := h.repo.QueryAll(search, status, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
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

func (h *RujukanHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ID tidak valid"})
		return
	}

	data, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Rujukan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

type CreateInput struct {
	IDPasien       string  `json:"id_pasien" binding:"required"`
	IDFaskesAsal   string  `json:"id_faskes_asal" binding:"required"`
	IDFaskesTujuan string  `json:"id_faskes_tujuan" binding:"required"`
	KodeICD        *string `json:"kode_icd"`
	Diagnosa       *string `json:"diagnosa"`
	Catatan        *string `json:"catatan"`
}

func (h *RujukanHandler) Create(c *gin.Context) {
	var input CreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	userID := extractUserID(c)

	rujukan := domain.Rujukan{
		IDPasien:       input.IDPasien,
		IDFaskesAsal:   input.IDFaskesAsal,
		IDFaskesTujuan: input.IDFaskesTujuan,
		KodeICD:        input.KodeICD,
		Diagnosa:       input.Diagnosa,
		Catatan:        input.Catatan,
		IDUser:         userID,
	}

	if err := h.repo.Create(&rujukan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": rujukan})
}

type UpdateStatusInput struct {
	Status string `json:"status" binding:"required,oneof=menunggu diterima ditolak selesai"`
}

func (h *RujukanHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "ID tidak valid"})
		return
	}

	var input UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if err := h.repo.UpdateStatus(id, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Status rujukan berhasil diperbarui"})
}

func extractUserID(c *gin.Context) uint {
	v, _ := c.Get("userID")
	switch id := v.(type) {
	case uint:
		return id
	case uint16:
		return uint(id)
	case int:
		return uint(id)
	}
	return 0
}
