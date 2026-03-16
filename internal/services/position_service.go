package services

import (
	"fmt"
	"time"

	"coinpilot/internal/filehandler"
	"coinpilot/internal/models"

	"github.com/shopspring/decimal"
)

// PositionService handles position calculation business logic
type PositionService interface {
	CalculatePositions() ([]models.Position, error)
	GetPosition(symbol string) (*models.Position, error)
	RecalculateAll() error
}

// positionService implements PositionService
type positionService struct {
	csvHandler *filehandler.CSVHandler
}

// NewPositionService creates a new PositionService instance
func NewPositionService() PositionService {
	return &positionService{
		csvHandler: filehandler.NewCSVHandler(),
	}
}

// CalculatePositions calculates all current positions from trade history
func (s *positionService) CalculatePositions() ([]models.Position, error) {
	trades, err := s.csvHandler.LoadAllTrades()
	if err != nil {
		return nil, fmt.Errorf("failed to load trades: %w", err)
	}

	// Group trades by symbol and exchange
	positionMap := make(map[string]*models.Position)

	for _, trade := range trades {
		key := fmt.Sprintf("%s-%s", trade.Symbol, trade.Exchange)

		position, exists := positionMap[key]
		if !exists {
			position = &models.Position{
				Symbol:           trade.Symbol,
				Exchange:         trade.Exchange,
				Quantity:         decimal.Zero,
				AverageCost:      decimal.Zero,
				TotalCost:        decimal.Zero,
				CurrentPrice:     decimal.Zero,
				CurrentValue:     decimal.Zero,
				UnrealizedPnL:    decimal.Zero,
				RealizedPnL:      decimal.Zero,
				UnrealizedPnLPct: decimal.Zero,
				IsClosed:         false,
				LastUpdated:      time.Now(),
			}
			positionMap[key] = position
		}

		// Process trade based on type
		switch trade.Type {
		case models.BUY:
			position.UpdateAverageCost(trade.Quantity, trade.Price)
		case models.SELL:
			soldCost := position.ProcessSell(trade.Quantity)
			// For MVP, we don't track detailed realized PnL yet
			_ = soldCost
		}

		position.LastUpdated = time.Now()
	}

	// Convert map to slice, filtering out zero positions
	positions := make([]models.Position, 0, len(positionMap))
	for _, position := range positionMap {
		if !position.Quantity.IsZero() {
			positions = append(positions, *position)
		}
	}

	return positions, nil
}

// GetPosition returns position for a specific symbol
func (s *positionService) GetPosition(symbol string) (*models.Position, error) {
	positions, err := s.CalculatePositions()
	if err != nil {
		return nil, err
	}

	for _, position := range positions {
		if position.Symbol == symbol {
			return &position, nil
		}
	}

	return nil, fmt.Errorf("position for symbol %s not found", symbol)
}

// RecalculateAll recalculates all positions from trade history
func (s *positionService) RecalculateAll() error {
	// For MVP, this is the same as CalculatePositions since we don't cache positions
	_, err := s.CalculatePositions()
	return err
}
