package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllAddressHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

	var addresses []models.Address

	q := "SELECT * FROM addresses WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "'"

	row, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for row.Next() {
		var address models.Address

		err := row.Scan(&address.AddressID, &address.UserID, &address.Street, &address.City, &address.State, &address.Country, &address.ZipCode, &address.IsDefault)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		addresses = append(addresses, address)
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message":   "Address Found",
		"token":     token,
		"addresses": addresses,
	})
}

func GetAddressHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

	var address models.Address

	q := "SELECT * FROM addresses WHERE address_id = " + c.Param("id") + " AND user_id = '" + strconv.Itoa(int(user.UserID)) + "';"

	row := database.DB.QueryRow(q)

	err := row.Scan(&address.AddressID, &address.UserID, &address.Street, &address.City, &address.State, &address.Country, &address.ZipCode, &address.IsDefault)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "Address Found",
		"token":   token,
		"address": address,
	})
}

func CreateAddressHandler(c *gin.Context) {
	tempUser, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	user := tempUser.(models.User)

	var address models.Address

	err := c.ShouldBindJSON(&address)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	address.IsDefault = false

	q := "INSERT INTO addresses (user_id, street, city, state, country, zip_code, is_default) VALUES ('" + strconv.Itoa(int(user.UserID)) + "', '" + address.Street + "', '" + address.City + "', '" + address.State + "', '" + address.Country + "', '" + address.ZipCode + "', '" + strconv.FormatBool(address.IsDefault) + "');"

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
		"message": "Address Created",
		"token":   token,
	})
}

func UpdateAddressHandler(c *gin.Context) {

}

func DeleteAddressHandler(c *gin.Context) {

}
