package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsValidPassword(password string) bool {
	passwordPatternRegex := "[a-zA-Z0-9!@#$%^&*()_+{}|:<>?]{8,}"

	regexpObj := regexp.MustCompile(passwordPatternRegex)

	return regexpObj.MatchString(password)
}

func GenerateToken(user *models.User) (string, error) {
	JWTSECRET := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id":      user.UserID,
		"username":     user.Username,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"role":         user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWTSECRET))
}

func GetClaims(c *gin.Context, tokenString string) (string, error) {
	JWTSECRET := os.Getenv("JWT_SECRET")

	verifiedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSECRET), nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims")
		return "", fmt.Errorf("invalid token claims")
	}

	userId, userIdOk := claims["user_id"].(float64)
	if !userIdOk {
		fmt.Println("User Id not found in claims")
		return "", fmt.Errorf("user id not found in claims")
	}

	return strconv.FormatFloat(userId, 'f', -1, 64), nil
}
