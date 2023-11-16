package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AuthInit(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	//unAuth
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)

	//Auth
	auth.GET("/user", handlers.GetUser)
	auth.PUT("/user", handlers.UpdateUser)
}
