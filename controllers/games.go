package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"piServer/common"
	"piServer/models"
	"time"
)

// GetGame handler for get game api call
func GetGame(c *gin.Context) {
	// Get game object

	// Return game object

}

// CreateGame handler for create new game api
func CreateGame(c *gin.Context) {
	// Create game object
	var game models.GameInput
	err := c.BindJSON(&game)
	if err != nil {
		//log err
		//return error
	}

	uniqueID, err := uuid.NewUUID()
	if err != nil {
		//log error
		//return error
	}

	player, _, err := GetPlayerDoc(c, c.Param("playerId"))
	if err != nil {
		//log and return
	}

	newGame := models.GameDoc{
		ID:         c.Param("playerId") + uniqueID.String(),
		PlayerName: player.Name,
		Date:       time.Now(),
		PlayerID:   c.Param("playerId"),
		Score:      game.Score,
	}

	// Store game object
	resp, rev, err := SaveGame(c, newGame)
	if err != nil {
		//log error
		//return error
	}

	// Return response
	response.SetResponse(c, rev, resp)

}

// UpdateGame handler for update game api
func UpdateGame(c *gin.Context) {
	// Get game object

	// Verify match

	// Update game object

	// Store updated object

}

// DeleteGame handler for delete game api
func DeleteGame(c *gin.Context) {

}

func createUniqueGameID(userID string) (gameID string) {
	return gameID
}
