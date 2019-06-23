package models

// PlayerDoc holds all the player info
type PlayerDoc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// PlayerInput holds the data input by a user to create a new player
type PlayerInput struct {
	Name string `json:"name"`
}
