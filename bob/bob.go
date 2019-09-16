// Package bob is an Exercism teenager
// Heavily based on @paiv solution, thanks mate!
package bob

import (
	"regexp"
	"strings"
)

var (
	letters       = regexp.MustCompile(`\pL`)
	lower         = regexp.MustCompile(`\p{Ll}`)
	interrogation = regexp.MustCompile(`\?$`)
)

// Hey is a normal teenager bot
func Hey(remark string) string {
	entry := strings.TrimSpace(remark)

	question := interrogation.MatchString(entry)
	letters := letters.MatchString(entry)
	lowercase := lower.MatchString(entry)
	shout := letters && !lowercase

	switch {
	case shout && question:
		return "Calm down, I know what I'm doing!"
	case shout:
		return "Whoa, chill out!"
	case question:
		return "Sure."
	case entry == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
