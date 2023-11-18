package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func AddressInit(user *gin.RouterGroup) {
	address := user.Group("/address")

	address.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetAllAddressHandler)
	address.GET("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.GetAddressHandler)
	address.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreateAddressHandler)
	address.PUT("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateAddressHandler)
	address.DELETE("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.DeleteAddressHandler)
}
