package challenges

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Solution
func ch4Solution(filePath string) ([]byte, byte) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	highScore := 0
	var plainText []byte
	var secretKey byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexText := scanner.Text()
		buf := unhexString(hexText)

		for k := byte('\x00'); k < 255; k++ {
			res, err := xorBufferWithKey(buf, k)
			check(err)

			score := englishLetterCount(res)
			if score > highScore {
				log("raw result: %2d %s\n", score, buf)
				highScore = score
				plainText = res
				secretKey = k
			}
		}
	}

	err = scanner.Err()
	check(err)

	log("score: %d key: %x, plaintext: %s\n", highScore, secretKey, plainText)
	return plainText, secretKey
}

// Command definition and help
var ch4Cmd = &cobra.Command{
	Use:   "ch4",
	Short: "Find a string XOR'd by a 1-byte key from a file of hex-encoded data",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			panic(newError("Must provide path to hex-encoded data"))
		}

		text, key := ch4Solution(args[0])
		fmt.Printf("key: %c (\\x%x)\n", key, key)
		fmt.Printf("plain text: %s", text)
	},
}

func init() {
	RootCmd.AddCommand(ch4Cmd)
}
