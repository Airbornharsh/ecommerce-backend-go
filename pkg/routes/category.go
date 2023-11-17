package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func CategoryInit(r *gin.RouterGroup) {
	category := r.Group("/category")

	//User
	category.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetAllCategoryHandler)
	category.GET("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.GetCategoryHandler)

	//Admin
	category.POST("/", middlewares.AdminTokenVerifyMiddleWare, handlers.CreateCategoryHandler)
	category.PUT("/:id", middlewares.AdminTokenVerifyMiddleWare, handlers.UpdateCategoryHandler)
	category.DELETE("/:id", middlewares.AdminTokenVerifyMiddleWare, handlers.DeleteCategoryHandler)
}
