package handlers

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")

	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

	c.JSON(200, gin.H{
		"message": "User Found",
		"user":    user,
	})
}

func UpdateUserHandler(c *gin.Context) {
}
