package main

import (
	"github.com/graphql-go/graphql"
)

var PlayerSchema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})

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
				"fetch": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return playerResolver(&p)
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
