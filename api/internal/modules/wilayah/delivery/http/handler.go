package http

import (
	"strconv"

	"rujukan/internal/modules/wilayah/repository"
	"rujukan/internal/pkg/response"

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

	var jenis int8 = 0
	if jenisStr != "" {
		if j, err := strconv.Atoi(jenisStr); err == nil {
			jenis = int8(j)
		}
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	data, total, err := h.repo.QueryWilayah(search, jenis, parentID, limit, offset)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKList(c, data, int(total), limit, offset)
}

func (h *WilayahHandler) GetWilayahByID(c *gin.Context) {
	id := c.Param("id")
	data, err := h.repo.GetWilayahByID(id)
	if err != nil {
		response.NotFound(c, "Wilayah tidak ditemukan")
		return
	}

	response.OK(c, data)
}

func (h *WilayahHandler) QueryJenisWilayah(c *gin.Context) {
	data, err := h.repo.QueryJenisWilayah()
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, data)
}

func (h *WilayahHandler) GetJenisWilayahByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	data, err := h.repo.GetJenisWilayahByID(int8(id))
	if err != nil {
		response.NotFound(c, "Jenis wilayah tidak ditemukan")
		return
	}

	response.OK(c, data)
}
