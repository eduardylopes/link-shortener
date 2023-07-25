package link

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateLink(ctx *gin.Context) {
	r := CreateLinkReq{
		URL: ctx.Query("url"),
	}

	if err := r.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.Service.CreateLink(ctx, &r)

	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetLinkByCode(ctx *gin.Context) {
	r := GetLinkByCodeReq{
		Code: ctx.Param("code"),
	}

	if err := r.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.Service.GetLinkByCode(ctx, r.Code)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, res.URL)
}
