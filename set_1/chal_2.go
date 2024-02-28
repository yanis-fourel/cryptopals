package set1

import "fmt"

// takes two equal-length buffers and produces their XOR combination
// Returns hexadecimal string
func FixedXOR(hex1, hex2 string) (string, error) {
	if len(hex1) != len(hex2) {
		return "", fmt.Errorf("hex1 and hex2 are of different length, %d and %d respectively", len(hex1), len(hex2))
	}
	b1, err := HexToBytes(hex1)
	if err != nil {
		return "", fmt.Errorf("hex1: %w", err)
	}
	b2, err := HexToBytes(hex2)
	if err != nil {
		return "", fmt.Errorf("hex2: %w", err)
	}

	result := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		result[i] = b1[i] ^ b2[i]
	}
	return BytesToHex(result), nil
}

func BytesToHex(bytes []byte) string {
	base := "0123456789abcdef"

	result := make([]byte, len(bytes)*2)
	for i := 0; i < len(bytes); i++ {
		result[2*i] = base[bytes[i]>>4]
		result[2*i+1] = base[bytes[i]&0x0f]
	}
	return string(result)
}
