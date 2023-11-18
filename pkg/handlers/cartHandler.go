package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

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

func GetCartHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

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

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message":   "Cart items retrieved successfully",
		"cartItems": cartItems,
		"token":     token,
	})
}

func AddProductCartHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

	var cartItem CartItem

	err := c.ShouldBindJSON(&cartItem)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q := "SELECT EXISTS (SELECT 1 FROM cartitems WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND product_id = '" + c.Param("productId") + "');"

	row := database.DB.QueryRow(q)

	var existsInCart bool

	err = row.Scan(&existsInCart)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if existsInCart {
		q = "UPDATE cartitems SET quantity = quantity + " + strconv.Itoa(cartItem.Quantity) + " WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND product_id = '" + c.Param("productId") + "';"
	} else {
		q = "INSERT INTO cartitems (user_id, product_id, quantity) VALUES ('" + strconv.Itoa(int(user.UserID)) + "', '" + c.Param("productId") + "', '" + strconv.Itoa(cartItem.Quantity) + "');"
	}

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Product added to cart",
		"token":   token,
	})
}

func UpdateProductCartHandler(c *gin.Context) {

}

func DeleteProductCartHandler(c *gin.Context) {

}

func DeleteCartHandler(c *gin.Context) {

}
