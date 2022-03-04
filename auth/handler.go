package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"PLAYLISTBOOK/utils"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return Handler{
		service: s,
	}
}

func (h Handler) handleRegister(c *gin.Context) {
	ctx := context.Background()
	var params RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := h.service.Register(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}
