package app

import "github.com/gin-gonic/gin"

func NewRouter(router *gin.Engine, service Service) {
	handler := NewHandler(service)
	router.GET("/", handler.handleTest)
}
