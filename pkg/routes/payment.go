package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func PaymentInit(user *gin.RouterGroup) {
	payment := user.Group("/payment")

	payment.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreatePaymentHandler)
	payment.PUT("/failed/:paymentId", middlewares.UserTokenVerifyMiddleWare, handlers.FailedPaymentHandler)
	payment.PUT("/success/:paymentId", middlewares.UserTokenVerifyMiddleWare, handlers.SuccessPaymentHandler)
	payment.PUT("/cancel/:paymentId", middlewares.UserTokenVerifyMiddleWare, handlers.CancelPaymentHandler)

	payment.GET("/status", handlers.ListPaymentStatusHandler)
	payment.GET("/method", handlers.ListPaymentMethodHandler)
}
