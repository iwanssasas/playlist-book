package app

import (
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
	c.JSON(http.StatusOK, "OK")
}
