package challenges

import (
	"testing"
)

var ch5Tests = []struct {
	input     string
	key       string
	expectHex string
}{
	{"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal",
		"ICE",
		"0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"},
}

func Test_ch5Solution(t *testing.T) {
	for _, tt := range ch5Tests {
		actualHex := ch5Solution(tt.input, tt.key)
		expect(t, tt.expectHex, actualHex)
	}
}
