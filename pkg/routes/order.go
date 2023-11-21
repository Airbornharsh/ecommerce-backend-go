package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderInit(user *gin.RouterGroup) {
	order := user.Group("/order")

	order.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetAllOrderHandler)
	order.GET("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.GetOrderHandler)
	order.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreateOrderHandler)
	order.PUT("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateOrderHandler)
	order.DELETE("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.DeleteOrderHandler)
}
