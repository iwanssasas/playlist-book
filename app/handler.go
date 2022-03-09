package app

import (
	"PLAYLISTBOOK/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return Handler{
		service: s,
	}
}

func (h Handler) handleTest(c *gin.Context) {
	name := c.GetString("username")
	c.JSON(http.StatusOK, name)
}

func (h Handler) handleSubmitBook(c *gin.Context) {
	ctx := context.Background()
	var params BookParams
	name := c.GetString("username")
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	data, err := h.service.SubmitBook(ctx, params, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}

func (h Handler) handleSelectAllBook(c *gin.Context) {
	ctx := context.Background()
	response, err := h.service.selectAllBook(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(response))
}

func (h Handler) handleDeleteBookById(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteBookById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response("success"))
}

func (h Handler) handleUpdateBookById(c *gin.Context) {
	ctx := context.Background()
	var params BookParams
	id := c.Param("id")
	name := c.GetString("username")
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	data, err := h.service.UpdateBookById(ctx, params, name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.Response(data))
}
