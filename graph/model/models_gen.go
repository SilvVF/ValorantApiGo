// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Player struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Tag                 string  `json:"tag"`
	Rank                string  `json:"rank"`
	IconURL             *string `json:"iconUrl,omitempty"`
	MatchesPlayed       int     `json:"matchesPlayed"`
	MatchWinPct         float64 `json:"matchWinPct"`
	KillsPerMatch       float64 `json:"killsPerMatch"`
	Kd                  float64 `json:"kd"`
	Kda                 float64 `json:"kda"`
	DmgPerRound         float64 `json:"dmgPerRound"`
	HeadshotPct         float64 `json:"headshotPct"`
	FirstBloodsPerMatch float64 `json:"firstBloodsPerMatch"`
	FirstDeathsPerRound float64 `json:"firstDeathsPerRound"`
	MostKillsInMatch    int     `json:"mostKillsInMatch"`
}

type PlayerInput struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}