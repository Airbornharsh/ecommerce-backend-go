package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func ShippingInit(user *gin.RouterGroup) {
	shipping := user.Group("/shipping")

	shipping.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreateShippingHandler)
}
