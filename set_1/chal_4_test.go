package set1

import (
	"os"
	"testing"
)

func TestChal4(t *testing.T) {
	file, err := os.Open("./chal_4_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	_, res, err := GetBestSbxor(file)

	if err != nil {
		t.Fatalf("Err: %v", err)
	}
	expected := "Now that the party is jumping\n"
	if res.decrypted != expected {
		t.Fatalf("Result differ, expected '%s', got '%s'", expected, res.decrypted)
	}
}
