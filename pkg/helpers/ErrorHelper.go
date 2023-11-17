package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, err error, statusCode int) bool {
	if err != nil {
		fmt.Println(err)
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return true
	}
	return false
}
