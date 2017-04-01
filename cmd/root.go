package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Verbose bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cryptopals",
	Short: "Answers to the Cryptopals exercises",
	Long: `Working my way through the Cryptopal exercises
Inconjunction with the other b1t tw1ddlers at $WORK.`,
// Uncomment the following line if your bare application
// has an action associated with it:
//	Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Persistent Flags will be global for your application.

	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
