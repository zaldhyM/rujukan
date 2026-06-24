package http

import (
	"strconv"

	"rujukan/internal/modules/pasien/domain"
	"rujukan/internal/modules/pasien/repository"
	"rujukan/internal/pkg/response"

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
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	data, total, err := h.repo.QueryAll(search, limit, offset)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.OKList(c, data, int(total), limit, offset)
}

func (h *PasienHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.repo.GetByID(id)
	if err != nil {
		response.NotFound(c, "Pasien tidak ditemukan")
		return
	}

	response.OK(c, data)
}

func (h *PasienHandler) Create(c *gin.Context) {
	var pasien domain.Pasien
	if err := c.ShouldBindJSON(&pasien); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := h.repo.Create(&pasien); err != nil {
		response.ServerError(c, err)
		return
	}

	response.Created(c, pasien)
}

func (h *PasienHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var pasien domain.Pasien
	if err := c.ShouldBindJSON(&pasien); err != nil {
		response.ValidationError(c, err)
		return
	}

	pasien.ID = id
	if err := h.repo.Update(&pasien); err != nil {
		response.ServerError(c, err)
		return
	}

	response.OK(c, pasien)
}

func (h *PasienHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		response.ServerError(c, err)
		return
	}

	response.OKMessage(c, "Pasien berhasil dihapus")
}
