package models

import "time"

type User struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Role        string `json:"role"`
}

type Product struct {
	ProductID   uint   `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	CategoryID  uint   `json:"category_id"`
	Image       string `json:"image"`
	Quantity    uint   `json:"quantity"`
}

type Category struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type UserProduct struct {
	UserProductID uint `json:"userproduct_id"`
	ProductID     uint `json:"product_id"`
	Quantity      uint `json:"quantity"`
	OrderID       uint `json:"order_id"`
}

type Address struct {
	AddressID uint   `json:"address_id"`
	UserID    uint   `json:"user_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	ZipCode   string `json:"zip_code"`
	IsDefault bool   `json:"is_default"`
}
type Shipping struct {
	ShippingID            uint   `json:"shipping_id"`
	UserID                uint   `json:"user_id"`
	Method                string `json:"method"` // "standard", "express", etc.
	AddressID             uint   `json:"address_id"`
	EstimatedDeliveryDays int    `json:"estimated_delivery_days"`
	// Other shipping-related fields
}

type PaymentMethod string

const (
	PaymentMethodCreditCard     PaymentMethod = "credit_card"
	PaymentMethodDebitCard      PaymentMethod = "debit_card"
	PaymentMethodNetBanking     PaymentMethod = "net_banking"
	PaymentMethodCashOnDelivery PaymentMethod = "cash_on_delivery"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentFailed    PaymentStatus = "failed"
	PaymentSuccess   PaymentStatus = "success"
	PaymentCancelled PaymentStatus = "cancelled"
)

type Payment struct {
	PaymentID     uint          `json:"payment_id"`
	UserID        uint          `json:"user_id"`
	OrderID       uint          `json:"order_id"`
	Amount        uint          `json:"amount"`
	Method        PaymentMethod `json:"method"`
	TransactionID string        `json:"transaction_id"`
	Status        PaymentStatus `json:"status"`
}

type Order struct {
	OrderID      uint      `json:"order_id"`
	UserID       uint      `json:"user_id"`
	TotalAmount  uint      `json:"total_amount"`
	PaymentID    uint      `json:"payment_id"`
	ShippingID   uint      `json:"shipping_id"`
	Status       string    `json:"status"` // "pending", "shipped", "delivered", etc.
	OrderDate    time.Time `json:"order_date"`
	DeliveryDate time.Time `json:"delivery_date"`
}

type OrderItem struct {
	OrderItemID uint `json:"orderitem_id"`
	OrderID     uint `json:"order_id"`
	ProductID   uint `json:"product_id"`
	Quantity    int  `json:"quantity"`
}

type CartItem struct {
	CartItemID uint `json:"cartitem_id"`
	UserID     uint `json:"user_id"`
	ProductID  uint `json:"productId"`
	Quantity   int  `json:"quantity"`
}

type WishlistItem struct {
	WishlistItemID uint `json:"wishlistitem_id"`
	UserID         uint `json:"user_id"`
	ProductID      uint `json:"productId"`
}

type Review struct {
	ReviewID  uint   `json:"review_id"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}
