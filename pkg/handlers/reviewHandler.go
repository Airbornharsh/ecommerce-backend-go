package handlers

import (
	"strconv"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllReviewHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	q := "SELECT * FROM reviews WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "'"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var reviews []models.Review

	for rows.Next() {
		var review models.Review
		err = rows.Scan(&review.ReviewID, &review.UserID, &review.ProductID, &review.Rating, &review.Comment)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}
		reviews = append(reviews, review)
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Reviews fetched successfully",
		"token":   token,
		"reviews": reviews,
	})
}

func GetAllProductReviewHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	productID := c.Param("productId")

	q := "SELECT * FROM reviews WHERE product_id = '" + productID + "'"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var reviews []models.Review

	for rows.Next() {
		var review models.Review
		err = rows.Scan(&review.ReviewID, &review.UserID, &review.ProductID, &review.Rating, &review.Comment)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}
		reviews = append(reviews, review)
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Reviews fetched successfully",
		"token":   token,
		"reviews": reviews,
	})
}

func GetReviewHandler(c *gin.Context) {
}

func CreateReviewHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var review models.Review

	if err := c.ShouldBindJSON(&review); helpers.ErrorResponse(c, err, 400) {
		return
	}

	review.UserID = user.UserID

	q := "INSERT INTO reviews (user_id, product_id, rating, comment) VALUES ('" + strconv.Itoa(int(user.UserID)) + "', '" + strconv.Itoa(int(review.ProductID)) + "', '" + strconv.Itoa(review.Rating) + "', '" + review.Comment + "') RETURNING review_id"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var reviewID uint

	for rows.Next() {
		err = rows.Scan(&reviewID)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}
	}

	review.ReviewID = reviewID

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Review created successfully",
		"token":   token,
		"review":  review,
	})
}

func UpdateReviewHandler(c *gin.Context) {
}

func DeleteReviewHandler(c *gin.Context) {
}
