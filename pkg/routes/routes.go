package routes

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	api := r.Group("/api")

	AuthInit(api)
	ProductInit(api)

	api.GET("/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World from Go!",
			})
		})
}
