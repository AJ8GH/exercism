// Package gigasecond provides gigasecond related time operations.
package gigasecond

import "time"

// AddGigasecond adds gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1_000_000_000)
}
