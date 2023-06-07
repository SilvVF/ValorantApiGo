package types

import (
	"LFGbackend/graph/model"
	"time"
)

type PlayerData struct {
	SeasonId                string
	SeasonName              string
	Playlist                string
	Rank                    string
	IconUrl                 string
	MatchesPlayed           int
	MatchWinPct             float64
	Kills                   int
	KillsPercentile         float64
	KillsPerRound           float64
	KillsPerMatch           float64
	ScorePerRound           float64
	ScorePerRoundPercentile float64
	Assists                 int
	AssistsPerRound         float64
	AssistsPerMatch         float64
	Kd                      float64
	KdPercentile            float64
	Kda                     float64
	DmgPerRound             float64
	HeadshotPct             float64
	HeadshotPctPercentile   float64
	EconRating              float64
	FirstBloodsPerMatch     float64
	FirstDeathsPerRound     float64
	MostKillsInMatch        int
	TimePlayed              int
	TrnPerformanceScore     float64
	PeakRank                string
	PeakRankIconUrl         string
	PeakRankActName         string
}

type JoinPostRequest struct {
	User     User
	ClientId string
	PostId   string
}

type LfgSession struct {
	ClientId     string
	JoinedPostId string
	Data         PlayerData
}

type GormPlayer struct {
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	model.Player
}
