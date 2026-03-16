package utils

import (
	"errors"
	"strings"

	"coinpilot/internal/models"

	"github.com/shopspring/decimal"
)

// ValidateTradeParameters validates required trade parameters
func ValidateTradeParameters(tradeType models.TradeType, symbol string, quantity, price decimal.Decimal) error {
	if tradeType == "" {
		return errors.New("trade type is required")
	}

	if strings.TrimSpace(symbol) == "" {
		return errors.New("symbol is required")
	}

	if quantity.IsZero() || quantity.IsNegative() {
		return errors.New("quantity must be positive")
	}

	if price.IsZero() || price.IsNegative() {
		return errors.New("price must be positive")
	}

	// Validate trade type
	validTypes := map[models.TradeType]bool{
		models.BUY:          true,
		models.SELL:         true,
		models.TRANSFER_IN:  true,
		models.TRANSFER_OUT: true,
		models.FEE:          true,
	}

	if !validTypes[tradeType] {
		return errors.New("invalid trade type")
	}

	return nil
}
