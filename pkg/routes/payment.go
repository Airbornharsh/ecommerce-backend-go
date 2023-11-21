package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func PaymentInit(user *gin.RouterGroup) {
	payment := user.Group("/payment")

	payment.POST("/", handlers.CreatePaymentHandler)
}
