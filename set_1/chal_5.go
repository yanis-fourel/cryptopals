package set1

import ()

func EncryptWithRepeatedXor(plaintext, key string) string {
	res := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		res[i] = plaintext[i] ^ key[i%len(key)]
	}
	return BytesToHex(res)
}
