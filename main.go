package main

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/server"
)

func main() {
	config.Init()
	server.NewServer()
}

//created_at
// created_by
// updated_at
// updated_by
