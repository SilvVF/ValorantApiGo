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
