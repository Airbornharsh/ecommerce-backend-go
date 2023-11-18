package routes

import (
	"github.com/airbornharsh/ecommerce-backend-go/pkg/handlers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func ReviewInit(user *gin.RouterGroup) {
	review := user.Group("/review")

	review.GET("/", middlewares.UserTokenVerifyMiddleWare, handlers.GetAllReviewHandler)
	review.GET("/product/:productId", middlewares.UserTokenVerifyMiddleWare, handlers.GetAllProductReviewHandler)
	review.GET("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.GetReviewHandler)
	review.POST("/", middlewares.UserTokenVerifyMiddleWare, handlers.CreateReviewHandler)
	review.PUT("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.UpdateReviewHandler)
	review.DELETE("/:id", middlewares.UserTokenVerifyMiddleWare, handlers.DeleteReviewHandler)
}
