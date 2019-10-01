package luhn

import (
	"unicode"
)

func Valid(input string) bool {
	var x []rune
	for _, i := range input {
		if i == ' ' {
			continue
		} else if unicode.IsNumber(i) {
			x = append(x, i)
		} else {
			return false
		}
	}

	l := len(x)

	if l < 2 {
		return false
	}

	digits := x[:l-1]
	sum := int(x[l-1] - '0')

	for i, j := 0, len(digits)-1; i <= j; i++ {
		if char := digits[j-i]; !unicode.IsNumber(char) {
			return false
		} else if num := int(char - '0'); i%2 != 0 {
			sum += num
		} else if dbl := num * 2; dbl > 9 {
			sum += dbl - 9
		} else {
			sum += dbl
		}
	}

	return sum%10 == 0
}
