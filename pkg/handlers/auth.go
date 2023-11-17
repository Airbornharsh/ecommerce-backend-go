package handlers

import "github.com/gin-gonic/gin"

func Register(c *gin.Context){
	
	
	c.JSON(200, gin.H{
		"message": "Register",
	})
}

func Login(c *gin.Context){
}