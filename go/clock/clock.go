package clock

import "fmt"

type clock struct {
	minutes int
}

func mod(n, m int) int {
	x := n % m
	if x < 0 {
		return m + x
	}
	return x
}

func New(hour, minute int) clock {
	return clock{mod(hour*60+minute, 1440)}
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func (c clock) Add(minutes int) clock {
	n := mod(c.minutes+minutes, 1440)
	return clock{n}
}

func (c clock) Subtract(minutes int) clock {
	return c.Add(-minutes)
}
