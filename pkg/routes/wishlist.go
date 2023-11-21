package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func WishlistInit(user *gin.RouterGroup) {
	wishlist := user.Group("/wishlist")

	wishlist.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetAllWishlistHandler)
	wishlist.GET("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.GetWishlistHandler)
	wishlist.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreateWishlistHandler)
	wishlist.PUT("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateWishlistHandler)
	wishlist.DELETE("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.DeleteWishlistHandler)

	wishlist.PUT("/:id/add/:productId", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateWishlistAddProductHandler)
	wishlist.PUT("/:id/remove/:productId", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateWishlistRemoveProductHandler)
}
