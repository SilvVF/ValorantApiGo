package main

import (
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"player": &graphql.Field{
			Type:        playerType,
			Description: "Get single player",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"tag": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, nameOk := p.Args["name"].(string)
				tag, tagOk := p.Args["tag"].(string)
				if nameOk && tagOk {
					// Search for el with name
					var player Player
					if err := db.First(&player, "id = ?", playerKey(name, tag)).Error; err != nil {
						pd := getPlayerData(name, tag)
						player = Player{
							ID:        playerKey(name, tag),
							Name:      name,
							Tag:       tag,
							Rank:      pd.rank,
							Kd:        pd.kd,
							Kda:       pd.kda,
							HsPct:     pd.headshotPct,
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						}
						if err := db.Create(&player).Error; err != nil {
							log.Println(err)
						}
						return player, nil
					} else {
						return player, nil
					}
				}
				return Player{}, nil
			},
		},
	},
})

var playerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Player",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"tag": &graphql.Field{
			Type: graphql.String,
		},
		"rank": &graphql.Field{
			Type: graphql.String,
		},
		"kd": &graphql.Field{
			Type: graphql.Float,
		},
		"kda": &graphql.Field{
			Type: graphql.Float,
		},
		"hsPct": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var playerDataType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PlayerData",
	Fields: graphql.Fields{
		"rank": &graphql.Field{
			Type: graphql.String,
		},
		"kd": &graphql.Field{
			Type: graphql.Float,
		},
		"kda": &graphql.Field{
			Type: graphql.Float,
		},
		"hsPct": &graphql.Field{
			Type: graphql.Float,
		},
	},
})

var PlayerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
