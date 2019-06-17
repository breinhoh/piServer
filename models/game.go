package models

import "time"

type Game struct {
	Id         string    `json:"id"`
	PlayerId   string    `json:"player_id"`
	PlayerName string    `json:"player_name"`
	Date       time.Time `json:"date"`
	Score      []int     `json:"score"`
}
