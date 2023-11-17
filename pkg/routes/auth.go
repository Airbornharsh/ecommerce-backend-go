package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthInit(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	//unAuth
	auth.POST("/register", handlers.RegisterHandler)
	auth.POST("/login", handlers.LoginHandler)

	//Auth
	auth.GET("/user", middlewares.UserTokenVerifyMiddleWare, handlers.GetUserHandler)
	auth.PUT("/user", handlers.UpdateUserHandler)
}
