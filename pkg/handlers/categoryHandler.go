package handlers

import "github.com/gin-gonic/gin"

func GetAllCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetAllCategoryHandler",
	})
}

func GetCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetCategoryHandler",
	})
}

func CreateCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateCategoryHandler",
	})
}

func UpdateCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateCategoryHandler",
	})
}

func DeleteCategoryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteCategoryHandler",
	})
}
