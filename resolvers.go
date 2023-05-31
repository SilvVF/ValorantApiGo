package main

import (
	"github.com/graphql-go/graphql"
	"log"
	"time"
)

func playerResolver(p *graphql.ResolveParams) (Player, error) {
	name, nameOk := p.Args["name"].(string)
	tag, tagOk := p.Args["tag"].(string)
	fetch, fetchOk := p.Args["fetch"].(bool)
	if nameOk && tagOk {
		// Search for el with name
		var player Player
		if err := db.First(&player, "id = ?", playerKey(name, tag)).Error; err != nil || fetch || fetchOk {
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
			if err := db.Save(&player).Error; err != nil {
				log.Println(err)
			}
			return player, nil
		} else {
			return player, nil
		}
	}
	return Player{}, nil
}
