package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func CartInit(user *gin.RouterGroup) {
	cart := user.Group("/cart")

	cart.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetCartHandler)
	cart.POST("/add/:productId", middlewares.UserTokenVerifyMiddleWare, handlers.AddProductCartHandler)
	cart.PUT("/update/:productId", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateProductCartHandler)
	cart.DELETE("/remove/:productId", middlewares.UserTokenVerifyMiddleWare, handlers.DeleteProductCartHandler)
	cart.DELETE("/clear", middlewares.UserTokenVerifyMiddleWare, handlers.DeleteCartHandler)
}
