package helpers

import (
	"fmt"
	"os"
	"regexp"
	"time"

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
		"name":         user.Name,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"role":         user.Role,
		"exp":          time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWTSECRET))
}

func GetClaims(c *gin.Context, tokenString string) (models.User, error) {
	JWTSECRET := os.Getenv("JWT_SECRET")
	var user models.User

	tokenString = tokenString[7:]

	verifiedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSECRET), nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return user, err
	}

	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid token claims")
		return user, fmt.Errorf("invalid token claims")
	}

	userId, userIdOk := claims["user_id"].(float64)
	if !userIdOk {
		fmt.Println("User Id not found in claims")
		return user, fmt.Errorf("user id not found in claims")
	}

	name, nameOk := claims["name"].(string)
	if !nameOk {
		fmt.Println("User Name not found in claims")
		return user, fmt.Errorf("user name not found in claims")
	}

	email, emailOk := claims["email"].(string)
	if !emailOk {
		fmt.Println("Email not found in claims")
		return user, fmt.Errorf("email not found in claims")
	}

	phoneNumber, phoneNumberOk := claims["phone_number"].(string)
	if !phoneNumberOk {
		fmt.Println("Phone Number not found in claims")
		return user, fmt.Errorf("phone number not found in claims")
	}

	role, roleOk := claims["role"].(string)
	if !roleOk {
		fmt.Println("Role not found in claims")
		return user, fmt.Errorf("role not found in claims")
	}

	user.UserID = uint(userId)
	user.Name = name
	user.Email = email
	user.PhoneNumber = phoneNumber
	user.Role = role

	return user, nil
}
