package cmd

import (
  "encoding/base64"
  "encoding/hex"
  
	"github.com/spf13/cobra"
)

func Base64Bytes(src []byte) []byte {
  buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
  base64.StdEncoding.Encode(buf, src)
  return buf
}

func UnhexString(src string) []byte {
  buf, err := hex.DecodeString(src)
  check(err)
  log("decoded hex: %s\n", buf)
  return buf
}

func ex1Solution(input string) string {
  buf := UnhexString(input)
  encoded := Base64Bytes(buf)
  return string(encoded)
}

// Command definition and help
var ex1Cmd = &cobra.Command{
	Use:   "ex1",
	Short: "Convert hex to base64",
	Run: func(cmd *cobra.Command, args []string) {

		log("input: %s\n", args[0])
    encoded := ex1Solution(args[0])

    log("answer: %s\n", string(encoded))
	},
}

func init() {
	RootCmd.AddCommand(ex1Cmd)
}
