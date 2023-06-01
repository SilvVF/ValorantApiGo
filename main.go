package main

import (
	"LFGbackend/graph"
	"LFGbackend/keys"
	"LFGbackend/src"
	"LFGbackend/types"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	db, err := gorm.Open(postgres.Open(keys.Dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(types.GormPlayer{})
	if err != nil {
		log.Println(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				Db: db,
			},
		},
	))
	mux := http.NewServeMux()

	mux.Handle("/", src.Middleware(playground.Handler("GraphQL playground", "/query")))
	mux.Handle("/query", src.Middleware(srv))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
