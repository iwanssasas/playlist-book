package auth

import (
	"PLAYLISTBOOK/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, service Service) {
	handler := NewHandler(service)
	router.POST("register", handler.handleRegister)
	router.POST("login", handler.handleLogin)
	router.GET("auth/google", handler.handleOauthGoogle)
	router.GET("callback/google", handler.handleGoogleCallback)

	auth := router.Group("", utils.TokenVerify())
	{
		auth.GET("ping", utils.Authorization([]string{"user", "admin"}), handler.handlePing)
		auth.GET("admin", utils.Authorization([]string{"admin"}), handler.handleAdmin)
	}

	profile := router.Group("", utils.TokenVerify())
	{
		profile.GET("profile", utils.Authorization([]string{"user"}), handler.handleGetProfile)
		profile.POST("profile", utils.Authorization([]string{"user"}), handler.handleEditProfile)
		profile.DELETE("profile", utils.Authorization([]string{"user"}), handler.handleDeleteProfile)
	}
}
