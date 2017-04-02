package challenges

import (
	"testing"
)

var ch3Tests = []struct {
  input  string
	expected byte 
}{
  {"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
	 'X'},
}

func Test_ch3Solution(t *testing.T) {
  for _, tt := range ch3Tests {
    actual := ch3Solution(tt.input)
    expect(t, tt.expected, actual)
  }
}