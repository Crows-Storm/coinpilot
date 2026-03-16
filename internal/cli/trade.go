package cli

import (
	"fmt"
	"strings"
	"time"

	"coinpilot/internal/models"
	"coinpilot/internal/services"
	"coinpilot/pkg/utils"

	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var tradeService = services.NewTradeService()

var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Manage cryptocurrency trades",
	Long:  "Add, list, and delete cryptocurrency trade records",
}

var tradeAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new trade record",
	Long:  "Add a new cryptocurrency trade record with type, symbol, quantity, and price",
	RunE:  runTradeAdd,
}

var tradeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all trade records",
	Long:  "Display all recorded cryptocurrency trades",
	RunE:  runTradeList,
}

var tradeDeleteCmd = &cobra.Command{
	Use:   "delete [trade-id]",
	Short: "Delete a trade record",
	Long:  "Delete a specific trade record by its ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runTradeDelete,
}

// Command flags
var (
	tradeType string
	symbol    string
	quantity  string
	price     string
	exchange  string
	fee       string
	notes     string
)

func init() {
	// Add trade subcommands
	tradeCmd.AddCommand(tradeAddCmd)
	tradeCmd.AddCommand(tradeListCmd)
	tradeCmd.AddCommand(tradeDeleteCmd)

	// Add flags for trade add command
	tradeAddCmd.Flags().StringVarP(&tradeType, "type", "t", "", "Trade type (BUY, SELL)")
	tradeAddCmd.Flags().StringVarP(&symbol, "symbol", "s", "", "Trading symbol (e.g., BTC)")
	tradeAddCmd.Flags().StringVarP(&quantity, "quantity", "q", "", "Trade quantity")
	tradeAddCmd.Flags().StringVarP(&price, "price", "p", "", "Trade price")
	tradeAddCmd.Flags().StringVarP(&exchange, "exchange", "e", "default", "Exchange name")
	tradeAddCmd.Flags().StringVarP(&fee, "fee", "f", "0", "Trade fee")
	tradeAddCmd.Flags().StringVarP(&notes, "notes", "n", "", "Trade notes")

	// Mark required flags
	tradeAddCmd.MarkFlagRequired("type")
	tradeAddCmd.MarkFlagRequired("symbol")
	tradeAddCmd.MarkFlagRequired("quantity")
	tradeAddCmd.MarkFlagRequired("price")

	// Add trade command to root
	rootCmd.AddCommand(tradeCmd)
}

func runTradeAdd(cmd *cobra.Command, args []string) error {
	// Parse decimal values
	qty, err := decimal.NewFromString(quantity)
	if err != nil {
		return fmt.Errorf("invalid quantity: %v", err)
	}

	prc, err := decimal.NewFromString(price)
	if err != nil {
		return fmt.Errorf("invalid price: %v", err)
	}

	feeDecimal, err := decimal.NewFromString(fee)
	if err != nil {
		return fmt.Errorf("invalid fee: %v", err)
	}

	// Validate parameters
	tradeTypeEnum := models.TradeType(tradeType)
	if err := utils.ValidateTradeParameters(tradeTypeEnum, symbol, qty, prc); err != nil {
		return err
	}

	// Create trade record
	trade := models.Trade{
		ID:        generateTradeID(),
		Timestamp: time.Now(),
		Exchange:  exchange,
		Pair:      symbol + "/USDT", // Default pairing for MVP
		Symbol:    symbol,
		Type:      tradeTypeEnum,
		Quantity:  qty,
		Price:     prc,
		Fee:       feeDecimal,
		Notes:     notes,
		CreatedAt: time.Now(),
	}

	// Add trade
	if err := tradeService.AddTrade(trade); err != nil {
		return fmt.Errorf("failed to add trade: %v", err)
	}

	fmt.Printf("Trade added successfully: %s %s %s at %s\n",
		trade.Type, trade.Quantity, trade.Symbol, trade.Price)
	return nil
}

func runTradeList(cmd *cobra.Command, args []string) error {
	trades, err := tradeService.ListTrades()
	if err != nil {
		return fmt.Errorf("failed to list trades: %v", err)
	}

	if len(trades) == 0 {
		fmt.Println("No trades found")
		return nil
	}

	// Simple table output for MVP
	fmt.Printf("%-8s %-12s %-8s %-8s %-12s %-12s %-8s\n",
		"ID", "Date", "Type", "Symbol", "Quantity", "Price", "Exchange")
	fmt.Println(strings.Repeat("-", 80))

	for _, trade := range trades {
		fmt.Printf("%-8s %-12s %-8s %-8s %-12s %-12s %-8s\n",
			trade.ID[:8],
			trade.Timestamp.Format("2006-01-02"),
			trade.Type,
			trade.Symbol,
			trade.Quantity.String(),
			trade.Price.String(),
			trade.Exchange)
	}

	return nil
}

func runTradeDelete(cmd *cobra.Command, args []string) error {
	tradeID := args[0]

	if err := tradeService.DeleteTrade(tradeID); err != nil {
		return fmt.Errorf("failed to delete trade: %v", err)
	}

	fmt.Printf("Trade %s deleted successfully\n", tradeID)
	return nil
}

// generateTradeID generates a simple trade ID for MVP
func generateTradeID() string {
	return fmt.Sprintf("trade_%d", time.Now().Unix())
}
