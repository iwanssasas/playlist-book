package server

import (
	"PLAYLISTBOOK/app"
	"PLAYLISTBOOK/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	router := gin.Default()

	appService := app.NewService()
	app.NewRouter(router, appService)

	conf := config.Get()
	host := fmt.Sprintf(":%v", conf.Port)
	router.Run(host)
}
