package set1

import ()

func RepeatedXor(plaintext, key []byte) []byte {
	res := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		res[i] = plaintext[i] ^ key[i%len(key)]
	}
	return res
}

func RepeatedXorhex(plaintext, key []byte) string {
	return BytesToHex(RepeatedXor(plaintext, key))
}
