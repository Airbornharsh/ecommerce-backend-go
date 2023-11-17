package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AuthInit(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	//unAuth
	auth.POST("/register", handlers.RegisterHandler)
	auth.POST("/login", handlers.LoginHandler)

	//Auth
	auth.GET("/user", handlers.GetUserHandler)
	auth.PUT("/user", handlers.UpdateUserHandler)
}
