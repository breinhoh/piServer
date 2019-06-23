package models

import "time"

// GameDoc holds all data for a saved game
type GameDoc struct {
	ID         string    `json:"id"`
	PlayerID   string    `json:"player_id"`
	PlayerName string    `json:"player_name"`
	Date       time.Time `json:"date"`
	Score      []int     `json:"score"`
}

// GameInput holds the data input by a user to save a new game
type GameInput struct {
	Score []int `json:"score"`
}
