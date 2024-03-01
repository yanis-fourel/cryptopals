package set1

import (
	"testing"
)

func TestChal3_1(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedOutput := "Cooking MC's like a pound of bacon"
	_, actualOutput, err := DecryptEnglishSingleByteXORhex(input)
	if actualOutput.decrypted != expectedOutput || err != nil {
		t.Fatalf(`Expected %s, %v but got %s, %v`, expectedOutput, nil, actualOutput.decrypted, err)
	}
}
