package clock

import (
	"fmt"
)

const testVersion = 4

// Clock type set here
type Clock int

// Copied this whole thing from another submission to learn from!
func New(hour, minute int) Clock {
	minute += hour * 60
	for minute < 0 {
		minute += 1440
	}
	return Clock(minute % 1440)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	for minutes < 0 {
		minutes += 1440
	}

	return (c + Clock(minutes)) % 1440
}
