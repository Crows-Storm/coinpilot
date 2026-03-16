package filehandler

import (
	"coinpilot/internal/models"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

// CSVHandler handles CSV file operations for trades
type CSVHandler struct {
	dataDir string
	mutex   sync.RWMutex
}

// NewCSVHandler creates a new CSV handler instance
func NewCSVHandler() *CSVHandler {
	// Get current working directory (script directory)
	wd, err := os.Getwd()
	if err != nil {
		// Fallback to current directory
		wd = "."
	}

	return &CSVHandler{
		dataDir: wd,
	}
}

// GetTradesFilePath returns the full path to trades.csv
func (h *CSVHandler) GetTradesFilePath() string {
	return filepath.Join(h.dataDir, "trades.csv")
}

// InitializeTradesFile creates trades.csv with headers if it doesn't exist
func (h *CSVHandler) InitializeTradesFile() error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	filePath := h.GetTradesFilePath()

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		// File exists, validate headers
		return h.validateTradesFileHeaders()
	}

	// Create new file with headers
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create trades.csv: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{
		"id", "timestamp", "exchange", "pair", "symbol",
		"type", "quantity", "price", "fee", "notes", "created_at",
	}

	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	return nil
}

// validateTradesFileHeaders checks if the CSV file has correct headers
func (h *CSVHandler) validateTradesFileHeaders() error {
	file, err := os.Open(h.GetTradesFilePath())
	if err != nil {
		return fmt.Errorf("failed to open trades.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV headers: %w", err)
	}

	expectedHeaders := []string{
		"id", "timestamp", "exchange", "pair", "symbol",
		"type", "quantity", "price", "fee", "notes", "created_at",
	}

	if len(headers) != len(expectedHeaders) {
		return fmt.Errorf("invalid CSV headers: expected %d columns, got %d", len(expectedHeaders), len(headers))
	}

	for i, expected := range expectedHeaders {
		if headers[i] != expected {
			return fmt.Errorf("invalid CSV header at position %d: expected '%s', got '%s'", i, expected, headers[i])
		}
	}

	return nil
}

// LoadAllTrades reads all trades from the CSV file
func (h *CSVHandler) LoadAllTrades() ([]models.Trade, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	filePath := h.GetTradesFilePath()

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File doesn't exist, return empty slice
		return []models.Trade{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open trades.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV records: %w", err)
	}

	// Skip header row
	if len(records) <= 1 {
		return []models.Trade{}, nil
	}

	trades := make([]models.Trade, 0, len(records)-1)

	for i, record := range records[1:] {
		trade, err := h.parseTradeRecord(record)
		if err != nil {
			return nil, fmt.Errorf("failed to parse trade record at line %d: %w", i+2, err)
		}
		trades = append(trades, trade)
	}

	return trades, nil
}

// SaveAllTrades writes all trades to the CSV file
func (h *CSVHandler) SaveAllTrades(trades []models.Trade) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	filePath := h.GetTradesFilePath()
	tempPath := filePath + ".tmp"

	// Create temporary file
	file, err := os.Create(tempPath)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	headers := []string{
		"id", "timestamp", "exchange", "pair", "symbol",
		"type", "quantity", "price", "fee", "notes", "created_at",
	}

	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	// Write trade records
	for _, trade := range trades {
		record := h.formatTradeRecord(trade)
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write trade record: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("failed to flush CSV writer: %w", err)
	}

	file.Close()

	// Atomic rename to replace original file
	if err := os.Rename(tempPath, filePath); err != nil {
		return fmt.Errorf("failed to replace trades.csv: %w", err)
	}

	return nil
}

// AppendTrade adds a single trade to the CSV file
func (h *CSVHandler) AppendTrade(trade models.Trade) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	filePath := h.GetTradesFilePath()

	// Open file in append mode
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open trades.csv for append: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := h.formatTradeRecord(trade)
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write trade record: %w", err)
	}

	return nil
}

// parseTradeRecord converts CSV record to Trade struct
func (h *CSVHandler) parseTradeRecord(record []string) (models.Trade, error) {
	if len(record) != 11 {
		return models.Trade{}, fmt.Errorf("invalid record length: expected 11 fields, got %d", len(record))
	}

	// Parse timestamp
	timestamp, err := time.Parse(time.RFC3339, record[1])
	if err != nil {
		return models.Trade{}, fmt.Errorf("invalid timestamp format: %w", err)
	}

	// Parse quantity
	quantity, err := decimal.NewFromString(record[6])
	if err != nil {
		return models.Trade{}, fmt.Errorf("invalid quantity format: %w", err)
	}

	// Parse price
	price, err := decimal.NewFromString(record[7])
	if err != nil {
		return models.Trade{}, fmt.Errorf("invalid price format: %w", err)
	}

	// Parse fee
	fee, err := decimal.NewFromString(record[8])
	if err != nil {
		return models.Trade{}, fmt.Errorf("invalid fee format: %w", err)
	}

	// Parse created_at
	createdAt, err := time.Parse(time.RFC3339, record[10])
	if err != nil {
		return models.Trade{}, fmt.Errorf("invalid created_at format: %w", err)
	}

	return models.Trade{
		ID:        record[0],
		Timestamp: timestamp,
		Exchange:  record[2],
		Pair:      record[3],
		Symbol:    record[4],
		Type:      models.TradeType(record[5]),
		Quantity:  quantity,
		Price:     price,
		Fee:       fee,
		Notes:     record[9],
		CreatedAt: createdAt,
	}, nil
}

// formatTradeRecord converts Trade struct to CSV record
func (h *CSVHandler) formatTradeRecord(trade models.Trade) []string {
	return []string{
		trade.ID,
		trade.Timestamp.Format(time.RFC3339),
		trade.Exchange,
		trade.Pair,
		trade.Symbol,
		string(trade.Type),
		trade.Quantity.String(),
		trade.Price.String(),
		trade.Fee.String(),
		trade.Notes,
		trade.CreatedAt.Format(time.RFC3339),
	}
}

// FindTradeByID searches for a trade by ID
func (h *CSVHandler) FindTradeByID(id string) (*models.Trade, error) {
	trades, err := h.LoadAllTrades()
	if err != nil {
		return nil, err
	}

	for _, trade := range trades {
		if trade.ID == id {
			return &trade, nil
		}
	}

	return nil, fmt.Errorf("trade with ID %s not found", id)
}

// DeleteTradeByID removes a trade by ID
func (h *CSVHandler) DeleteTradeByID(id string) error {
	trades, err := h.LoadAllTrades()
	if err != nil {
		return err
	}

	// Find and remove the trade
	found := false
	filteredTrades := make([]models.Trade, 0, len(trades))

	for _, trade := range trades {
		if trade.ID != id {
			filteredTrades = append(filteredTrades, trade)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("trade with ID %s not found", id)
	}

	// Save the updated trades list
	return h.SaveAllTrades(filteredTrades)
}
