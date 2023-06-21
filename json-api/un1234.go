package main

import (
	"json-api/api"
	"json-api/cmd_ui"
	"log"
)

func main() {

	log.Printf("start process")
	conf := api.ServerConfig{
		Sockaddr: "127.0.0.1:6666",
	}
	go cmd_ui.Spinner()
	server := api.NewServer(conf)
	server.Run()
}
