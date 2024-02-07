package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
   layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout,date)
    if err != nil {
        panic(err)
    }
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    foo, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
    return foo.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	foo, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	if foo.Hour() >= 12 && foo.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    foo := Schedule(date)
    return fmt.Sprintf("You have an appointment on %s.", foo.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().UTC().Year()
    anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)

    return anniversary
}
