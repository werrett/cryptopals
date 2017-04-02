package challenges 

import (
  "encoding/base64"
  "encoding/hex"
	"github.com/spf13/cobra"
)

// Convert from a hex string to bytes
func unhex_string(src string) []byte {
  buf, err := hex.DecodeString(src)
  check(err)
  log("decoded hex: %s\n", buf)
  return buf
}

// Base64 encode a byte array
func b64_bytes(src []byte) []byte {
  buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
  base64.StdEncoding.Encode(buf, src)
  return buf
}

// Solution
func ch1_solution(str string) string {
  buf := unhex_string(str)
  return string(b64_bytes(buf))
}

// Command definition and help
var ch1Cmd = &cobra.Command{
	Use:   "ch1",
	Short: "Convert hex to base64",
	Run: func(cmd *cobra.Command, args []string) {

    buf := ch1_solution(args[0])
    println(buf) 
	},
}

func init() {
	RootCmd.AddCommand(ch1Cmd)
}