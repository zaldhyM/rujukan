package http

import (
	"strconv"

	"rujukan/internal/modules/referensi/repository"
	"rujukan/internal/pkg/response"

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
		if j, err := strconv.Atoi(jenisStr); err == nil {
			jenis = int8(j)
		}
	}

	data, err := h.repo.QueryReferensi(jenis, search)
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.OK(c, data)
}

func (h *ReferensiHandler) GetReferensiByKeys(c *gin.Context) {
	jenis, err := strconv.Atoi(c.Param("jenis"))
	if err != nil {
		response.BadRequest(c, "Format jenis tidak valid")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	data, err := h.repo.GetReferensiByKeys(int8(jenis), int16(id))
	if err != nil {
		response.NotFound(c, "Referensi tidak ditemukan")
		return
	}

	response.OK(c, data)
}

func (h *ReferensiHandler) QueryJenisReferensi(c *gin.Context) {
	data, err := h.repo.QueryJenisReferensi()
	if err != nil {
		response.ServerError(c, err)
		return
	}

	response.OK(c, data)
}

func (h *ReferensiHandler) GetJenisReferensiByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	data, err := h.repo.GetJenisReferensiByID(int8(id))
	if err != nil {
		response.NotFound(c, "Jenis referensi tidak ditemukan")
		return
	}

	response.OK(c, data)
}
