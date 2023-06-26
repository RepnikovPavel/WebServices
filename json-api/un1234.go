package main

import (
	"fmt"
	"json-api/Storage"
	"json-api/cmd_ui"
	"log"
)

func main() {

	log.Printf("start process")
	go cmd_ui.Spinner()
	// conf := api.ServerConfig{
	// 	Sockaddr: "127.0.0.1:6666",
	// }
	// server := api.NewServer(conf)
	// server.Run()
	storage_layer, err := Storage.NewPostgreLayer()
	if err != nil {
		cmd_ui.LOGERR(err)
	}
	fmt.Println(&storage_layer)

}
