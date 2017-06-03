package clock

import (
	"fmt"
	"strings"
)

const testVersion = 4

// Clock type set here
type Clock struct {
	hour   int
	minute int
}

// New creates a new instance of clock type, returns clock.
func New(hour, minute int) Clock {
	var retHour int

	minute, retHour = calcMin(minute)
	//add or subtract returned hours from calculated minutes
	hour = calcHour(hour + retHour)

	return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {

	h := fmt.Sprintf("%02d", c.hour)
	m := fmt.Sprintf("%02d", c.minute)
	t := []string{h, m}

	return strings.Join(t, ":")
}

// Add function calculates the new time on clock passed in
// and returns clock with adjusted time.
func (c Clock) Add(minutes int) Clock {
	var retHour int

	c.minute, retHour = calcMin(c.minute + minutes)
	//add or subtract returned hours from calculated minutes
	c.hour = calcHour(c.hour + retHour)
	return c
}

// calcMin calculates hours and remaining minutes
// from raw minutes passed in.
func calcMin(minute int) (int, int) {
	var newHours, hour int
	hour = 0
	switch {
	case minute > 59:
		newHours = minute / 60
		hour = newHours
		minute = minute % 60
	case minute < 0 && minute > -61:
		minute = 60 + minute
		hour = -1
	case minute < -61:
		remMin := minute % 60
		newHours = minute / 60
		if remMin == 0 {
			hour = newHours
			minute = remMin
		} else {
			//have partial negative hour so subtract an
			//additional hour from the hour calculation
			hour = newHours - 1
			minute = 60 + remMin
		}
	}
	return minute, hour
}

// calcHour reformats raw hours into 24 hour format.
func calcHour(hour int) int {
	switch {
	case hour > 23:
		hour = hour % 24
	case hour < 0 && hour >= -23:
		hour = 24 + hour
	case hour < -23:
		if hour%24 == 0 {
			hour = 0
		} else {
			hour = 24 + (hour % 24)
		}
	}

	return hour
}
