package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommand
var RootCmd = &cobra.Command{
	Use:   "policy server",
	Short: "the policy server application",
	Long:  "Run the 'serve' subcommand to start the http server",
}

// Execute adds all child command to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
