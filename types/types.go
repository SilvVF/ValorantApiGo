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

type JoinRequest struct {
	Player model.Player
	PostId string
	UserId string
}

type LeaveRequest struct {
	Player model.Player
	PostId string
	UserId string
}

type PostSession struct {
	ClientId string
}

type UserInfo struct {
	ClientId string
	Player   *model.Player
}

type User struct {
	Info  UserInfo
	State chan *model.Post
}

type GormPlayer struct {
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	model.Player
}

func (data *GormPlayer) AsPlayer() *model.Player {
	return &model.Player{
		Name:                    data.Name,
		Tag:                     data.Tag,
		SeasonID:                data.SeasonID,
		SeasonName:              data.SeasonName,
		Playlist:                data.Playlist,
		Rank:                    data.Rank,
		IconURL:                 data.IconURL,
		MatchesPlayed:           data.MatchesPlayed,
		MatchWinPct:             data.MatchWinPct,
		Kills:                   data.Kills,
		KillsPercentile:         data.KillsPercentile,
		KillsPerRound:           data.KillsPerRound,
		KillsPerMatch:           data.KillsPerMatch,
		ScorePerRound:           data.ScorePerRound,
		ScorePerRoundPercentile: data.ScorePerRoundPercentile,
		Assists:                 data.Assists,
		AssistsPerRound:         data.AssistsPerRound,
		AssistsPerMatch:         data.AssistsPerMatch,
		Kd:                      data.Kd,
		KdPercentile:            data.KdPercentile,
		Kda:                     data.Kda,
		DmgPerRound:             data.DmgPerRound,
		HeadshotPct:             data.HeadshotPct,
		HeadshotPctPercentile:   data.HeadshotPctPercentile,
		EconRating:              data.EconRating,
		FirstBloodsPerMatch:     data.FirstBloodsPerMatch,
		FirstDeathsPerRound:     data.FirstDeathsPerRound,
		MostKillsInMatch:        data.MostKillsInMatch,
		TimePlayed:              data.TimePlayed,
		TrnPerformanceScore:     data.TrnPerformanceScore,
		PeakRank:                data.PeakRank,
		PeakRankIconURL:         data.PeakRankIconURL,
		PeakRankActName:         data.PeakRankActName,
	}
}

func (data PlayerData) AsPlayer(name string, tag string) model.Player {
	return model.Player{
		Name:                    name,
		Tag:                     tag,
		SeasonID:                data.SeasonId,
		SeasonName:              data.SeasonName,
		Playlist:                data.Playlist,
		Rank:                    data.Rank,
		IconURL:                 data.IconUrl,
		MatchesPlayed:           data.MatchesPlayed,
		MatchWinPct:             data.MatchWinPct,
		Kills:                   data.Kills,
		KillsPercentile:         data.KillsPercentile,
		KillsPerRound:           data.KillsPerRound,
		KillsPerMatch:           data.KillsPerMatch,
		ScorePerRound:           data.ScorePerRound,
		ScorePerRoundPercentile: data.ScorePerRoundPercentile,
		Assists:                 data.Assists,
		AssistsPerRound:         data.AssistsPerRound,
		AssistsPerMatch:         data.AssistsPerMatch,
		Kd:                      data.Kd,
		KdPercentile:            data.KdPercentile,
		Kda:                     data.Kda,
		DmgPerRound:             data.DmgPerRound,
		HeadshotPct:             data.HeadshotPct,
		HeadshotPctPercentile:   data.HeadshotPctPercentile,
		EconRating:              data.EconRating,
		FirstBloodsPerMatch:     data.FirstBloodsPerMatch,
		FirstDeathsPerRound:     data.FirstDeathsPerRound,
		MostKillsInMatch:        data.MostKillsInMatch,
		TimePlayed:              data.TimePlayed,
		TrnPerformanceScore:     data.TrnPerformanceScore,
		PeakRank:                data.PeakRank,
		PeakRankIconURL:         data.PeakRankIconUrl,
		PeakRankActName:         data.PeakRankActName,
	}
}
