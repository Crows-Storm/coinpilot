package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// TradeType represents the type of trade operation
type TradeType string

const (
	// MVP Phase 1: Only BUY and SELL supported
	BUY  TradeType = "BUY"
	SELL TradeType = "SELL"

	// Phase 2: Additional trade types
	TRANSFER_IN  TradeType = "TRANSFER_IN"
	TRANSFER_OUT TradeType = "TRANSFER_OUT"
	FEE          TradeType = "FEE"
)

// IsValidMVPTradeType checks if the trade type is supported in MVP
func (t TradeType) IsValidMVPTradeType() bool {
	return t == BUY || t == SELL
}

// Trade represents a cryptocurrency trade record
type Trade struct {
	ID        string          `json:"id" csv:"id"`
	Timestamp time.Time       `json:"timestamp" csv:"timestamp"`
	Exchange  string          `json:"exchange" csv:"exchange"`
	Pair      string          `json:"pair" csv:"pair"`
	Symbol    string          `json:"symbol" csv:"symbol"`
	Type      TradeType       `json:"type" csv:"type"`
	Quantity  decimal.Decimal `json:"quantity" csv:"quantity"`
	Price     decimal.Decimal `json:"price" csv:"price"`
	Fee       decimal.Decimal `json:"fee" csv:"fee"`
	Notes     string          `json:"notes" csv:"notes"`
	CreatedAt time.Time       `json:"created_at" csv:"created_at"`
}
