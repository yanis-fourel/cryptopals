package set1

import (
	"bufio"
	"fmt"
	"io"
)

func GetBestSbxor(reader io.Reader) (byte, DecryptAttempt, error) {
	scanner := bufio.NewScanner(reader)
	line := -1

	var bestk byte
	var bestAttempt DecryptAttempt

	for scanner.Scan() {
		line += 1
		cipherHex := scanner.Text()
		k, res, err := DecryptEnglishSingleByteXORhex(cipherHex)
		if err != nil {
			return 0, DecryptAttempt{}, fmt.Errorf("Error on line %d: %w\n", line, err)
		}
		if res.engScore > bestAttempt.engScore {
			bestk = k
			bestAttempt = res
		}
	}

	return bestk, bestAttempt, nil
}
