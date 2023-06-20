package main

import (
	// "database/sql"
	"log"
	// "json_api_project/NodeStorage"
	rendering "json_api_project/cmd_ui"
)

func main() {
	log.Printf("start process")

	// conf := api.ServerConfig{
	// Sockaddr: ":6666",
	// }
	// go rendering.Spinner()
	// server := api.NewServer(conf)
	// server.Run()

	// NodeStorage.SQLiteConf{DriverName: "sqlite3",DataSourceName: ""}
	rendering.Spinner()
}
