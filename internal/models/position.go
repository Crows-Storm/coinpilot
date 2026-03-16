package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Position represents a calculated position for a specific symbol
type Position struct {
	Symbol           string          `json:"symbol" csv:"symbol"`
	Exchange         string          `json:"exchange" csv:"exchange"`
	Quantity         decimal.Decimal `json:"quantity" csv:"quantity"`
	AverageCost      decimal.Decimal `json:"average_cost" csv:"average_cost"`
	TotalCost        decimal.Decimal `json:"total_cost" csv:"total_cost"`
	CurrentPrice     decimal.Decimal `json:"current_price" csv:"current_price"`
	CurrentValue     decimal.Decimal `json:"current_value" csv:"current_value"`
	UnrealizedPnL    decimal.Decimal `json:"unrealized_pnl" csv:"unrealized_pnl"`
	RealizedPnL      decimal.Decimal `json:"realized_pnl" csv:"realized_pnl"`
	UnrealizedPnLPct decimal.Decimal `json:"unrealized_pnl_pct" csv:"unrealized_pnl_pct"`
	IsClosed         bool            `json:"is_closed" csv:"is_closed"`
	LastUpdated      time.Time       `json:"last_updated" csv:"last_updated"`
}

// CalculateCurrentValue calculates current value based on quantity and current price
func (p *Position) CalculateCurrentValue() {
	p.CurrentValue = p.Quantity.Mul(p.CurrentPrice)
}

// CalculateUnrealizedPnL calculates unrealized profit/loss
func (p *Position) CalculateUnrealizedPnL() {
	p.UnrealizedPnL = p.CurrentValue.Sub(p.TotalCost)

	// Calculate percentage if total cost is not zero
	if !p.TotalCost.IsZero() {
		p.UnrealizedPnLPct = p.UnrealizedPnL.Div(p.TotalCost).Mul(decimal.NewFromInt(100))
	} else {
		p.UnrealizedPnLPct = decimal.Zero
	}
}

// UpdateAverageCost updates average cost when adding a buy trade
func (p *Position) UpdateAverageCost(newQuantity, newPrice decimal.Decimal) {
	if newQuantity.IsZero() {
		return
	}

	// Calculate new total cost and quantity
	newTotalCost := newQuantity.Mul(newPrice)
	totalQuantity := p.Quantity.Add(newQuantity)

	if totalQuantity.IsZero() {
		p.AverageCost = decimal.Zero
		p.TotalCost = decimal.Zero
	} else {
		p.TotalCost = p.TotalCost.Add(newTotalCost)
		p.AverageCost = p.TotalCost.Div(totalQuantity)
	}

	p.Quantity = totalQuantity
}

// ProcessSell processes a sell trade and updates position
func (p *Position) ProcessSell(sellQuantity decimal.Decimal) decimal.Decimal {
	if sellQuantity.IsZero() || p.Quantity.IsZero() {
		return decimal.Zero
	}

	// Calculate realized PnL for the sold portion
	soldCost := sellQuantity.Mul(p.AverageCost)

	// Update position
	p.Quantity = p.Quantity.Sub(sellQuantity)
	p.TotalCost = p.TotalCost.Sub(soldCost)

	// Mark as closed if quantity is zero
	if p.Quantity.IsZero() {
		p.IsClosed = true
		p.AverageCost = decimal.Zero
		p.TotalCost = decimal.Zero
	}

	return soldCost
}
