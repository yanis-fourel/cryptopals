package set1

import "testing"

func TestChal2(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expectedOutput := "746865206b696420646f6e277420706c6179"

	actualOutput, err := FixedXOR(input1, input2)

	if err != nil || actualOutput != expectedOutput {
		t.Fatalf(`Expected %s, %v but got %s, %v`, expectedOutput, nil, actualOutput, err)
	}
}
