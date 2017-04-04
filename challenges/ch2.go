package challenges

import (
	"encoding/hex"
	"github.com/spf13/cobra"
)

// Convert from a bytes to hex
func hexBytes(src []byte) string {
	str := hex.EncodeToString(src)
	log(LogDebug, "encoded hex: %s\n", str)
	return str
}

// XOR two byte buffers of the same length
func xorFixedBuffer(buf1, buf2 []byte) ([]byte, error) {
	// Gotta be the same length
	if len(buf1) != len(buf2) {
		return nil, newError("Buffers to XOR must be same length")
	}

	res := make([]byte, len(buf1))

	for i := range res {
		res[i] = buf1[i] ^ buf2[i]
	}

	return res, nil
}

// Solution
func ch2Solution(str1, str2 string) string {
	buf1 := unhexString(str1)
	buf2 := unhexString(str2)

	res, err := xorFixedBuffer(buf1, buf2)
	check(err)
	log(LogInfo, "xor result: %s\n", res)

	hexRes := hexBytes(res)
	print("hex result: %s\n", hexRes)
	return hexRes
}

// Command definition and help
var ch2Cmd = &cobra.Command{
	Use:   "ch2",
	Short: "XOR two, equal length, strings",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			panic(newError("Must provided two hex strings to XOR"))
		}

		if len(args[0]) != len(args[1]) {
			panic(newError("Hex strings must be same length"))
		}

		ch2Solution(args[0], args[1])
	},
}

func init() {
	RootCmd.AddCommand(ch2Cmd)
}
