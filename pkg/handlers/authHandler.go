package handlers

import (
	"fmt"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	fmt.Println("Error:")
	var user models.User

	err := c.BindJSON(&user)
	if helpers.ErrorResponse(c, err, 400) {
		return
	}

	q := `SELECT EXISTS (SELECT 1 FROM users WHERE email = '` + user.Email + `' OR phone_number = '` + user.PhoneNumber + `');`

	var exists bool
	rows, err := database.DB.Query(q)

	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		err := rows.Scan(&exists)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		if exists {
			c.JSON(200, gin.H{
				"message": "User Already Exists",
			})
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	user.Password = string(hashedPassword)
	user.Role = "user"

	q = `INSERT INTO users (name, phone_number, password, email, role) VALUES ('` + user.Name + `', '` + user.PhoneNumber + `', '` + user.Password + `', '` + user.Email + `', '` + user.Role + `');`

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.Writer.Header().Set("Authorization", token)
	c.JSON(200, gin.H{
		"message": "User Registered Successfully",
		"token":   token,
		"userData": gin.H{
			"name":         user.Name,
			"phone_number": user.PhoneNumber,
			"email":        user.Email,
			"role":         user.Role,
		},
	})
}

func LoginHandler(c *gin.Context) {
	type Login struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}

	var tempUser Login
	var user models.User

	err := c.BindJSON(&tempUser)
	if helpers.ErrorResponse(c, err, 400) {
		return
	}

	q := `SELECT * FROM users WHERE name = '` + tempUser.Name + `' OR email = '` + tempUser.Email + `' OR phone_number = '` + tempUser.PhoneNumber + `';`

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Name, &user.PhoneNumber, &user.Password, &user.Email, &user.Role)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tempUser.Password))
		if helpers.ErrorResponse(c, err, 401) {
			return
		}

		token, err := helpers.GenerateToken(&user)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		c.Writer.Header().Set("Authorization", token)
		c.JSON(200, gin.H{
			"message": "User Logged In Successfully",
			"token":   token,
			"userData": gin.H{
				"name":         user.Name,
				"phone_number": user.PhoneNumber,
				"email":        user.Email,
				"role":         user.Role,
			},
		})
	}
}
