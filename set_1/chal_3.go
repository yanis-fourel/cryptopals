package set1

import (
	"math"
	"slices"
	"strings"
)

type attempt = struct {
	b         byte
	decrypted string
	engScore  float32
}

func DecryptEnglishSingleByteXOR(hex string) (string, error) {
	bytes, err := HexToBytes(hex)
	if err != nil {
		return "", err
	}

	attempts := make([]attempt, 0, 128)
	for b := 0; b < 128; b++ {
		bb := byte(b)
		decrypted := string(SingleByteXOR(bytes, bb))
		attempts = append(attempts, attempt{
			b:         bb,
			decrypted: decrypted,
			engScore:  EnglishLetterFreqScore(decrypted),
		})
	}

	slices.SortFunc(attempts, func(a, b attempt) int {
		if a.engScore > b.engScore {
			return -1
		}
		if a.engScore < b.engScore {
			return 1
		}
		return 0
	})

	return attempts[0].decrypted, nil
}

func SingleByteXOR(bytes []byte, b byte) []byte {
	res := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		res[i] = bytes[i] ^ b
	}
	return res
}

// Returns value in range [0;1], 1 matching exactly the english letter frequency
func EnglishLetterFreqScore(str string) float32 {
	singleFreq := map[byte]float32{
		byte('a'): 0.082,
		byte('b'): 0.015,
		byte('c'): 0.028,
		byte('d'): 0.043,
		byte('e'): 0.127,
		byte('f'): 0.022,
		byte('g'): 0.02,
		byte('h'): 0.061,
		byte('i'): 0.07,
		byte('j'): 0.0015,
		byte('k'): 0.0077,
		byte('l'): 0.04,
		byte('m'): 0.024,
		byte('n'): 0.067,
		byte('o'): 0.075,
		byte('p'): 0.019,
		byte('q'): 9.5e-4,
		byte('r'): 0.06,
		byte('s'): 0.063,
		byte('t'): 0.091,
		byte('u'): 0.028,
		byte('v'): 0.0098,
		byte('w'): 0.024,
		byte('x'): 0.0015,
		byte('y'): 0.02,
		byte('z'): 0.074,
		byte(' '): 0.15,
	}

	var score float32 = 1.0
	str = strings.ToLower(str)
	for l := range singleFreq {
		letter := byte(l)
		c := strings.Count(str, string(letter))
		freq := float32(c) / float32(len(str))
		score *= float32(1.0 - math.Abs((float64(freq - singleFreq[letter]))))
	}

	return score
}
