package http

import (
	"strconv"

	"rujukan/internal/modules/faskes/domain"
	"rujukan/internal/modules/faskes/repository"
	"rujukan/internal/pkg/response"

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
		response.InternalError(c, err.Error())
		return
	}

	response.OKList(c, data, int(total), limit, offset)
}

func (h *FaskesHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	data, err := h.repo.GetByID(int16(id))
	if err != nil {
		response.NotFound(c, "Faskes tidak ditemukan")
		return
	}

	response.OK(c, data)
}

func (h *FaskesHandler) Create(c *gin.Context) {
	var faskes domain.Faskes
	if err := c.ShouldBindJSON(&faskes); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := h.repo.Create(&faskes); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, faskes)
}

func (h *FaskesHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	var faskes domain.Faskes
	if err := c.ShouldBindJSON(&faskes); err != nil {
		response.ValidationError(c, err)
		return
	}

	faskes.ID = int16(id)
	if err := h.repo.Update(&faskes); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, faskes)
}

func (h *FaskesHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Format ID tidak valid")
		return
	}

	if err := h.repo.Delete(int16(id)); err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.OKMessage(c, "Faskes berhasil dihapus")
}
