package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ex1Cmd represents the ex1 command
var ex1Cmd = &cobra.Command{
	Use:   "ex1",
	Short: "Convert hex to base64",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("ex1 called")
	},
}

func init() {
	RootCmd.AddCommand(ex1Cmd)
}
