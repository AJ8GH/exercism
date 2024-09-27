package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	v, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err.Error()))
	}
	return v
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	v, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err.Error()))
	}
	return time.Now().After(v)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	v, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err.Error()))
	}
	return v.Hour() > 11 && v.Hour() <= 19
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsed := Schedule(date)
	formatted := parsed.Format("Monday, January 2, 2006, at 15:04.")
	return "You have an appointment on " + formatted
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	v, err := time.Parse(
		"2006-01-02 15:04:05 +0000 UTC",
		"2024-09-15 00:00:00 +0000 UTC",
	)

	if err != nil {
		fmt.Println(fmt.Errorf("error: %s", err.Error()))
	}

	return v
}
