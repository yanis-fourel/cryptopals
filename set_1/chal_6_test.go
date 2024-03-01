package set1

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestHamming(t *testing.T) {
	d := HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	if d != 37 {
		t.Fatalf("Expected 37, got %d", d)
	}
}

func TestChal6(t *testing.T) {
	file, err := os.Open("./chal_6_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	b64cipher, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	cipher := make([]byte, len(b64cipher)+1)
	_, err = base64.StdEncoding.Decode(cipher, b64cipher)
	if err != nil {
		t.Fatal(err)
	}

	_, _, plaintext := DecryptRkxor(cipher)
	fmt.Printf("Chat 6 plaintext is: %s", plaintext)
}
