package main

import (
	"LFGbackend/graph"
	"LFGbackend/lfg"
	"LFGbackend/middleware"
	"LFGbackend/types"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	lfgServer := lfg.NewLfgServer()
	sessionManger := lfg.NewManager()

	router.Use(middleware.Middleware())

	graphqlServer := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				Db:             db,
				SessionManager: sessionManger,
				Server:         lfgServer,
			},
		},
	))

	graphqlServer.AddTransport(&transport.Websocket{})

	router.HandleFunc(
		"/",
		playground.Handler("playground", "/graphql"),
	)
	router.Handle("/graphql", graphqlServer)
	router.HandleFunc("/post/{id}", websocketJoinPostHandler(sessionManger, lfgServer))
	router.HandleFunc("/create/{needed}/{minrank}/{gamemode}", websocketCreatePostHandler(sessionManger, lfgServer))

	log.Fatal(http.ListenAndServe(":8080", router))
}
