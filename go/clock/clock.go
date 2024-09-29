package clock

import "fmt"

const minutesInHour = 60
const hoursInDay = 24

type Clock struct {
	hours, minutes int
}

func New(h, m int) Clock {
	return Clock{0, 0}.Add(m + h*minutesInHour)
}

func (c Clock) Add(m int) Clock {
	hoursToAdd := m / minutesInHour
	m -= hoursToAdd * minutesInHour
	c.hours += hoursToAdd
	c.minutes += m
	if c.minutes < 0 {
		c.minutes = minutesInHour + c.minutes%minutesInHour
		c.hours--
	} else if c.minutes >= minutesInHour {
		c.hours++
		c.minutes = c.minutes % minutesInHour
	}
	if c.hours < 0 {
		c.hours = hoursInDay + c.hours%hoursInDay
	}
	c.hours %= hoursInDay
	return c
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
