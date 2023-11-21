package helpers

import (
	"errors"

	"github.com/airbornharsh/ecommerce-backend-go/pkg/models"
)

func PaymentMethodConverter(status string) (models.PaymentMethod, error) {
	switch status {
	case "credit_card":
		return models.PaymentMethod("credit_card"), nil
	case "debit_card":
		return models.PaymentMethod("debit_card"), nil
	case "net_banking":
		return models.PaymentMethod("net_banking"), nil
	case "cash_on_delivery":
		return models.PaymentMethod("cash_on_delivery"), nil
	case "upi":
		return models.PaymentMethod("upi"), nil
	default:
		return models.PaymentMethod("invalid"), errors.New("invalid payment method")
	}
}
