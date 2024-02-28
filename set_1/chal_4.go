package set1

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

func GetBestSbxor(reader io.Reader) (SbxorResult, error) {
	scanner := bufio.NewScanner(reader)
	line := -1
	results := make([]SbxorResult, 0)
	for scanner.Scan() {
		line += 1
		cipherHex := scanner.Text()
		res, err := DecryptEnglishSingleByteXOR(cipherHex)
		if err != nil {
			return SbxorResult{}, fmt.Errorf("Error on line %d: %w\n", line, err)
		}
		results = append(results, res)
	}
	slices.SortFunc(results, func(a, b SbxorResult) int {
		if a.engScore < b.engScore {
			return 1
		}
		if a.engScore > b.engScore {
			return -1
		}
		return 0
	})

	return results[0], nil
}
