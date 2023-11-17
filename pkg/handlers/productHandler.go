package handlers

import (
	"strconv"

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

	row, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for row.Next() {
		var product models.Product

		err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Image, &product.Quantity)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		products = append(products, product)
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
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var product models.Product

	q := "SELECT * FROM products WHERE product_id = " + c.Param("id") + ";"

	row := database.DB.QueryRow(q)

	err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Image, &product.Quantity)

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
		"message": "Product Found",
		"token":   token,
		"product": product,
	})
}

// func FilterCategoryHandler(c *gin.Context) {

// }

func PostProductsHandler(c *gin.Context) {
	tempAdmin, exists := c.Get("admin")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var product models.Product

	err := c.ShouldBindJSON(&product)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q := "INSERT INTO products (name, description, price, category_id, image, quantity) VALUES ('" + product.Name + "', '" + product.Description + "', " + strconv.Itoa(int(product.Price)) + ", " + strconv.Itoa(int(product.CategoryID)) + ", '" + product.Image + "', " + strconv.Itoa(int(product.Quantity)) + ");"

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
		"message": "Product Added",
		"token":   token,
	})

}

func PutProductsHandler(c *gin.Context) {

}

func DeleteProductsHandler(c *gin.Context) {

}
