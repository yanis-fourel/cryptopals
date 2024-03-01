package set1

import (
	"cmp"
	"log"
	"slices"
)

// returns Key, engScore, Decrypted
func DecryptRkxor(cipher []byte) ([]byte, float32, string) {
	keySizes := getKeySizes(cipher)

	var bestkey []byte
	var bestscore float32

	for _, keySize := range keySizes {
		k, v := solveWithKeySize(cipher, keySize)
		if v > bestscore {
			bestkey = k
			bestscore = v
		}
	}

	decrypted := RepeatedXor(cipher, bestkey)
	return bestkey, bestscore, string(decrypted)
}

func solveWithKeySize(cipher []byte, ks int) ([]byte, float32) {
	bestkey := make([]byte, ks)
	var bestsumscore float32

	for i := 0; i < ks; i++ {
		sameByteXor := make([]byte, 0)

		for j := i; j < len(cipher); j += ks {
			sameByteXor = append(sameByteXor, cipher[j])
		}

		k, s, _ := DecryptEnglishSingleByteXOR(sameByteXor)
		bestkey[i] = k
		bestsumscore += s.engScore
	}

	return bestkey, bestsumscore / float32(ks)
}

type keysizeScore = struct {
	keysize int
	score   float32
}

func getKeySizes(cipher []byte) []int {
	keysizeScores := make([]keysizeScore, 0)
	for i := 2; i < 40; i++ {
		s1 := cipher[:i]
		s2 := cipher[i+1 : 2*i+1]
		keysizeScores = append(keysizeScores, keysizeScore{
			keysize: i,
			score:   float32(HammingDistance(s1, s2)) / float32(i),
		})
	}
	slices.SortFunc(keysizeScores, func(a, b keysizeScore) int { return cmp.Compare(b.score, a.score) })

	keysizes := make([]int, 3)
	for i := 0; i < len(keysizes); i++ {
		keysizes[i] = keysizeScores[i].keysize
	}
	return keysizes
}

func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		log.Fatalln("Tina says it can't")
	}

	distance := 0
	for i := 0; i < len(a); i++ {
		v := a[i] ^ b[i]
		for j := 0; j < 8; j++ {
			distance += int(v & 1)
			v >>= 1
		}
	}
	return distance
}
