package challenges

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Global log verbosity
var Verbose int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cryptopals",
	Short: "Solutions to the Cryptopals exercises",
	Long: `Working my way through the Cryptopal exercises
Along with the other b1t tw1ddlers at $WORK.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// Persistent Flags will be global for your application.
	RootCmd.PersistentFlags().IntVarP(&Verbose, "verbose", "v", LogInfo, "verbose output")
}
