package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ProductID   uint   `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	CategoryID  uint   `json:"category_id"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Quantity    uint   `json:"quantity"`
}

func GetProductsHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var products []Product

	q := "SELECT p.product_id, p.name, p.description, p.price, p.category_id, cat.name, p.image, p.quantity FROM products p INNER JOIN categories cat ON p.category_id = cat.category_id;"

	row, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for row.Next() {
		var product Product

		err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Category, &product.Image, &product.Quantity)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		products = append(products, product)
	}

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
	user := c.MustGet("user").(models.User)

	var product Product

	q := "SELECT  p.product_id, p.name, p.description, p.price, p.category_id, cat.name, p.image, p.quantity  FROM products p INNER JOIN categories cat ON p.category_id = cat.category_id WHERE product_id = " + c.Param("id") + ";"

	row := database.DB.QueryRow(q)

	err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Category, &product.Image, &product.Quantity)

	if helpers.ErrorResponse(c, err, 500) {
		return
	}

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

func FilterCategoryHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var products []Product

	q := "SELECT  p.product_id, p.name, p.description, p.price, p.category_id, cat.name, p.image, p.quantity  FROM products p INNER JOIN categories cat ON p.category_id = cat.category_id WHERE category_id = " + c.Param("category") + ";"

	row, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for row.Next() {
		var product Product

		err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Category, &product.Image, &product.Quantity)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		products = append(products, product)
	}

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

func PostProductsHandler(c *gin.Context) {
	admin := c.MustGet("admin").(models.User)

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
	admin := c.MustGet("admin").(models.User)

	var newProduct models.Product

	err := c.ShouldBindJSON(&newProduct)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q := "UPDATE products SET "
	if newProduct.Name != "" {
		q += "name = '" + newProduct.Name + "', "
	}
	if newProduct.Description != "" {
		q += "description = '" + newProduct.Description + "', "
	}
	if newProduct.Price != 0 {
		q += "price = " + strconv.Itoa(int(newProduct.Price)) + ", "
	}
	if newProduct.CategoryID != 0 {
		q += "category_id = " + strconv.Itoa(int(newProduct.CategoryID)) + ", "
	}
	if newProduct.Image != "" {
		q += "image = '" + newProduct.Image + "', "
	}
	if newProduct.Quantity != 0 {
		q += "quantity = " + strconv.Itoa(int(newProduct.Quantity)) + ", "
	}

	q = q[:len(q)-2]
	q += " WHERE product_id = " + c.Param("id") + ";"

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&admin)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Product Updated",
		"token":   token,
	})

}

func DeleteProductsHandler(c *gin.Context) {
	admin := c.MustGet("admin").(models.User)

	q := "DELETE FROM products WHERE product_id = " + c.Param("id") + ";"

	_, err := database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&admin)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Product Deleted",
		"token":   token,
	})
}
