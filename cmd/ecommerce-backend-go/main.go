package main

import (
	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	database.DBInit()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.Run(":8080")
}
