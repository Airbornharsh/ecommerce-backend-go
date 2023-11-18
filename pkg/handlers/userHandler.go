package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "User Found",
		"token":   token,
		"user":    user,
	})
}

func UpdateUserHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var newUser models.User

	err := c.BindJSON(&newUser)
	if helpers.ErrorResponse(c, err, 400) {
		return
	}

	if newUser.Name == "" && newUser.Email == "" && newUser.PhoneNumber == "" {
		c.JSON(200, gin.H{
			"message": "Nothing to Update",
		})
		return
	}

	if newUser.Name == "" {
		newUser.Name = user.Name
	}

	if newUser.Email == "" {
		newUser.Email = user.Email
	}

	if newUser.PhoneNumber == "" {
		newUser.PhoneNumber = user.PhoneNumber
	}

	q := `UPDATE users SET name = '` + newUser.Name + `', email = '` + newUser.Email + `', phone_number = '` + newUser.PhoneNumber + `' WHERE user_id = ` + strconv.Itoa(int(user.UserID)) + `;`

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	newUser.UserID = user.UserID
	newUser.Role = user.Role

	token, err := helpers.GenerateToken(&newUser)

	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "User Updated",
		"token":   token,
	})
}
