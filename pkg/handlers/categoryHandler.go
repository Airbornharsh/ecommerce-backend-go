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
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var category models.Category

	q := "SELECT * FROM categories WHERE category_id = " + c.Param("id") + ";"

	row := database.DB.QueryRow(q)

	err := row.Scan(&category.CategoryID, &category.Name)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	user := tempUser.(models.User)
	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message":  "Category Found",
		"token":    token,
		"category": category,
	})
}

func CreateCategoryHandler(c *gin.Context) {
	tempAdmin, exists := c.Get("admin")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var category models.Category

	err := c.ShouldBindJSON(&category)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q := "INSERT INTO categories (name) VALUES ('" + category.Name + "');"

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	admin := tempAdmin.(models.User)
	token, err := helpers.GenerateToken(&admin)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Category Created",
		"token":   token,
	})
}

func UpdateCategoryHandler(c *gin.Context) {
	tempAdmin, exists := c.Get("admin")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var category models.Category

	err := c.ShouldBindJSON(&category)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if category.Name == "" {
		c.JSON(400, gin.H{
			"message": "Name is required",
		})
		return
	}

	q := "UPDATE categories SET name = '" + category.Name + "' WHERE category_id = " + c.Param("id") + ";"

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	admin := tempAdmin.(models.User)
	token, err := helpers.GenerateToken(&admin)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Category Updated",
		"token":   token,
	})
}

func DeleteCategoryHandler(c *gin.Context) {
	tempAdmin, exists := c.Get("admin")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	q := "DELETE FROM categories WHERE category_id = " + c.Param("id") + ";"

	_, err := database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	admin := tempAdmin.(models.User)
	token, err := helpers.GenerateToken(&admin)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Category Deleted",
		"token":   token,
	})
}
