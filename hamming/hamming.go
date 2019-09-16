package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("both strings need to be of the same size")
	}

	diff := 0
	rb := []rune(strings.ToUpper(b))

	for idx, nuc := range []rune(strings.ToUpper(a)) {
		if nuc != rb[idx] {
			diff += 1
		}
	}
	return diff, nil
}
