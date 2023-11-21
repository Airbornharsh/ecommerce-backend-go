package handlers

import (
	"strconv"
	"time"

	"github.com/airbornharsh/ecommerce-backend-go/internal/database"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/helpers"
	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
	"github.com/gin-gonic/gin"
)

type OrderProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type Order struct {
	OrderID      uint               `json:"order_id"`
	UserID       uint               `json:"user_id"`
	TotalAmount  uint               `json:"total_amount"`
	PaymentID    uint               `json:"payment_id"`
	Products     []OrderProduct     `json:"products"`
	ShippingID   uint               `json:"shipping_id"`
	Status       string             `json:"status"`
	OrderStatus  models.OrderStatus `json:"order_status"`
	OrderDate    time.Time          `json:"order_date"`
	DeliveryDate time.Time          `json:"delivery_date"`
}

type TempOrderProduct struct {
	ProductID   uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Description string `json:"product_description"`
	Price       uint   `json:"product_price"`
	CategoryID  uint   `json:"product_category_id"`
	Category    string `json:"product_category"`
	Image       string `json:"product_image"`
	Quantity    int    `json:"product_quantity"`
}

type TempOrder struct {
	OrderID               uint               `json:"order_id"`
	UserID                uint               `json:"user_id"`
	TotalAmount           uint               `json:"order_total_amount"`
	PaymentID             uint               `json:"payment_id"`
	Amount                uint               `json:"payment_amount"`
	PaymentMethod         string             `json:"payment_method"`
	Products              []TempOrderProduct `json:"products"`
	PaymentStatus         string             `json:"payment_status"`
	CreatedAt             time.Time          `json:"payment_created_at"`
	ShippingID            uint               `json:"shipping_id"`
	ShippingMethod        string             `json:"shipping_method"`
	Street                string             `json:"shipping_street"`
	City                  string             `json:"shipping_city"`
	State                 string             `json:"shipping_state"`
	Country               string             `json:"shipping_country"`
	ZipCode               string             `json:"shipping_zipcode"`
	IsDefault             bool               `json:"shipping_is_default"`
	EstimatedDeliveryDays int                `json:"shipping_estimated_delivery_days"`
	OrderStatus           string             `json:"order_status"`
	OrderDate             time.Time          `json:"order_date"`
	DeliveryDate          time.Time          `json:"order_delivery_date"`
}

func GetAllOrderHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var orders []TempOrder

	// q := "SELECT order.order_id, order.user_id, order.total_amount, order.payment_id, "

	q := "SELECT orders.order_id, orders.user_id, orders.total_amount, orders.payment_id, payments.amount, payments.method, payments.status, payments.created_at, orders.shipping_id, shippings.method, addresses.street, addresses.city, addresses.state, addresses.country, addresses.zip_code, addresses.is_default, shippings.estimated_delivery_days, orders.status, orders.order_date, orders.delivery_date FROM orders INNER JOIN payments ON orders.payment_id = payments.payment_id INNER JOIN shippings ON orders.shipping_id = shippings.shipping_id INNER JOIN addresses ON shippings.address_id = addresses.address_id WHERE orders.user_id = '" + strconv.Itoa(int(user.UserID)) + "'"

	rows, err := database.DB.Query(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		var order TempOrder

		err := rows.Scan(&order.OrderID, &order.UserID, &order.TotalAmount, &order.PaymentID, &order.Amount, &order.PaymentMethod, &order.PaymentStatus, &order.CreatedAt, &order.ShippingID, &order.ShippingMethod, &order.Street, &order.City, &order.State, &order.Country, &order.ZipCode, &order.IsDefault, &order.EstimatedDeliveryDays, &order.OrderStatus, &order.OrderDate, &order.DeliveryDate)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		var products []TempOrderProduct

		q = "SELECT orderitems.product_id, products.name, products.description, products.price, products.category_id, categories.name, products.image, orderitems.quantity FROM orderitems INNER JOIN products ON orderitems.product_id = products.product_id INNER JOIN categories ON products.category_id = categories.category_id WHERE orderitems.order_id = '" + strconv.Itoa(int(order.OrderID)) + "'"

		row, err := database.DB.Query(q)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		for row.Next() {
			var product TempOrderProduct

			err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Category, &product.Image, &product.Quantity)

			if helpers.ErrorResponse(c, err, 500) {
				return
			}
			products = append(products, product)
		}

		order.Products = products
		orders = append(orders, order)
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "All Orders",
		"token":   token,
		"orders":  orders,
	})
}

func GetOrderHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	orderID := c.Param("id")

	var order TempOrder

	q := "SELECT orders.order_id, orders.user_id, orders.total_amount, orders.payment_id, payments.amount, payments.method, payments.status, payments.created_at, orders.shipping_id, shippings.method, addresses.street, addresses.city, addresses.state, addresses.country, addresses.zip_code, addresses.is_default, shippings.estimated_delivery_days, orders.status, orders.order_date, orders.delivery_date FROM orders INNER JOIN payments ON orders.payment_id = payments.payment_id INNER JOIN shippings ON orders.shipping_id = shippings.shipping_id INNER JOIN addresses ON shippings.address_id = addresses.address_id WHERE orders.user_id = '" + strconv.Itoa(int(user.UserID)) + "' AND orders.order_id = '" + orderID + "'"

	err := database.DB.QueryRow(q).Scan(&order.OrderID, &order.UserID, &order.TotalAmount, &order.PaymentID, &order.Amount, &order.PaymentMethod, &order.PaymentStatus, &order.CreatedAt, &order.ShippingID, &order.ShippingMethod, &order.Street, &order.City, &order.State, &order.Country, &order.ZipCode, &order.IsDefault, &order.EstimatedDeliveryDays, &order.OrderStatus, &order.OrderDate, &order.DeliveryDate)

	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var products []TempOrderProduct

	q = "SELECT orderitems.product_id, products.name, products.description, products.price, products.category_id, categories.name, products.image, orderitems.quantity FROM orderitems INNER JOIN products ON orderitems.product_id = products.product_id INNER JOIN categories ON products.category_id = categories.category_id WHERE orderitems.order_id = '" + orderID + "'"

	rows, err := database.DB.Query(q)

	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	for rows.Next() {
		var product TempOrderProduct

		err := rows.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.Category, &product.Image, &product.Quantity)

		if helpers.ErrorResponse(c, err, 500) {
			return
		}
		products = append(products, product)
	}

	order.Products = products

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Order",
		"token":   token,
		"order":   order,
	})
}

func CreateOrderHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var order Order

	if err := c.ShouldBindJSON(&order); helpers.ErrorResponse(c, err, 400) {
		return
	}

	var shipping models.Shipping

	q := "SELECT estimated_delivery_days FROM shippings WHERE shipping_id = '" + strconv.Itoa(int(order.ShippingID)) + "'"

	err := database.DB.QueryRow(q).Scan(&shipping.EstimatedDeliveryDays)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	order.UserID = user.UserID
	order.OrderDate = time.Now()
	order.DeliveryDate = time.Now().AddDate(0, 0, shipping.EstimatedDeliveryDays)

	q = "INSERT INTO orders (user_id, total_amount, payment_id, shipping_id, status, order_date, delivery_date) VALUES ('" + strconv.Itoa(int(order.UserID)) + "', '0', '" + strconv.Itoa(int(order.PaymentID)) + "', '" + strconv.Itoa(int(order.ShippingID)) + "', 'pending', '" + order.OrderDate.Format("2006-01-02 15:04:05") + "', '" + order.DeliveryDate.Format("2006-01-02 15:04:05") + "') RETURNING order_id"

	err = database.DB.QueryRow(q).Scan(&order.OrderID)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q = "INSERT INTO orderitems (order_id, product_id, quantity) VALUES "
	for i, product := range order.Products {
		q += "('" + strconv.Itoa(int(order.OrderID)) + "', '" + strconv.Itoa(int(product.ProductID)) + "', '" + strconv.Itoa(product.Quantity) + "')"
		if i == len(order.Products)-1 {
			q += ";"
		} else {
			q += ", "
		}
	}

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	q = "UPDATE orders SET total_amount = (SELECT SUM(products.price * orderitems.quantity) FROM orderitems INNER JOIN products ON orderitems.product_id = products.product_id WHERE orderitems.order_id = '" + strconv.Itoa(int(order.OrderID)) + "') WHERE order_id = '" + strconv.Itoa(int(order.OrderID)) + "' RETURNING total_amount"

	var totalAmount uint

	err = database.DB.QueryRow(q).Scan(&totalAmount)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	var payment models.Payment

	q = "SELECT amount FROM payments WHERE payment_id = '" + strconv.Itoa(int(order.PaymentID)) + "'"

	err = database.DB.QueryRow(q).Scan(&payment.Amount)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	if totalAmount > payment.Amount {
		q = "UPDATE orders SET status = 'cancelled' WHERE order_id = '" + strconv.Itoa(int(order.OrderID)) + "' "

		_, err = database.DB.Exec(q)
		if helpers.ErrorResponse(c, err, 500) {
			return
		}

		c.JSON(400, gin.H{
			"message": "Payment amount does not match with order amount",
		})
		return
	}

	q = "UPDATE orders SET status = 'confirmed' WHERE order_id = '" + strconv.Itoa(int(order.OrderID)) + "' "

	_, err = database.DB.Exec(q)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	token, err := helpers.GenerateToken(&user)
	if helpers.ErrorResponse(c, err, 500) {
		return
	}

	c.JSON(200, gin.H{
		"message": "Order created successfully",
		"token":   token,
		"order":   order,
	})
}

func UpdateOrderHandler(c *gin.Context) {

}

func DeleteOrderHandler(c *gin.Context) {

}
