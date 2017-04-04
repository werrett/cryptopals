package challenges

import (
	"github.com/spf13/cobra"
)

// XOR byte buffer with a repeating key
func xorBufferWithKey(buf []byte, key []byte) ([]byte, error) {

	res := make([]byte, len(buf))

	for i := 0; i < len(res); i += len(key) {
		for j := range key {
			res[i+j] = buf[i+j] ^ key[j]
		}
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
	var plainText []byte
	var secretKey byte

	for k := byte('\x00'); k < 255; k++ {
		res, err := xorBufferWithKey(buf, []byte{k})
		check(err)

		score := englishLetterCount(res)
		log(LogDebug, "score: %2d decrypt: %s\n", score, res)
		if score > highScore {
			log(LogInfo, "new high score: %2d decrypt: %s\n", score, res)
			highScore = score
			plainText = res
			secretKey = k
		}
	}

	log(LogInfo, "high score: %d\n", highScore)
	print("unxor'd text: %s\n", plainText)
	print("key: %c (\\x%x)\n", secretKey, secretKey)
	return secretKey
}

// Command definition and help
var ch3Cmd = &cobra.Command{
	Use:   "ch3",
	Short: "Find repeating XOR key in STRING",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			panic(newError("Must provide STRING in hex format"))
		}

		ch3Solution(args[0])
	},
}

func init() {
	RootCmd.AddCommand(ch3Cmd)
}
