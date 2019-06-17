package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"piServer/models"
)

func SetResponse(c *gin.Context, game *models.Game) {

	c.JSON(http.StatusOK, game)
}

func SetErrorResponse(c *gin.Context, err *error) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "We are having difficulties at the moment.  Please try again later.",
	})
}
