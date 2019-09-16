// Package gigasecond is my solution to Exercism Gigasecond
package gigasecond

import "time"

// AddGigasecond Add 1 Gigasecond (10^9 seconds) to a date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
