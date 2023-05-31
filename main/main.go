package main

import (
	"LFGbackend/graph"
	"LFGbackend/keys"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ *gorm.DB

func main() {

	db, err := gorm.Open(postgres.Open(keys.Dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{Db: db},
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
