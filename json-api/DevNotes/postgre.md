[page of go package](https://pkg.go.dev/github.com/lib/pq)  
```go
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	…
}  
```

# how to run sql script under running docker container  

```bash  
docker cp ./localfile.sql containername:/container/path/file.sql &&
docker exec -u postgresuser containername psql dbname postgresuser -f /container/path/file.sql  
```