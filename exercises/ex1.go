package exercises 

import (
  "encoding/base64"
  "encoding/hex"
  
	"github.com/spf13/cobra"
)

// Convert from a hex string to bytes
func unhex(src string) []byte {
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
func ex1_solution(str string) []byte {
  buf := unhex(str)
  return b64_bytes(buf) 
}

// Command definition and help
var ex1Cmd = &cobra.Command{
	Use:   "ex1",
	Short: "Convert hex to base64",
	Run: func(cmd *cobra.Command, args []string) {

    buf := ex1_solution(args[0])
    println(string(buf)) 
	},
}

func init() {
	RootCmd.AddCommand(ex1Cmd)
}
