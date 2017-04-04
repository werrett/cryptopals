package challenges

import (
	"bufio"
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
			res, err := xorBufferWithKey(buf, []byte{k})
			check(err)

			score := englishLetterCount(res)
			log(LogDebug, "score: %2d decrypt: %s\n", score, res)
			if score > highScore {
				log(LogInfo, "high score: %2d decrypt: %s\n", score, res)
				highScore = score
				plainText = res
				secretKey = k
			}
		}
	}

	err = scanner.Err()
	check(err)

	log(LogInfo, "high score: %d\n", highScore)
	print("key: %c (\\x%x)\n", secretKey, secretKey)
	print("plaintext: %s\n", plainText)
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

		ch4Solution(args[0])
	},
}

func init() {
	RootCmd.AddCommand(ch4Cmd)
}
