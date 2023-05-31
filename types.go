package main

import (
	"time"
)

type Player struct {
	ID        string `gorm:"primaryKey"`
	Name      string `json:"name"`
	Tag       string `json:"tag"`
	Rank      string
	Kd        float64
	Kda       float64
	HsPct     float64
	CreatedAt time.Time
	UpdatedAt time.Time
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
