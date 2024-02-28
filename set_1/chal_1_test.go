package set1

import (
	"testing"
)

func TestChal1(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	actualOutput, err := HexToBase64(input)
	if err != nil || expectedOutput != actualOutput {
		t.Fatalf(`Expected %s, %v but got %s, %v`, expectedOutput, nil, actualOutput, err)
	}
}
