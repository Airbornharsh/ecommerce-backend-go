package main

import (
	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DBInit()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	routes.Init(r)

	r.Run(":8080")
}
