package filehandler

import (
	"coinpilot/internal/models"
	"os"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTradeFileHandler_Create(t *testing.T) {
	// Setup temporary directory for test
	tempDir := t.TempDir()
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	handler := NewTradeFileHandler()
	err := handler.Initialize()
	require.NoError(t, err)

	// Test valid trade creation
	trade := models.Trade{
		ID:        "test_trade_1",
		Timestamp: time.Now(),
		Exchange:  "binance",
		Pair:      "BTC/USDT",
		Symbol:    "BTC",
		Type:      models.BUY,
		Quantity:  decimal.NewFromFloat(0.5),
		Price:     decimal.NewFromFloat(45000.0),
		Fee:       decimal.NewFromFloat(22.5),
		Notes:     "Test trade",
		CreatedAt: time.Now(),
	}

	err = handler.Create(trade)
	assert.NoError(t, err)

	// Verify trade was created
	trades, err := handler.FindAll()
	require.NoError(t, err)
	assert.Len(t, trades, 1)
	assert.Equal(t, "test_trade_1", trades[0].ID)
	assert.Equal(t, "BTC", trades[0].Symbol)
}

func TestTradeFileHandler_Create_ValidationErrors(t *testing.T) {
	tempDir := t.TempDir()
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	handler := NewTradeFileHandler()
	err := handler.Initialize()
	require.NoError(t, err)

	tests := []struct {
		name  string
		trade models.Trade
		error string
	}{
		{
			name: "empty ID",
			trade: models.Trade{
				ID:        "",
				Symbol:    "BTC",
				Type:      models.BUY,
				Quantity:  decimal.NewFromFloat(1.0),
				Price:     decimal.NewFromFloat(45000.0),
				Timestamp: time.Now(),
				CreatedAt: time.Now(),
			},
			error: "ID is required",
		},
		{
			name: "empty symbol",
			trade: models.Trade{
				ID:        "test_1",
				Symbol:    "",
				Type:      models.BUY,
				Quantity:  decimal.NewFromFloat(1.0),
				Price:     decimal.NewFromFloat(45000.0),
				Timestamp: time.Now(),
				CreatedAt: time.Now(),
			},
			error: "Symbol is required",
		},
		{
			name: "zero quantity",
			trade: models.Trade{
				ID:        "test_1",
				Symbol:    "BTC",
				Type:      models.BUY,
				Quantity:  decimal.Zero,
				Price:     decimal.NewFromFloat(45000.0),
				Timestamp: time.Now(),
				CreatedAt: time.Now(),
			},
			error: "Quantity must be positive",
		},
		{
			name: "negative price",
			trade: models.Trade{
				ID:        "test_1",
				Symbol:    "BTC",
				Type:      models.BUY,
				Quantity:  decimal.NewFromFloat(1.0),
				Price:     decimal.NewFromFloat(-45000.0),
				Timestamp: time.Now(),
				CreatedAt: time.Now(),
			},
			error: "Price must be positive",
		},
		{
			name: "unsupported trade type",
			trade: models.Trade{
				ID:        "test_1",
				Symbol:    "BTC",
				Type:      models.TRANSFER_IN,
				Quantity:  decimal.NewFromFloat(1.0),
				Price:     decimal.NewFromFloat(45000.0),
				Timestamp: time.Now(),
				CreatedAt: time.Now(),
			},
			error: "unsupported trade type 'TRANSFER_IN' in MVP phase",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.Create(tt.trade)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.error)
		})
	}
}

func TestTradeFileHandler_FindAll(t *testing.T) {
	tempDir := t.TempDir()
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	handler := NewTradeFileHandler()
	err := handler.Initialize()
	require.NoError(t, err)

	// Test empty file
	trades, err := handler.FindAll()
	require.NoError(t, err)
	assert.Len(t, trades, 0)

	// Add some trades
	trade1 := models.Trade{
		ID:        "trade_1",
		Timestamp: time.Now(),
		Exchange:  "binance",
		Symbol:    "BTC",
		Type:      models.BUY,
		Quantity:  decimal.NewFromFloat(0.5),
		Price:     decimal.NewFromFloat(45000.0),
		Fee:       decimal.Zero,
		CreatedAt: time.Now(),
	}

	trade2 := models.Trade{
		ID:        "trade_2",
		Timestamp: time.Now(),
		Exchange:  "okx",
		Symbol:    "ETH",
		Type:      models.SELL,
		Quantity:  decimal.NewFromFloat(2.0),
		Price:     decimal.NewFromFloat(2500.0),
		Fee:       decimal.Zero,
		CreatedAt: time.Now(),
	}

	err = handler.Create(trade1)
	require.NoError(t, err)
	err = handler.Create(trade2)
	require.NoError(t, err)

	// Verify both trades are returned
	trades, err = handler.FindAll()
	require.NoError(t, err)
	assert.Len(t, trades, 2)
}

func TestTradeFileHandler_Delete(t *testing.T) {
	tempDir := t.TempDir()
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tempDir)

	handler := NewTradeFileHandler()
	err := handler.Initialize()
	require.NoError(t, err)

	// Add a trade
	trade := models.Trade{
		ID:        "trade_to_delete",
		Timestamp: time.Now(),
		Exchange:  "binance",
		Symbol:    "BTC",
		Type:      models.BUY,
		Quantity:  decimal.NewFromFloat(0.5),
		Price:     decimal.NewFromFloat(45000.0),
		Fee:       decimal.Zero,
		CreatedAt: time.Now(),
	}

	err = handler.Create(trade)
	require.NoError(t, err)

	// Verify trade exists
	trades, err := handler.FindAll()
	require.NoError(t, err)
	assert.Len(t, trades, 1)

	// Delete the trade
	err = handler.Delete("trade_to_delete")
	assert.NoError(t, err)

	// Verify trade is deleted
	trades, err = handler.FindAll()
	require.NoError(t, err)
	assert.Len(t, trades, 0)

	// Test deleting non-existent trade
	err = handler.Delete("non_existent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// Test deleting with empty ID
	err = handler.Delete("")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ID cannot be empty")
}
