package helpers

import (
	"errors"

	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
)

func OrderStatusConverter(status string) (models.OrderStatus, error) {
	switch status {
	case "pending":
		return models.OrderStatus("pending"), nil
	case "confirmed":
		return models.OrderStatus("confirmed"), nil
	case "shipped":
		return models.OrderStatus("shipped"), nil
	case "delivered":
		return models.OrderStatus("delivered"), nil
	case "cancelled":
		return models.OrderStatus("cancelled"), nil
	default:
		return models.OrderStatus("invalid"), errors.New("invalid order status")
	}
}
