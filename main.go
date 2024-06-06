// main.go
package main

import (
	"log"
	"database/sql"
	"myblog/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:31650000@localhost/my_blog?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
