package routes

import (
	"github.com/gin-gonic/gin"
	"piServer/controllers"
)

func Init() *gin.Engine {

	r := gin.Default()

	r.GET("/:gameID", games.GetGame)

	return r
}
