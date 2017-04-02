package challenges 

import (
  "fmt"
	"github.com/spf13/cobra"
)

// XOR byte buffer with a repeating key
// FIXME: Generalize this to a multiple byte key
func xorBufferWithKey(buf []byte, key byte) ([]byte, error) {
  
  res := make([]byte, len(buf))
  
  for i := range res {
    res[i] = buf[i] ^ key
  } 
  
  return res, nil
}

// Return number of characters in a buffer that are either alphanumerica or 
// common punctuation
func englishLetterCount(buf []byte) int {
  count := 0
  
  for i := range buf {
    c := buf[i]
    switch {
    case 'a' <= c && c <= 'z':
      count++
    case 'A' <= c && c <= 'Z':
    	count++
    case '0' <= c && c <= '9':
    	count++
    case c == '!' || c == '?': 
    	count++
    case c == '\'' || c == '"': 
    	count++
    case c == ',' || c == '.': 
    	count++
    case c == ':' || c == ';': 
    	count++
    case c == ' ':
    	count++
    }   
  }
  return count
}

// Solution
// XXX: Extra credit: Use channels to search for key in parallel
func ch3Solution(input string) byte {
  buf := unhexString(input)
  
  highScore := 0 
  var cipherText []byte
  var secretKey byte
  
  for k := byte('\x00'); k < 255; k++ {
    res, err := xorBufferWithKey(buf, k)
    check(err)
    
    score := englishLetterCount(res)
    log("raw result: %2d %s\n", score, res) 
    if (score > highScore) {
      highScore = score
      cipherText = res
      secretKey = k
    }
  }
  
  fmt.Printf("raw answer: %d %s\n", highScore, cipherText) 
  return secretKey
}

// Command definition and help
var ch3Cmd = &cobra.Command{
	Use:   "ch3",
	Short: "Find repeating XOR key in STRING",
	Run: func(cmd *cobra.Command, args []string) {

    if (len(args) != 1) {
      panic( newError("Must provide STRING in hex format") )
    }
    
    key := ch3Solution(args[0])
    fmt.Printf("Secret key: %c (\\x%x)\n", key, key) 
	},
}

func init() {
	RootCmd.AddCommand(ch3Cmd)
}
