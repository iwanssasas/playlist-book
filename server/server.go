package server

import (
	"PLAYLISTBOOK/app"
	"PLAYLISTBOOK/auth"
	"PLAYLISTBOOK/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	router := gin.Default()

	appService := app.NewService()
	authService := auth.NewService()

	app.NewRouter(router, appService)
	auth.NewRouter(router, authService)

	conf := config.Get()
	host := fmt.Sprintf(":%v", conf.Port)
	router.Run(host)
}
