package cmd

import (
	"fmt"
  "encoding/base64"
  "encoding/hex"
  
	"github.com/spf13/cobra"
)

func base64Bytes(src []byte) []byte {
  buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
  base64.StdEncoding.Encode(buf, src)
  return buf
}

func unhexString(src string) []byte {
  buf, err := hex.DecodeString(src)
  check(err)
  return buf
}

// Command definition and help
var ex1Cmd = &cobra.Command{
	Use:   "ex1",
	Short: "Convert hex to base64",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("input: %s\n", args[0])
    buf := unhexString(args[0])
    
    fmt.Printf("bytes: %s\n", buf)
    encoded := base64Bytes(buf)

    fmt.Printf("base64: %s\n", string(encoded))
	},
}

func init() {
	RootCmd.AddCommand(ex1Cmd)
}
