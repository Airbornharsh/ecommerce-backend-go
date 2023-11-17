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

}

func CreateAddressHandler(c *gin.Context) {

}

func UpdateAddressHandler(c *gin.Context) {

}

func DeleteAddressHandler(c *gin.Context) {

}
