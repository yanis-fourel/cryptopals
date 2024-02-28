package set1

import (
	"fmt"
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
	// https://en.wikipedia.org/wiki/Letter_frequency
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

	// https://www.sttmedia.com/syllablefrequency-english
	diagramFreq := map[string]float32{
		"al": 0.0093,
		"an": 0.0217,
		"ar": 0.0106,
		"as": 0.0109,
		"at": 0.0117,
		"ea": 0.0084,
		"ed": 0.0129,
		"en": 0.0137,
		"er": 0.0211,
		"es": 0.01,
		"ha": 0.0117,
		"he": 0.0365,
		"hi": 0.0107,
		"in": 0.021,
		"is": 0.0099,
		"it": 0.0124,
		"le": 0.0095,
		"me": 0.0083,
		"nd": 0.0162,
		"ne": 0.0075,
		"ng": 0.0099,
		"nt": 0.0077,
		"on": 0.0136,
		"or": 0.0109,
		"ou": 0.0141,
		"re": 0.0164,
		"se": 0.0085,
		"st": 0.0096,
		"te": 0.01,
		"th": 0.0399,
		"ti": 0.0092,
		"to": 0.0124,
		"ve": 0.0111,
		"wa": 0.0084,
	}

	// https://www.sttmedia.com/syllablefrequency-english
	trigramFreq := map[string]float32{
		"all": 0.0045,
		"and": 0.017,
		"are": 0.0033,
		"but": 0.003,
		"ent": 0.0048,
		"era": 0.0033,
		"ere": 0.0047,
		"eve": 0.0043,
		"for": 0.0056,
		"had": 0.0035,
		"hat": 0.0058,
		"hen": 0.0035,
		"her": 0.0073,
		"hin": 0.0033,
		"his": 0.0049,
		"ing": 0.0106,
		"ion": 0.0047,
		"ith": 0.0047,
		"not": 0.0056,
		"ome": 0.0031,
		"oul": 0.0041,
		"our": 0.0033,
		"sho": 0.0033,
		"ted": 0.0032,
		"ter": 0.0036,
		"tha": 0.0054,
		"the": 0.0367,
		"thi": 0.0055,
		"tio": 0.0037,
		"uld": 0.004,
		"ver": 0.0069,
		"was": 0.0066,
		"wit": 0.0046,
		"you": 0.0072,
	}

	var score float32
	str = strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		score += singleFreq[byte(str[i])]
		if i < len(str)-1 {
			score += diagramFreq[str[i:i+1]]
		}
		if i < len(str)-2 {
			score += trigramFreq[str[i:i+2]]
		}
	}

	return score / float32(len(str))
}
