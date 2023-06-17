package api

import (
	"fmt"
)

func DoSomeStuff() {
	fmt.Println("some stuff")
}

type ServerConfig struct {
	Sockaddr string
}

type Server struct {
}
