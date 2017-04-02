package challenges

import (
	"testing"
)

var ch2Tests = []struct {
  str1  string
  str2  string
	expected string 
}{
  {"1c0111001f010100061a024b53535009181c",
	 "686974207468652062756c6c277320657965",
	 "746865206b696420646f6e277420706c6179"},
}

func Test_ch2_solution(t *testing.T) {
  for _, tt := range ch2Tests {
    actual := ch2_solution(tt.str1, tt.str2)
    expect(t, tt.expected, actual)
  }
}