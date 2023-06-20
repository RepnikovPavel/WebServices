```go
package main

import (
	"database/sql"
	"json_api_project/NodeStorage"
	rendering "json_api_project/cmd_ui"
	"log"
	"path"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Printf("start process")
	const (
		SQLiteDBFolder string = "/home/user/WebServices/json-api-DataBase/SQLite"
		SqliteDBName   string = "json_api.db"
	)
	DataSourceName := path.Join(SQLiteDBFolder, SqliteDBName)
	log.Println(DataSourceName)
	dbconf := NodeStorage.SQLiteConf{DriverName: "sqlite3", DataSourceName: DataSourceName}
	db, err := sql.Open(dbconf.DriverName, dbconf.DataSourceName)
	if err != nil {
		rendering.LOGERR(err)
	}
     
	SQL_Schema_AccountCold := `
	CREATE TABLE IF NOT EXISTS cold_account_info(
		id INT NOT NULL,
		time_of_registration TIMESTAMP,
		PRIMARY KEY (id)
	);`
	_, err = db.Exec(SQL_Schema_AccountCold)
	if err != nil {
		rendering.LOGERR(err)
	}
	db.Close()
}
```