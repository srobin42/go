// Package gigasecond calculates the moment someone has lived a gigasecond
package gigasecond

import (
	"time"
)

const testVersion = 4
const gigasec = 1000000000

// AddGigasecond calculates the moment someone has lived
// 10^9 seconds returning the date.
func AddGigasecond(dt time.Time) time.Time {
	sec := time.Second * gigasec
	return dt.Add(sec)
}
