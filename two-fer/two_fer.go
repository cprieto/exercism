package twofer

import (
	"fmt"
	"strings"
)

func ShareWith(name string) string {
	other := strings.TrimSpace(name)
	if other == "" {
		other = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", other)
}
