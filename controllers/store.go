package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"piServer/common/database"
	"piServer/models"
)

// SaveGame function that save the input game doc to couchDB
func SaveGame(c *gin.Context, newgame models.GameDoc) (models.GameDoc, string, error) {

	//get database
	db := database.GetDb()
	exist := db.DbExists()
	if exist != nil {
		fmt.Println("We have a problem here")
	}

	rev, err := db.Save(newgame, newgame.ID, "")
	if err != nil {
		//log and return
	}

	return newgame, rev, nil

}

// GetPlayerDoc returns a player doc from the database if found
func GetPlayerDoc(c *gin.Context, playerID string) (models.PlayerDoc, string, error) {
	//get database
	db := database.GetDb()
	exist := db.DbExists()
	if exist != nil {
		fmt.Println("We have a problem here")
	}

	var player models.PlayerDoc

	rev, err := db.Read(playerID, &player, nil)
	if err != nil {
		//log and return
	}
	return player, rev, nil
}

// SavePlayer function that saves the new player doc to the database
func SavePlayer(c *gin.Context, newplayer models.PlayerDoc) (models.PlayerDoc, string, error) {
	//get database
	db := database.GetDb()
	exist := db.DbExists()
	if exist != nil {
		fmt.Println("We have a problem here")
	}

	rev, err := db.Save(newplayer, newplayer.ID, "")
	if err != nil {
		//log and return
	}

	return newplayer, rev, nil
}
