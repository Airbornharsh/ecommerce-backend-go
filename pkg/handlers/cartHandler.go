package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetCartHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

	type CartItem struct {
		CartItemID  uint   `json:"cartitem_id"`
		UserID      uint   `json:"user_id"`
		ProductID   uint   `json:"product_id"`
		Quantity    int    `json:"quantity"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       uint   `json:"price"`
		CategoryID  uint   `json:"category_id"`
		Category    string `json:"category"`
		Image       string `json:"image"`
	}

	var cartItems []CartItem

	q := "SELECT c.cartitem_id, c.user_id, c.product_id, c.quantity, p.name, p.description, p.price, p.category_id, cat.name, p.image FROM cartitems c INNER JOIN products p ON c.product_id = p.product_id INNER JOIN categories cat ON p.category_id = cat.category_id WHERE c.user_id = '" + strconv.Itoa(int(user.UserID)) + "';"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		var cartItem CartItem

		rows.Scan(&cartItem.CartItemID, &cartItem.UserID, &cartItem.ProductID, &cartItem.Quantity, &cartItem.Name, &cartItem.Description, &cartItem.Price, &cartItem.CategoryID, &cartItem.Category, &cartItem.Image)
		cartItems = append(cartItems, cartItem)
	}

	c.JSON(200, gin.H{
		"cartItems": cartItems,
	})
}

func AddProductCartHandler(c *gin.Context) {

}

func UpdateProductCartHandler(c *gin.Context) {

}

func DeleteProductCartHandler(c *gin.Context) {

}

func DeleteCartHandler(c *gin.Context) {

}
