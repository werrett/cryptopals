package challenges

import (
	"testing"
)

var ch3Tests = []struct {
	input      string
	expectKey  byte
	expectText []byte
}{
	{"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
		byte('\x58'),
		[]byte("Cooking MC's like a pound of bacon")},
	{"52461f4c5a5c4d5a4b4c1f5e4d5a1f4c5e595a1f48564b571f46504a1f584a464c",
		byte('\x3f'),
		[]byte("my secrets are safe with you guys")},
}

func Test_ch3Solution(t *testing.T) {
	for _, tt := range ch3Tests {
		actualText, actualKey := ch3Solution(tt.input)
		expect(t, string(tt.expectText), string(actualText))
		expect(t, tt.expectKey, actualKey)
	}
}

// xorBufferWithKey is a little complex so lets test it.
var xorBufferWithKeyTests = []struct {
	plaintext  []byte
	key        []byte
	ciphertext []byte
}{
	{[]byte("deadbeef"),
		[]byte("12345"),
		[]byte("UWRPWTWU")},

	{[]byte("d"),
		[]byte("12345"),
		[]byte("U")},

	{[]byte("deadbeef"),
		[]byte("cafebabe"),
		[]byte("\a\x04\a\x01\x00\x04\a\x03")},

	{[]byte("deadbeef"),
		[]byte("1"),
		[]byte("UTPUSTTW")},
}

func Test_xorBufferWithKey(t *testing.T) {
	for _, tt := range xorBufferWithKeyTests {
		actual, _ := xorBufferWithKey(tt.plaintext, tt.key)
		expect(t, tt.ciphertext, actual)
	}
}
