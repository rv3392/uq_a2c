package calendar

import (
	"errors"
	"fmt"
	"os"
	"uq_a2c/scraper"
	"uq_a2c/scraper/date"

	goics "github.com/jordic/goics"
)

// CreateAssessmentsCal creates an ics file with the given assessmentsToSave slice
// and returns a pointer to an ics.Calendar type.
func CreateAssessmentsCal(assessments []scraper.Assessment) *goics.Component {
	cal := createEmptyAssessmentsCal()
	for _, a := range assessments {
		err, events := createEventsFromAssessment(&a)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to write assessments to ics file!")
			error.Error(err)
			os.Exit(2)
		}

		for _, e := range events {
			cal.AddComponent(e)
		}
	}

	return cal
}

func createEmptyAssessmentsCal() *goics.Component {
	// Create a new assessments cal with some basic properties
	cal := goics.NewComponent()
	cal.SetType("VCALENDAR")
	cal.AddProperty("CALSCAL", "GREGORIAN")
	cal.AddProperty("PRODID", "-//richalverma.com/")

	return cal
}

// Recurse through the given Assessment's DueDate until a SingleDateType or RangeDateType is reached
func createEventsFromAssessment(a *scraper.Assessment) (error, []*goics.Component) {
	events := make([]*goics.Component, 0)
	var err error = nil
	switch a.DueDate.DateType {
	case date.RangeDateType:
		err = addEventFromRangeDate(a, &events)
	case date.SingleDateType:
		err = addEventFromSingleDate(a, &events)
	case date.MultiDateType:
		for _, d := range a.DueDate.ChildDates {
			newA := a
			newA.DueDate = d

			var subEvents []*goics.Component
			err, subEvents = createEventsFromAssessment(newA)
			events = append(events, subEvents...)
		}
	}

	if err != nil {
		return err, nil
	}

	return nil, events
}

func addEventFromRangeDate(a *scraper.Assessment, events *[]*goics.Component) error {
	if a.DueDate.DateType != date.RangeDateType {
		return errors.New("warning: can't convert non range date type to VEVENT")
	}

	start := a.DueDate.ChildDates[0].ToString()
	end := a.DueDate.ChildDates[1].ToString()
	*events = append(*events, newEvent(a.Name, a.Description, start, end))

	return nil
}

// Create a new assessment todo with the details specified in the assessment struct
// If the given Assessment struct contains a DueDate with type MultiDateType then an error is returned.
func addEventFromSingleDate(a *scraper.Assessment, events *[]*goics.Component) error {
	if a.DueDate.DateType != date.SingleDateType {
		return errors.New("warning: can't convert non single date type to VEVENT")
	}
	*events = append(*events, newEvent(a.Name, a.Description, a.DueDate.ToString(), a.DueDate.ToString()))
	return nil
}

func newEvent(name string, description string, start string, end string) *goics.Component {
	event := goics.NewComponent()
	event.SetType("VEVENT")
	event.AddProperty("DTSTART", start)
	event.AddProperty("DTEND", end)
	event.AddProperty("SUMMARY", name)
	event.AddProperty("DESCRIPTION", description)
	return event
}
