package games

import (
	"github.com/gin-gonic/gin"
	"piServer/common"
	"piServer/models"
	"time"
)

func GetGame(c *gin.Context) {
	game := models.Game{
		Id:         "123",
		PlayerId:   "abcd",
		PlayerName: "Joe Bob",
		Date:       time.Now(),
		Score:      []int{1, 0, 1, 1, 1, 1, 1, 1, 1},
	}
	response.SetResponse(c, &game)
}
