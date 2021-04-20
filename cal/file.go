// Package calendar provides functions to help with saving calendars
package calendar

import (
	"bufio"
	"os"

	goics "github.com/jordic/goics"
)

// Save saves the calendar data in "calendar" to the filename provided as
// a string. Include the ".ics" extension if wanted.
func Save(filename string, calendar *goics.Component) {
	f, _ := os.Create(filename)
	writer := bufio.NewWriter(f)
	calendar.Write(goics.NewICalEncode(writer))
	writer.Flush()
	f.Close()
}
