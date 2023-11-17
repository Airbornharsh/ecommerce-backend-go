package handlers

import (
	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetProductsHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var products []models.Product

	q := "SELECT * FROM products"

	err := database.DB.QueryRow(q).Scan(&products)
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
		"message":  "Products Found",
		"token":    token,
		"products": products,
	})
}

func GetProductHandler(c *gin.Context) {

}

func FilterCategoryHandler(c *gin.Context) {

}

func PostProductsHandler(c *gin.Context) {

}

func PutProductsHandler(c *gin.Context) {

}

func DeleteProductsHandler(c *gin.Context) {

}
