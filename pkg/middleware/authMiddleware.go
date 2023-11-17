package middleware

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func UserTokenVerifyMiddleWare(c *gin.Context) {
	Auth := c.Request.Header.Get("Authorization")

	if Auth == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Set("user", nil)

		c.Abort()
		return
	}

	tempUser, err := helpers.GetClaims(c, Auth)
	if helpers.ErrorResponse(c, err, 401) {
		c.Abort()
		return
	}

	var user models.User

	q := `SELECT user_id, username, email, phone_number, role FROM users WHERE user_id = ` + strconv.Itoa(int(tempUser.UserID)) + `;`

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		c.Abort()
		return
	}

	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.PhoneNumber, &user.Role)
		if helpers.ErrorResponse(c, err, 500) {
			c.Abort()
			return
		}
	}

	if user.UserID == 0 {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Set("user", nil)

		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}
