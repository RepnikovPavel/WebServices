package main

import (
	"fmt"
	"json_api_project/api"
	rendering "json_api_project/cmd_ui"
	"time"
)

func main() {
	conf := api.ServerConfig{
		Sockaddr: "1234",
	}
	fmt.Println(conf.Sockaddr)
	x := 1
	go rendering.Spinner()
	fmt.Print(x)
	api.DoSomeStuff()
	time.Sleep(10 * time.Second)
}
