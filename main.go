package main

import (
	"LFGbackend/graph"
	"LFGbackend/srv"
	"LFGbackend/types"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

func main() {

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to db. \n", err)
	}
	err = db.AutoMigrate(&types.GormPlayer{})
	if err != nil {
		log.Println(err)
	}

	server := srv.NewServer()
	sessions := make(map[string]*types.PostSession)

	s := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				Db:     db,
				Server: server,
			},
		},
	))
	s.AddTransport(&transport.Websocket{})
	mux := http.NewServeMux()

	mux.Handle("/", srv.Middleware(server, sessions, playground.Handler("GraphQL playground", "/query")))
	mux.Handle("/query", srv.Middleware(server, sessions, s))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
