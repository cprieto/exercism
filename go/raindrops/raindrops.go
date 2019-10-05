package raindrops

import (
	"strconv"
)

func Convert(input int) string {
	// NOTE: I tried using a map but nothing beats the performance of simple if's
	var output string
	if input % 3 == 0 {
		output += "Pling"
	}
	if input % 5 == 0 {
		output += "Plang"
	}
	if input % 7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return strconv.Itoa(input)
	} else {
		return output
	}
}
