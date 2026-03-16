package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "coinpilot",
	Short: "CoinPilot - Cryptocurrency Portfolio Management CLI",
	Long: `CoinPilot is an offline cryptocurrency portfolio management CLI tool 
for recording trades, calculating positions, analyzing returns, and displaying asset dashboards.

All data is stored locally without requiring internet connectivity or exchange API integration.`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CoinPilot - Cryptocurrency Portfolio Management CLI")
		fmt.Println("Use 'coinpilot --help' for available commands")
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands here as they are implemented
}
