package routes

import (
	"github.com/gin-gonic/gin"
	"piServer/controllers"
)

func Init() *gin.Engine {

	r := gin.Default()

	r.GET("/get-player/:playerId", controllers.GetPlayer)
	r.POST("/new-player", controllers.CreatePlayer)
	r.PUT("/update-player/:playerId", controllers.UpdatePlayer)
	r.DELETE("delete-player/:playerId", controllers.DeletePlayer)

	r.GET("/get-game/:playerId/:gameId", controllers.GetGame)
	r.POST("/new-game/:playerId", controllers.CreateGame)
	r.PUT("/update-game/:playerId/:gameId", controllers.UpdateGame)
	r.DELETE("/delete-game/:playerId/:gameId", controllers.DeleteGame)

	return r
}
