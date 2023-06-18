package main

import (
	"json_api_project/api"
	rendering "json_api_project/cmd_ui"
)

func main() {
	conf := api.ServerConfig{
		Sockaddr: ":6666",
	}
	go rendering.Spinner()
	server := api.NewServer(conf)
	server.Run()
}
