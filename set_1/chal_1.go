package set1

import (
	"fmt"
	"log"
	"strings"
)

func HexToBase64(hex string) (string, error) {
	bytes, err := HexToBytes(hex)
	if err != nil {
		log.Fatal(err)
	}
	return BytesToBase64(bytes), nil
}

func HexToBytes(hex string) ([]byte, error) {
	if len(hex)%2 != 0 {
		return nil, fmt.Errorf("Hex input length is an odd number (%d)", len(hex))
	}

	base := "0123456789abcdef"
	result := make([]byte, 0, len(hex)/2)
	hex = strings.ToLower(hex)

	for i := 0; i < len(hex); i += 2 {
		v1 := strings.IndexByte(base, hex[i])
		if v1 == -1 {
			return nil, fmt.Errorf("Character %c not in base %s", hex[i], base)
		}
		v2 := strings.IndexByte(base, hex[i+1])
		if v2 == -1 {
			return nil, fmt.Errorf("Character %c not in base %s", hex[i+1], base)
		}
		value := v1<<4 + v2
		result = append(result, byte(value))
	}
	return result, nil
}

func BytesToBase64(b []byte) string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	result := make([]byte, 0, len(b))
	for i := 0; i < len(b); i++ {
		switch i % 3 {
		case 0:
			result = append(result, base[(b[i]&0b11111100)>>2])
			if i == len(b)-1 {
				result = append(result, base[b[i]&0b00000011]<<4)
				result = append(result, byte('='))
				result = append(result, byte('='))
			}
		case 1:
			result = append(result, base[(b[i-1]&0b00000011)<<4|(b[i]&0b11110000)>>4])
			if i == len(b)-1 {
				result = append(result, base[b[i]&0b00001111]<<2)
				result = append(result, byte('='))
			}
		case 2:
			result = append(result, base[(b[i-1]&0b00001111)<<2|(b[i]&0b11000000)>>6])
			result = append(result, base[(b[i]&0b00111111)])
		}
	}
	return string(result)
}
