package main

import (
	"fmt"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/middlewares"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DBInit()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(middlewares.CorsMiddleware())

	r.GET("/", func(c *gin.Context) {
		fmt.Println("Hello World")
		c.JSON(200, gin.H{
			"message": "Welcome to Ecommerce Backend",
		})
	})

	r.POST("/", func(c *gin.Context) {
		fmt.Println(c)
		fmt.Println("Hello World POST")
		c.JSON(200, gin.H{
			"message": "Welcome to Ecommerce Backend POST",
		})
	})

	routes.Init(r)

	fmt.Println("Server Started at http://localhost:5000")
	r.Run(":5000")
}
