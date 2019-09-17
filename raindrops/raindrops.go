package raindrops

import (
	"strconv"
	"strings"
)

var factors = map[string]int{
	"Pling": 3,
	"Plang": 5,
	"Plong": 7,
}

func Convert(input int) string {
	var builder strings.Builder
	for k, v := range factors {
		if mod := input % v; mod == 0 {
			builder.WriteString(k)
		}
	}

	if builder.Len() > 0 {
		return builder.String()
	} else {
		return strconv.Itoa(input)
	}
}
