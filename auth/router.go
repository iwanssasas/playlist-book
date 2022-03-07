package auth

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, service Service) {
	handler := NewHandler(service)
	router.POST("/register", handler.handleRegister)
	router.POST("/login", handler.handleLogin)
	router.GET("/ping", handler.handlePing)
}
