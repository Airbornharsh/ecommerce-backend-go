package handlers

import (
	"strconv"
	"time"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

type Payment struct {
	PaymentID uint   `json:"payment_id"`
	UserID    uint   `json:"user_id"`
	Amount    uint   `json:"amount"`
	Method    string `json:"method"`
	Status    string `json:"status"`
}

func CreatePaymentHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var payment Payment

	if err := c.ShouldBindJSON(&payment); helpers.ErrorResponse(c, err, 400) {
		return
	}

	payment.UserID = user.UserID
	payment.Status = "pending"
	method, err := helpers.PaymentMethodConverter(payment.Method)
	if helpers.ErrorResponse(c, err, 400) {
		return
	}

	payment.Method = string(method)

	q := "INSERT INTO payments (user_id, amount, method, status, created_at) VALUES ('" + strconv.Itoa(int(payment.UserID)) + "', '" + strconv.Itoa(int(payment.Amount)) + "', '" + payment.Method + "', '" + payment.Status + "', '" + time.Now().Format("2006-01-02 15:04:05") + "') RETURNING payment_id"

	err = database.DB.QueryRow(q).Scan(&payment.PaymentID)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	// if payment.PaymentMethod == "cash_on_delivery" {
	// 	q = "UPDATE payments SET status = 'success' WHERE payment_id = '" + strconv.Itoa(int(payment.PaymentID)) + "'"
	// 	_, err = database.DB.Exec(q)
	// 	if helpers.ErrorResponse(c, err, 500) {
	// 		return
	// 	}
	// }

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Payment created successfully",
		"token":   token,
		"payment": payment,
	})
}

func FailedPaymentHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	paymentID := c.Param("paymentId")

	q := "UPDATE payments SET status = 'failed' WHERE payment_id = '" + paymentID + "' AND user_id = '" + strconv.Itoa(int(user.UserID)) + "';"

	_, err := database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Payment failed successfully",
		"token":   token,
	})
}

func SuccessPaymentHandler(c *gin.Context) {

}

func CancelPaymentHandler(c *gin.Context) {

}
