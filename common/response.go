package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetResponse create valid response
func SetResponse(c *gin.Context, rev string, doc interface{}) {

	c.Writer.Header().Set("ETAG", rev)
	c.JSON(http.StatusOK, doc)
}

// SetErrorResponse set error response
func SetErrorResponse(c *gin.Context, err *error) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "We are having difficulties at the moment.  Please try again later.",
	})
}
