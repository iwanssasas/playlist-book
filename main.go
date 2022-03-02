package main

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/server"
)

func main() {
	config.Init()
	server.NewServer()
}
