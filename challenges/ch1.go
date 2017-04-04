package challenges

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/spf13/cobra"
)

// Convert from a hex string to bytes
func unhexString(src string) []byte {
	buf, err := hex.DecodeString(src)
	check(err)
	log(LogDebug, "decoded hex: %s\n", buf)
	return buf
}

// Base64 encode a byte array
func b64Bytes(src []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)

	log(LogDebug, "b64 bytesx: %s\n", buf)
	return buf
}

// Solution
func ch1Solution(str string) string {
	buf := unhexString(str)
	res := string(b64Bytes(buf))

	print("result: %s", res)
	return res
}

// Command definition and help
var ch1Cmd = &cobra.Command{
	Use:   "ch1",
	Short: "Convert hex to base64",
	Run: func(cmd *cobra.Command, args []string) {

		ch1Solution(args[0])
	},
}

func init() {
	RootCmd.AddCommand(ch1Cmd)
}
