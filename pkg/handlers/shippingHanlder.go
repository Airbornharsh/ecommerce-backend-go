package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateShippingHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var shipping models.Shipping

	err := c.ShouldBindJSON(&shipping)

	if helpers.ErrorResponse(c, err, 400) {
		return
	}

	shipping.UserID = user.UserID

	q := "INSERT INTO shippings (user_id, method, address_id, estimated_delivery_days) VALUES ('" + strconv.Itoa(int(shipping.UserID)) + "', '" + shipping.Method + "', '" + strconv.Itoa(int(shipping.AddressID)) + "', '" + strconv.Itoa(shipping.EstimatedDeliveryDays) + "') RETURNING shipping_id"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var shippingID uint

	for rows.Next() {
		err = rows.Scan(&shippingID)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}
	}

	shipping.ShippingID = shippingID

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message":  "Shipping created successfully",
		"token":    token,
		"shipping": shipping,
	})
}
