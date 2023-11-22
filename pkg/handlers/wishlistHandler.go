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
	user := c.MustGet("user").(models.User)

	var wishlists []models.Wishlist

	q := "SELECT wishlist_id, name, user_id, COALESCE(defaultproduct_id, 0) AS converted_defaultproduct_id FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "'"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		var wishlist models.Wishlist

		err := rows.Scan(&wishlist.WishlistID, &wishlist.Name, &wishlist.UserID, &wishlist.DefaultProductID)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		wishlists = append(wishlists, wishlist)
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message":   "wishlists fetched",
		"token":     token,
		"wishlists": wishlists,
	})
}

func GetWishlistHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	type Wishlist struct {
		models.Wishlist
		Products []models.Product `json:"products"`
	}

	var wishlist Wishlist

	q := "SELECT wishlist_id, name, user_id, COALESCE(defaultproduct_id, 0) AS converted_defaultproduct_id FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND wishlist_id = '" + c.Param("id") + "'"

	err := database.DB.QueryRow(q).Scan(&wishlist.WishlistID, &wishlist.Name, &wishlist.UserID, &wishlist.DefaultProductID)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q = "SELECT product_id, name, description, price, category_id, image, quantity FROM products  WHERE product_id IN (SELECT product_id FROM wishlistitems WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND wishlist_id = '" + c.Param("id") + "')"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Image, &product.Quantity)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		wishlist.Products = append(wishlist.Products, product)
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "wishlist fetched",
		"token":   token,
		"data":    wishlist,
	})
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
	user := c.MustGet("user").(models.User)

	var wishlist models.Wishlist

	if err := c.ShouldBindJSON(&wishlist); helpers.ErrorResponse(c, err, 500) {
		return
	}

	wishlist.UserID = user.UserID

	var wishlistExists bool

	q := "SELECT EXISTS(SELECT 1 FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND wishlist_id = '" + c.Param("id") + "' LIMIT 1)"

	err := database.DB.QueryRow(q).Scan(&wishlistExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if !wishlistExists {
		helpers.ErrorResponse(c, errors.New("not present"), 409)
		return
	}

	q = "UPDATE wishlists SET"

	if wishlist.Name != "" {
		q += " name = '" + wishlist.Name + "',"
	}
	if wishlist.DefaultProductID != 0 {
		q += " defaultproduct_id = '" + strconv.Itoa(int(wishlist.DefaultProductID)) + "',"
	}

	q = q[:len(q)-1]

	q += " WHERE wishlist_id = '" + c.Param("id") + "'"

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "wishlist updated",
		"token":   token,
	})
}

func DeleteWishlistHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var wishlistExists bool

	q := "SELECT EXISTS(SELECT 1 FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND wishlist_id = '" + c.Param("id") + "' LIMIT 1)"

	err := database.DB.QueryRow(q).Scan(&wishlistExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if !wishlistExists {
		helpers.ErrorResponse(c, errors.New("not present"), 409)
		return
	}

	q = "DELETE FROM wishlists WHERE wishlist_id = '" + c.Param("id") + "'"

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "wishlist deleted",
		"token":   token,
	})
}

func UpdateWishlistAddProductHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var wishlistExists bool

	q := "SELECT EXISTS (SELECT 1 FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND wishlist_id = '" + c.Param("id") + "')"

	err := database.DB.QueryRow(q).Scan(&wishlistExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if !wishlistExists {
		helpers.ErrorResponse(c, errors.New("not present"), 409)
		return
	}

	var productExists bool

	q = "SELECT EXISTS (SELECT 1 FROM wishlistitems WHERE product_id = '" + c.Param("productId") + "' AND wishlist_id = '" + c.Param("id") + "')"

	err = database.DB.QueryRow(q).Scan(&productExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if productExists {
		c.JSON(200, gin.H{
			"message": "Product Already Added",
		})
		return
	}

	var WishlistItem models.WishlistItem

	WishlistItem.UserID = user.UserID

	q = "INSERT INTO wishlistitems(user_id, product_id, wishlist_id) VALUES ('" + strconv.Itoa(int(user.UserID)) + "','" + c.Param("productId") + "','" + c.Param("id") + "') RETURNING wishlistitem_id,product_id,wishlist_id"

	err = database.DB.QueryRow(q).Scan(&WishlistItem.WishlistItemID, &WishlistItem.ProductID, &WishlistItem.WishlistID)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message":      "Added Product to Wishlist",
		"token":        token,
		"wishlistItem": WishlistItem,
	})
}

func UpdateWishlistRemoveProductHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var wishlistExists bool

	q := "SELECT EXISTS (SELECT 1 FROM wishlists WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND wishlist_id = '" + c.Param("id") + "')"

	err := database.DB.QueryRow(q).Scan(&wishlistExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if !wishlistExists {
		helpers.ErrorResponse(c, errors.New("not present"), 409)
		return
	}

	var productExists bool

	q = "SELECT EXISTS (SELECT 1 FROM wishlistitems WHERE product_id = '" + c.Param("productId") + "' AND wishlist_id = '" + c.Param("id") + "')"

	err = database.DB.QueryRow(q).Scan(&productExists)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if !productExists {
		c.JSON(200, gin.H{
			"message": "Product Not Present",
		})
		return
	}

	q = "DELETE FROM wishlistitems WHERE product_id = '" + c.Param("productId") + "' AND wishlist_id = '" + c.Param("id") + "'"

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Removed Product from Wishlist",
		"token":   token,
	})
}
