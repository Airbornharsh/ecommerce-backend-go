package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func ProductInit(r *gin.RouterGroup) {
	products := r.Group("/products")

	//user && UnAuth
	products.GET("/", handlers.GetProducts)
	products.GET("/:id", handlers.GetProduct)
	products.GET("/filter/:category", handlers.FilterCategory)

	//admin
	products.POST("/", handlers.PostProducts)
	products.PUT("/:id", handlers.PutProducts)
	products.DELETE("/:id", handlers.DeleteProducts)
}
