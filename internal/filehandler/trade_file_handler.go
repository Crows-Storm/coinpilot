package filehandler

import (
	"coinpilot/internal/models"
	"fmt"
	"strings"
)

// TradeFileHandler provides basic CRUD operations for trade records
// This is the MVP Phase 1 interface that wraps the CSVHandler
type TradeFileHandler struct {
	csvHandler *CSVHandler
}

// NewTradeFileHandler creates a new TradeFileHandler instance
func NewTradeFileHandler() *TradeFileHandler {
	return &TradeFileHandler{
		csvHandler: NewCSVHandler(),
	}
}

// Initialize sets up the trades.csv file if it doesn't exist
func (h *TradeFileHandler) Initialize() error {
	return h.csvHandler.InitializeTradesFile()
}

// Create adds a new trade record to the CSV file
// Performs basic validation for required fields only (MVP Phase 1)
func (h *TradeFileHandler) Create(trade models.Trade) error {
	// Basic validation for required fields
	if err := h.validateRequiredFields(trade); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// MVP Phase 1: Only support BUY and SELL trade types
	if !trade.Type.IsValidMVPTradeType() {
		return fmt.Errorf("unsupported trade type '%s' in MVP phase, only BUY and SELL are supported", trade.Type)
	}

	// Use append for single trade creation (more efficient than rewriting entire file)
	return h.csvHandler.AppendTrade(trade)
}

// FindAll retrieves all trade records from the CSV file
func (h *TradeFileHandler) FindAll() ([]models.Trade, error) {
	return h.csvHandler.LoadAllTrades()
}

// Delete removes a trade record by ID
func (h *TradeFileHandler) Delete(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("trade ID cannot be empty")
	}

	return h.csvHandler.DeleteTradeByID(id)
}

// validateRequiredFields performs basic validation for required fields
func (h *TradeFileHandler) validateRequiredFields(trade models.Trade) error {
	var errors []string

	// Validate ID
	if strings.TrimSpace(trade.ID) == "" {
		errors = append(errors, "ID is required")
	}

	// Validate Symbol
	if strings.TrimSpace(trade.Symbol) == "" {
		errors = append(errors, "Symbol is required")
	}

	// Validate Type
	if strings.TrimSpace(string(trade.Type)) == "" {
		errors = append(errors, "Type is required")
	}

	// Validate Quantity (must be positive)
	if trade.Quantity.IsZero() || trade.Quantity.IsNegative() {
		errors = append(errors, "Quantity must be positive")
	}

	// Validate Price (must be positive)
	if trade.Price.IsZero() || trade.Price.IsNegative() {
		errors = append(errors, "Price must be positive")
	}

	// Validate Timestamp (cannot be zero time)
	if trade.Timestamp.IsZero() {
		errors = append(errors, "Timestamp is required")
	}

	// Validate CreatedAt (cannot be zero time)
	if trade.CreatedAt.IsZero() {
		errors = append(errors, "CreatedAt is required")
	}

	// Fee can be zero or positive (negative fees not allowed)
	if trade.Fee.IsNegative() {
		errors = append(errors, "Fee cannot be negative")
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errors, ", "))
	}

	return nil
}
