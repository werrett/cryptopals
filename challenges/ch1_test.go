package challenges

import (
	"testing"
)

var ch1Tests = []struct {
  input    string
	expected string 
}{
  {"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
}

func Test_ch1_solution(t *testing.T) {
  for _, tt := range ch1Tests {
    actual := ch1_solution(tt.input)
    expect(t, tt.expected, actual)
  }
}