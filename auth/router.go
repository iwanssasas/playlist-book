package auth

import (
	"PLAYLISTBOOK/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, service Service) {
	handler := NewHandler(service)
	router.POST("/register", handler.handleRegister)
	router.POST("/login", handler.handleLogin)

	auth := router.Group("/", utils.TokenVerify())
	{
		auth.GET("/ping", utils.Authorization([]string{"guest", "admin"}), handler.handlePing)
		auth.GET("/admin", utils.Authorization([]string{"admin"}), handler.handleAdmin)
	}
}
