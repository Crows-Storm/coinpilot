package services

import (
	"fmt"
	"time"

	"coinpilot/internal/filehandler"
	"coinpilot/internal/models"
)

// TradeService handles trade-related business logic
type TradeService interface {
	AddTrade(trade models.Trade) error
	ListTrades() ([]models.Trade, error)
	DeleteTrade(id string) error
}

// tradeService implements TradeService
type tradeService struct {
	csvHandler *filehandler.CSVHandler
}

// NewTradeService creates a new TradeService instance
func NewTradeService() TradeService {
	csvHandler := filehandler.NewCSVHandler()

	// Initialize trades file if it doesn't exist
	if err := csvHandler.InitializeTradesFile(); err != nil {
		// Log error but don't fail service creation
		fmt.Printf("Warning: failed to initialize trades.csv: %v\n", err)
	}

	return &tradeService{
		csvHandler: csvHandler,
	}
}

// AddTrade adds a new trade record
func (s *tradeService) AddTrade(trade models.Trade) error {
	// Validate trade type for MVP
	if !trade.Type.IsValidMVPTradeType() {
		return fmt.Errorf("unsupported trade type '%s' in MVP phase. Only BUY and SELL are supported", trade.Type)
	}

	// Set created timestamp if not provided
	if trade.CreatedAt.IsZero() {
		trade.CreatedAt = time.Now()
	}

	// Set trade timestamp if not provided
	if trade.Timestamp.IsZero() {
		trade.Timestamp = time.Now()
	}

	// Validate required fields
	if trade.ID == "" {
		return fmt.Errorf("trade ID is required")
	}
	if trade.Symbol == "" {
		return fmt.Errorf("symbol is required")
	}
	if trade.Quantity.IsZero() || trade.Quantity.IsNegative() {
		return fmt.Errorf("quantity must be positive")
	}
	if trade.Price.IsZero() || trade.Price.IsNegative() {
		return fmt.Errorf("price must be positive")
	}

	// Set default exchange if not provided
	if trade.Exchange == "" {
		trade.Exchange = "default"
	}

	// Set default pair if not provided
	if trade.Pair == "" {
		trade.Pair = trade.Symbol + "/USDT"
	}

	return s.csvHandler.AppendTrade(trade)
}

// ListTrades returns all trade records
func (s *tradeService) ListTrades() ([]models.Trade, error) {
	return s.csvHandler.LoadAllTrades()
}

// DeleteTrade removes a trade record by ID
func (s *tradeService) DeleteTrade(id string) error {
	if id == "" {
		return fmt.Errorf("trade ID is required")
	}

	return s.csvHandler.DeleteTradeByID(id)
}
