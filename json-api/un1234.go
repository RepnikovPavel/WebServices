package main

import (
	"json_api_project/api"
	rendering "json_api_project/cmd_ui"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Printf("start process")
	conf := api.ServerConfig{
		Sockaddr: ":6666",
	}
	go rendering.Spinner()
	server := api.NewServer(conf)
	server.Run()
	rendering.Spinner()
}
