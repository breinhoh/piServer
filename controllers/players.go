package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"piServer/common"
	"piServer/models"
)

// GetPlayer handler for get player api
func GetPlayer(c *gin.Context) {
	player, rev, err := GetPlayerDoc(c, c.Param("playerId"))
	if err != nil {
		//log and return
	}

	response.SetResponse(c, rev, player)

}

// CreatePlayer handler for create player api
func CreatePlayer(c *gin.Context) {
	var player models.PlayerInput
	err := c.BindJSON(&player)
	if err != nil {
		//log err
		//return error
	}

	uniqueID, err := uuid.NewUUID()
	if err != nil {
		//log error
		//return error
	}

	//playerName := getPlayer()

	newPlayer := models.PlayerDoc{
		ID:   uniqueID.String(),
		Name: player.Name,
	}

	resp, rev, err := SavePlayer(c, newPlayer)
	if err != nil {
		//log error
		//return error
	}
	response.SetResponse(c, rev, resp)

}

// UpdatePlayer hadler for update player api
func UpdatePlayer(c *gin.Context) {

}

// DeletePlayer handler for delete player api
func DeletePlayer(c *gin.Context) {

}
