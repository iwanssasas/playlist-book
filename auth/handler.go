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
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	data, err := h.service.Register(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handleLogin(c *gin.Context) {
	ctx := context.Background()
	var params LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	data, err := h.service.Login(ctx, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handlePing(c *gin.Context) {
	// fmt.Println(c.Get("email"))
	ctx := context.Background()
	data, err := h.service.Ping(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handleAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, utils.Response("Admin Only!"))
}
