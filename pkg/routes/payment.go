package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func PaymentInit(user *gin.RouterGroup) {
	payment := user.Group("/payment")

	payment.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreatePaymentHandler)
	payment.POST("/failed", middlewares.UserTokenVerifyMiddleWare, handlers.FailedPaymentHandler)
	payment.POST("/success", middlewares.UserTokenVerifyMiddleWare, handlers.SuccessPaymentHandler)
	payment.POST("/cancel", middlewares.UserTokenVerifyMiddleWare, handlers.CancelPaymentHandler)

}
