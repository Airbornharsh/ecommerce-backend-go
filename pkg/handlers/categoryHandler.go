package handlers

import (
	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllCategoryHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var categories []models.Category

	q := "SELECT * FROM categories"

	row, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for row.Next() {
		var category models.Category

		err := row.Scan(&category.CategoryID, &category.Name)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		categories = append(categories, category)
	}

	user := tempUser.(models.User)
	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message":    "Categories Found",
		"token":      token,
		"categories": categories,
	})
}

func GetCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetCategoryHandler",
	})
}

func CreateCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateCategoryHandler",
	})
}

func UpdateCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateCategoryHandler",
	})
}

func DeleteCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteCategoryHandler",
	})
}
