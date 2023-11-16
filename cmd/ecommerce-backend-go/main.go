package main

import (
	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/routes"
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

	routes.Init(r)

	r.Run(":8080")
}
