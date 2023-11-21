package handlers

import (
	"errors"
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllWishlistHandler(c *gin.Context) {

}

func GetWishlistHandler(c *gin.Context) {

}

func CreateWishlistHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var wishlist models.Wishlist

	if err := c.ShouldBindJSON(&wishlist); helpers.ErrorResponse(c, err, 500) {
		return
	}

	wishlist.UserID = user.UserID

	var wishlistExists bool

	q := "SELECT EXISTS(SELECT 1 FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND name = '" + wishlist.Name + "' LIMIT 1)"

	err := database.DB.QueryRow(q).Scan(&wishlistExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if wishlistExists {
		helpers.ErrorResponse(c, errors.New("already present"), 409)
		return
	}

	if wishlist.DefaultProductID == 0 {
		q = "INSERT INTO wishlists (user_id, name) VALUES ('" + strconv.Itoa(int(user.UserID)) + "', '" + wishlist.Name + "') RETURNING wishlist_id"
	} else {
		q = "INSERT INTO wishlists (user_id, name, defaultproduct_id) VALUES ('" + strconv.Itoa(int(user.UserID)) + "', '" + wishlist.Name + "', '" + strconv.Itoa(int(wishlist.DefaultProductID)) + "') RETURNING wishlist_id"
	}

	err = database.DB.QueryRow(q).Scan(&wishlist.WishlistID)

	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "wishlist created",
		"token":   token,
		"data":    wishlist,
	})
}

func UpdateWishlistHandler(c *gin.Context) {

}

func DeleteWishlistHandler(c *gin.Context) {

}

func UpdateWishlistAddProductHandler(c *gin.Context) {

}

func UpdateWishlistRemoveProductHandler(c *gin.Context) {

}
