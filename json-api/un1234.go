package main

import (
	"json-api/Storage"
	"json-api/api"
	"log"
)

func main() {

	log.Printf("start process")
	pglayer, _ := Storage.NewPostgreLayer()

	conf := api.ServerConfig{
		Sockaddr: "127.0.0.1:6666",
	}

	server := api.NewServer(conf, pglayer)
	server.Run()

}
