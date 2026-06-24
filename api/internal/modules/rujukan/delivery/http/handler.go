package http

import (
	"strconv"

	"rujukan/internal/modules/rujukan/domain"
	"rujukan/internal/modules/rujukan/repository"
	"rujukan/internal/pkg/response"

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
		response.InternalError(c, err.Error())
		return
	}

	response.OKList(c, data, int(total), limit, offset)
}

func (h *RujukanHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "ID tidak valid")
		return
	}

	data, err := h.repo.GetByID(id)
	if err != nil {
		response.NotFound(c, "Rujukan tidak ditemukan")
		return
	}

	response.OK(c, data)
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
		response.ValidationError(c, err)
		return
	}

	rujukan := domain.Rujukan{
		IDPasien:       input.IDPasien,
		IDFaskesAsal:   input.IDFaskesAsal,
		IDFaskesTujuan: input.IDFaskesTujuan,
		KodeICD:        input.KodeICD,
		Diagnosa:       input.Diagnosa,
		Catatan:        input.Catatan,
		IDUser:         extractUserID(c),
	}

	if err := h.repo.Create(&rujukan); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, rujukan)
}

type UpdateStatusInput struct {
	Status string `json:"status" binding:"required,oneof=menunggu diterima ditolak selesai"`
}

func (h *RujukanHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "ID tidak valid")
		return
	}

	var input UpdateStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := h.repo.UpdateStatus(id, input.Status); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKMessage(c, "Status rujukan berhasil diperbarui")
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
