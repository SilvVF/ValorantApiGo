package main

import (
	"LFGbackend/graph"
	"LFGbackend/graph/model"
	"LFGbackend/keys"
	"LFGbackend/types"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	url := "dkjflakjf"
	db.Create(types.GormPlayer{
		Id:        "silv004",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Player: model.Player{
			Name:                "silv",
			Tag:                 "004",
			Rank:                "Unranked",
			IconURL:             &url,
			MatchesPlayed:       0,
			MatchWinPct:         0,
			KillsPerMatch:       0,
			Kd:                  0,
			Kda:                 0,
			DmgPerRound:         0,
			HeadshotPct:         0,
			FirstBloodsPerMatch: 0,
			FirstDeathsPerRound: 0,
			MostKillsInMatch:    0,
		},
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &graph.Resolver{
				Db: db,
			},
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
