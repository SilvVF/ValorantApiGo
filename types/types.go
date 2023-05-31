package types

import "LFGbackend/graph/model"

type PlayerData struct {
	SeasonId            string
	Playlist            string
	Rank                string
	IconUrl             string
	MatchesPlayed       int
	MatchWinPct         float64
	KillsPerMatch       float64
	Kd                  float64
	Kda                 float64
	DmgPerRound         float64
	HeadshotPct         float64
	FirstBloodsPerMatch float64
	FirstDeathsPerRound float64
	MostKillsInMatch    int
}

func (data PlayerData) AsPlayer(name string, tag string) model.Player {
	return model.Player{
		ID:                  name + tag,
		Name:                name,
		Tag:                 tag,
		Rank:                data.Rank,
		IconURL:             &data.IconUrl,
		MatchesPlayed:       data.MatchesPlayed,
		MatchWinPct:         data.MatchWinPct,
		KillsPerMatch:       data.KillsPerMatch,
		Kd:                  data.Kd,
		Kda:                 data.Kda,
		DmgPerRound:         data.DmgPerRound,
		HeadshotPct:         data.HeadshotPct,
		FirstBloodsPerMatch: data.FirstBloodsPerMatch,
		FirstDeathsPerRound: data.FirstDeathsPerRound,
		MostKillsInMatch:    data.MostKillsInMatch,
	}
}
