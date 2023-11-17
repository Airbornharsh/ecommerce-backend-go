package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Api Server",
			})
		})

	AuthInit(api)
	ProductInit(api)
	CategoryInit(api)
	AddressInit(api)
}
