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

	// q := "SELECT * FROM reviews WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "'"

	q := "SELECT reviews.review_id, reviews.user_id, reviews.product_id, products.name, products.description, products.price, products.category_id, categories.name, products.image, products.quantity, reviews.rating, reviews.comment FROM reviews INNER JOIN products ON reviews.product_id = products.product_id INNER JOIN categories ON products.category_id = categories.category_id WHERE reviews.user_id = '" + strconv.Itoa(int(user.UserID)) + "'"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	type Review struct {
		ReviewID           uint   `json:"review_id"`
		UserID             uint   `json:"user_id"`
		ProductID          uint   `json:"product_id"`
		ProductName        string `json:"product_name"`
		ProductDescription string `json:"product_description"`
		ProductPrice       uint   `json:"product_price"`
		ProductCategoryID  uint   `json:"product_category_id"`
		ProductCategory    string `json:"product_category"`
		ProductImage       string `json:"product_image"`
		ProductQuantity    uint   `json:"product_quantity"`
		Rating             uint   `json:"rating"`
		Comment            string `json:"comment"`
	}

	var reviews []Review

	for rows.Next() {
		var review Review
		err = rows.Scan(&review.ReviewID, &review.UserID, &review.ProductID, &review.ProductName, &review.ProductDescription, &review.ProductPrice, &review.ProductCategoryID, &review.ProductCategory, &review.ProductImage, &review.ProductQuantity, &review.Rating, &review.Comment)
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
	user := c.MustGet("user").(models.User)

	reviewID := c.Param("id")

	q := "SELECT * FROM reviews WHERE review_id = '" + reviewID + "'"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var review models.Review

	for rows.Next() {
		err = rows.Scan(&review.ReviewID, &review.UserID, &review.ProductID, &review.Rating, &review.Comment)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Review fetched successfully",
		"token":   token,
		"review":  review,
	})
}

func CreateReviewHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var review models.Review

	if err := c.ShouldBindJSON(&review); helpers.ErrorResponse(c, err, 400) {
		return
	}

	review.UserID = user.UserID

	var isReviewed bool

	q := "SELECT EXISTS (SELECT 1 FROM reviews WHERE user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND product_id = '" + strconv.Itoa(int(review.ProductID)) + "')"

	err := database.DB.QueryRow(q).Scan(&isReviewed)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if isReviewed {
		c.JSON(400, gin.H{
			"message": "You have already reviewed this product",
		})
		return
	}

	q = "INSERT INTO reviews (user_id, product_id, rating, comment) VALUES ('" + strconv.Itoa(int(user.UserID)) + "', '" + strconv.Itoa(int(review.ProductID)) + "', '" + strconv.Itoa(review.Rating) + "', '" + review.Comment + "') RETURNING review_id"

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
	user := c.MustGet("user").(models.User)

	reviewID := c.Param("id")

	var review models.Review

	if err := c.ShouldBindJSON(&review); helpers.ErrorResponse(c, err, 400) {
		return
	}

	review.UserID = user.UserID

	q := "UPDATE reviews SET "
	if review.Rating != 0 {
		q += "rating = '" + strconv.Itoa(review.Rating) + "', "
	}
	if review.Comment != "" {
		q += "comment = '" + review.Comment + "', "
	}
	q += "updated_at = NOW() WHERE review_id = '" + reviewID + "' RETURNING review_id"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var updatedReviewID uint

	for rows.Next() {
		err = rows.Scan(&updatedReviewID)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}
	}

	if updatedReviewID == 0 {
		c.JSON(400, gin.H{
			"message": "Review not found",
		})
		return
	}
}

func DeleteReviewHandler(c *gin.Context) {
}
