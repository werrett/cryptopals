package challenges

import (
	"github.com/spf13/cobra"
)

// Solution
func ch5Solution(input, key string) string {
	plainText := []byte(input)
	keyBytes := []byte(key)
	log(LogDebug, "plainText: %s", plainText)
	log(LogDebug, "keyBytes: %s", keyBytes)

	cipherBytes, _ := xorBufferWithKey(plainText, keyBytes)
	log(LogInfo, "cipherBytes: %s", cipherBytes)

	cipherHex := hexBytes(cipherBytes)
	print("cipherHex: %s\n", cipherHex)
	return cipherHex
}

// Command definition and help
var ch5Cmd = &cobra.Command{
	Use:   "ch5",
	Short: "XOR a string with a variable length key",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			panic(newError("Must provide plaintext STRING and KEY"))
		}

		ch5Solution(args[0], args[1])
	},
}

func init() {
	RootCmd.AddCommand(ch5Cmd)
}
