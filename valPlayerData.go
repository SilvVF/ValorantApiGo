package main

import (
	"encoding/json"
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"log"
	"strings"
	"time"
)

var client = cycletls.Init()

func playerKey(name string, tag string) string {
	return name + "#" + tag
}

func getPlayerData(name string, tag string) PlayerData {

	url := "https://tracker.gg/valorant/profile/riot/" + name + "%23" + tag + "/overview"

	response, err := client.Do(url, cycletls.Options{
		Body:      "",
		Ja3:       "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-51-57-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
		UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
	}, "GET")
	if err != nil {
		log.Print("Request Failed: " + err.Error())
	}

	start := strings.Index(response.Body, "\"segments\"")
	end := strings.Index(response.Body, "\"availableSegments\"")
	jsonString := "{" + response.Body[start:end-1] + "}"
	var data ScrapedPlayerData
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Println(err)
		return PlayerData{}
	}

	comp := data.Segments[0]

	compPlayerData := PlayerData{
		seasonId:            comp.Attributes.SeasonID,
		playlist:            comp.Attributes.Playlist,
		rank:                comp.Stats.Rank.Metadata.TierName,
		iconUrl:             comp.Stats.Rank.Metadata.IconURL,
		matchesPlayed:       comp.Stats.MatchesPlayed.Value,
		matchWinPct:         comp.Stats.MatchesWinPct.Value,
		killsPerMatch:       comp.Stats.KillsPerMatch.Value,
		kd:                  comp.Stats.KDRatio.Value,
		kda:                 comp.Stats.KDARatio.Value,
		dmgPerRound:         comp.Stats.DamagePerRound.Value,
		headshotPct:         comp.Stats.HeadshotsPercentage.Value,
		firstBloodsPerMatch: comp.Stats.FirstBloodsPerMatch.Value,
		firstDeathsPerRound: comp.Stats.FirstDeathsPerRound.Value,
	}

	return compPlayerData
}

type PlayerData struct {
	seasonId            string
	playlist            string
	rank                string
	iconUrl             string
	matchesPlayed       int
	matchWinPct         float64
	killsPerMatch       float64
	kd                  float64
	kda                 float64
	dmgPerRound         float64
	headshotPct         float64
	firstBloodsPerMatch float64
	firstDeathsPerRound float64
	mostKillsInMatch    int
}

type ScrapedPlayerData struct {
	Segments []struct {
		Type       string `json:"type"`
		Attributes struct {
			SeasonID string `json:"seasonId"`
			Playlist string `json:"playlist"`
		} `json:"attributes"`
		Metadata struct {
			Name         string    `json:"name"`
			PlaylistName string    `json:"playlistName"`
			StartTime    time.Time `json:"startTime"`
			EndTime      time.Time `json:"endTime"`
			Schema       string    `json:"schema"`
		} `json:"metadata"`
		ExpiryDate time.Time `json:"expiryDate"`
		Stats      struct {
			MatchesPlayed struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"matchesPlayed"`
			MatchesWon struct {
				Rank            interface{} `json:"rank"`
				Percentile      int         `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"matchesWon"`
			MatchesLost struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"matchesLost"`
			MatchesTied struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"matchesTied"`
			MatchesWinPct struct {
				Rank            interface{} `json:"rank"`
				Percentile      float64     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"matchesWinPct"`
			MatchesDuration struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"matchesDuration"`
			TimePlayed struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"timePlayed"`
			RoundsPlayed struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"roundsPlayed"`
			RoundsWon struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"roundsWon"`
			RoundsLost struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"roundsLost"`
			RoundsWinPct struct {
				Rank            interface{} `json:"rank"`
				Percentile      float64     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"roundsWinPct"`
			RoundsDuration struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"roundsDuration"`
			Score struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"score"`
			ScorePerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"scorePerMatch"`
			ScorePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      int         `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"scorePerRound"`
			Kills struct {
				Rank            interface{} `json:"rank"`
				Percentile      int         `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills"`
			KillsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"killsPerRound"`
			KillsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"killsPerMatch"`
			Deaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"deaths"`
			DeathsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"deathsPerRound"`
			DeathsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"deathsPerMatch"`
			Assists struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"assists"`
			AssistsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"assistsPerRound"`
			AssistsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"assistsPerMatch"`
			KDRatio struct {
				Rank            interface{} `json:"rank"`
				Percentile      float64     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"kDRatio"`
			KDARatio struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"kDARatio"`
			KADRatio struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"kADRatio"`
			Damage struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"damage"`
			DamageDelta struct {
				Rank            interface{} `json:"rank"`
				Percentile      int         `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"damageDelta"`
			DamageDeltaPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      float64     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     string      `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"damageDeltaPerRound"`
			DamagePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      int         `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"damagePerRound"`
			DamagePerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"damagePerMatch"`
			DamagePerMinute struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"damagePerMinute"`
			DamageReceived struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"damageReceived"`
			Headshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"headshots"`
			HeadshotsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"headshotsPerRound"`
			HeadshotsPercentage struct {
				Rank            interface{} `json:"rank"`
				Percentile      int         `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"headshotsPercentage"`
			GrenadeCasts struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"grenadeCasts"`
			GrenadeCastsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"grenadeCastsPerRound"`
			GrenadeCastsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"grenadeCastsPerMatch"`
			Ability1Casts struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"ability1Casts"`
			Ability1CastsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"ability1CastsPerRound"`
			Ability1CastsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"ability1CastsPerMatch"`
			Ability2Casts struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"ability2Casts"`
			Ability2CastsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"ability2CastsPerRound"`
			Ability2CastsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"ability2CastsPerMatch"`
			UltimateCasts struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"ultimateCasts"`
			UltimateCastsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"ultimateCastsPerRound"`
			UltimateCastsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"ultimateCastsPerMatch"`
			DealtHeadshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"dealtHeadshots"`
			DealtBodyshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"dealtBodyshots"`
			DealtLegshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"dealtLegshots"`
			ReceivedHeadshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"receivedHeadshots"`
			ReceivedBodyshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"receivedBodyshots"`
			ReceivedLegshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"receivedLegshots"`
			EconRating struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"econRating"`
			EconRatingPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"econRatingPerMatch"`
			EconRatingPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"econRatingPerRound"`
			Suicides struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"suicides"`
			FirstBloods struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"firstBloods"`
			FirstBloodsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"firstBloodsPerRound"`
			FirstBloodsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"firstBloodsPerMatch"`
			FirstDeaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"firstDeaths"`
			FirstDeathsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"firstDeathsPerRound"`
			LastDeaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"lastDeaths"`
			Survived struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"survived"`
			Traded struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"traded"`
			KAST struct {
				Rank            interface{} `json:"rank"`
				Percentile      float64     `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     string      `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"kAST"`
			MostKillsInMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"mostKillsInMatch"`
			Flawless struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"flawless"`
			Thrifty struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"thrifty"`
			Aces struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"aces"`
			TeamAces struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"teamAces"`
			Clutches struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutches"`
			ClutchesPercentage struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"clutchesPercentage"`
			ClutchesLost struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutchesLost"`
			Clutches1V1 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutches1v1"`
			Clutches1V2 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutches1v2"`
			Clutches1V3 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutches1v3"`
			Clutches1V4 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutches1v4"`
			Clutches1V5 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutches1v5"`
			ClutchesLost1V1 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutchesLost1v1"`
			ClutchesLost1V2 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutchesLost1v2"`
			ClutchesLost1V3 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutchesLost1v3"`
			ClutchesLost1V4 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutchesLost1v4"`
			ClutchesLost1V5 struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"clutchesLost1v5"`
			Kills1K struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills1K"`
			Kills2K struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills2K"`
			Kills3K struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills3K"`
			Kills4K struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills4K"`
			Kills5K struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills5K"`
			Kills6K struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"kills6K"`
			Plants struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"plants"`
			PlantsPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"plantsPerMatch"`
			PlantsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"plantsPerRound"`
			AttackKills struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackKills"`
			AttackKillsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackKillsPerRound"`
			AttackDeaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackDeaths"`
			AttackKDRatio struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackKDRatio"`
			AttackAssists struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackAssists"`
			AttackAssistsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackAssistsPerRound"`
			AttackRoundsWon struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackRoundsWon"`
			AttackRoundsLost struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackRoundsLost"`
			AttackRoundsPlayed struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackRoundsPlayed"`
			AttackRoundsWinPct struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackRoundsWinPct"`
			AttackScore struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackScore"`
			AttackScorePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackScorePerRound"`
			AttackDamage struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackDamage"`
			AttackDamagePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackDamagePerRound"`
			AttackHeadshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackHeadshots"`
			AttackTraded struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackTraded"`
			AttackSurvived struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackSurvived"`
			AttackFirstBloods struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackFirstBloods"`
			AttackFirstBloodsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackFirstBloodsPerRound"`
			AttackFirstDeaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"attackFirstDeaths"`
			AttackFirstDeathsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackFirstDeathsPerRound"`
			AttackKAST struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"attackKAST"`
			Defuses struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defuses"`
			DefusesPerMatch struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defusesPerMatch"`
			DefusesPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defusesPerRound"`
			DefenseKills struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseKills"`
			DefenseKillsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseKillsPerRound"`
			DefenseDeaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseDeaths"`
			DefenseKDRatio struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseKDRatio"`
			DefenseAssists struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseAssists"`
			DefenseAssistsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseAssistsPerRound"`
			DefenseRoundsWon struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseRoundsWon"`
			DefenseRoundsLost struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseRoundsLost"`
			DefenseRoundsPlayed struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseRoundsPlayed"`
			DefenseRoundsWinPct struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseRoundsWinPct"`
			DefenseScore struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseScore"`
			DefenseScorePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseScorePerRound"`
			DefenseDamage struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseDamage"`
			DefenseDamagePerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseDamagePerRound"`
			DefenseHeadshots struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseHeadshots"`
			DefenseTraded struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseTraded"`
			DefenseSurvived struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseSurvived"`
			DefenseFirstBloods struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseFirstBloods"`
			DefenseFirstBloodsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseFirstBloodsPerRound"`
			DefenseFirstDeaths struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseFirstDeaths"`
			DefenseFirstDeathsPerRound struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
				DisplayType  string  `json:"displayType"`
			} `json:"defenseFirstDeathsPerRound"`
			DefenseKAST struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory string      `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"defenseKAST"`
			Rank struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory interface{} `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
					IconURL  string `json:"iconUrl"`
					TierName string `json:"tierName"`
				} `json:"metadata"`
				Value        interface{} `json:"value"`
				DisplayValue string      `json:"displayValue"`
				DisplayType  string      `json:"displayType"`
			} `json:"rank"`
			TrnPerformanceScore struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory interface{} `json:"displayCategory"`
				Category        interface{} `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
					Stats []string `json:"stats"`
				} `json:"metadata"`
				Value        int    `json:"value"`
				DisplayValue string `json:"displayValue"`
				DisplayType  string `json:"displayType"`
			} `json:"trnPerformanceScore"`
			PeakRank struct {
				Rank            interface{} `json:"rank"`
				Percentile      interface{} `json:"percentile"`
				DisplayName     string      `json:"displayName"`
				DisplayCategory interface{} `json:"displayCategory"`
				Category        string      `json:"category"`
				Description     interface{} `json:"description"`
				Metadata        struct {
					IconURL  string `json:"iconUrl"`
					TierName string `json:"tierName"`
					ActID    string `json:"actId"`
					ActName  string `json:"actName"`
				} `json:"metadata"`
				Value        interface{} `json:"value"`
				DisplayValue string      `json:"displayValue"`
				DisplayType  string      `json:"displayType"`
			} `json:"peakRank"`
		} `json:"stats"`
	} `json:"segments"`
}
