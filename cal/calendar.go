package calendar

import (
	"uq_a2c/coursescraper"

	ics "github.com/arran4/golang-ical"
)

// CreateAssessmentsCal creates an ics file with the given assessmentsToSave slice
// and returns a pointer to an ics.Calendar type.
func CreateAssessmentsCal(assessments []coursescraper.Assessment) *ics.Calendar {
	cal := ics.NewCalendar()

	return cal
}
