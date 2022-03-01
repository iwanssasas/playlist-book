package main

import (
	"PLAYLISTBOOK/config"
	"fmt"
)

func main() {

	config.Init()
	fmt.Println(config.Get())

}
