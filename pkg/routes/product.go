package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func ProductInit(r *gin.RouterGroup) {
	products := r.Group("/products")

	//user && UnAuth
	products.GET("/", handlers.GetProductsHandler)
	products.GET("/:id", handlers.GetProductHandler)
	products.GET("/filter/:category", handlers.FilterCategoryHandler)

	//admin
	products.POST("/", handlers.PostProductsHandler)
	products.PUT("/:id", handlers.PutProductsHandler)
	products.DELETE("/:id", handlers.DeleteProductsHandler)
}
