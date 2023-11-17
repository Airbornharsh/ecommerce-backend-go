package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func ProductInit(r *gin.RouterGroup) {
	products := r.Group("/products")

	//user && UnAuth
	products.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetProductsHandler)
	products.GET("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.GetProductHandler)
	// products.GET("/filter/:category", handlers.FilterCategoryHandler)

	//admin
	products.POST("/", middlewares.AdminTokenVerifyMiddleWare, handlers.PostProductsHandler)
	products.PUT("/:id", middlewares.AdminTokenVerifyMiddleWare, handlers.PutProductsHandler)
	products.DELETE("/:id", middlewares.AdminTokenVerifyMiddleWare, handlers.DeleteProductsHandler)
}
