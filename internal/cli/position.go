package cli

import (
	"fmt"
	"strings"

	"coinpilot/internal/services"

	"github.com/spf13/cobra"
)

var positionService = services.NewPositionService()

var positionCmd = &cobra.Command{
	Use:   "position",
	Short: "Manage cryptocurrency positions",
	Long:  "View and manage calculated cryptocurrency positions",
}

var positionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all positions",
	Long:  "Display all calculated cryptocurrency positions",
	RunE:  runPositionList,
}

func init() {
	// Add position subcommands
	positionCmd.AddCommand(positionListCmd)

	// Add position command to root
	rootCmd.AddCommand(positionCmd)
}

func runPositionList(cmd *cobra.Command, args []string) error {
	positions, err := positionService.CalculatePositions()
	if err != nil {
		return fmt.Errorf("failed to calculate positions: %v", err)
	}

	if len(positions) == 0 {
		fmt.Println("No positions found")
		return nil
	}

	// Simple table output for MVP
	fmt.Printf("%-8s %-12s %-12s %-12s %-12s %-12s\n",
		"Symbol", "Exchange", "Quantity", "Avg Cost", "Total Cost", "Status")
	fmt.Println(strings.Repeat("-", 80))

	for _, position := range positions {
		status := "Open"
		if position.IsClosed {
			status = "Closed"
		}

		fmt.Printf("%-8s %-12s %-12s %-12s %-12s %-12s\n",
			position.Symbol,
			position.Exchange,
			position.Quantity.String(),
			position.AverageCost.String(),
			position.TotalCost.String(),
			status)
	}

	return nil
}
