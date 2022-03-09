package app

import (
	"PLAYLISTBOOK/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, service Service) {
	handler := NewHandler(service)

	auth := router.Group("/", utils.TokenVerify())
	{
		auth.GET("/test", handler.handleTest)
		auth.POST("/books", utils.Authorization([]string{"admin"}), handler.handleSubmitBook)
		auth.GET("/books", handler.handleSelectAllBook)
		auth.PUT("/books/:id", handler.handleUpdateBookById)
		auth.DELETE("/books/:id", handler.handleDeleteBookById)

	}
}
