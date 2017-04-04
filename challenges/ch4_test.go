package challenges

import (
	"testing"
)

var ch4Tests = []struct {
	input      string
	expectKey  byte
	expectText []byte
}{
	{"../data/4.txt",
		byte('\x35'),
		[]byte("Now that the party is jumping\n")},
}

func Test_ch4Solution(t *testing.T) {
	for _, tt := range ch4Tests {
		actualText, actualKey := ch4Solution(tt.input)
		expect(t, tt.expectText, actualText)
		expect(t, tt.expectKey, actualKey)
	}
}
