package main

import (
	"database/sql"
	"github.com/graphql-go/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func main() {

	sqlDB, err := sql.Open("pgx", postgresUrl)
	if err != nil {
		log.Fatal(err)
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	h := handler.New(&handler.Config{
		Schema:   &PlayerSchema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
