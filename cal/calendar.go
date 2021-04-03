package calendar

import (
	"uq_a2c/coursescraper"

	goics "github.com/jordic/goics"
)

// CreateAssessmentsCal creates an ics file with the given assessmentsToSave slice
// and returns a pointer to an ics.Calendar type.
func CreateAssessmentsCal(assessments []coursescraper.Assessment) *goics.Component {
	cal := createEmptyAssessmentsCal()

	for _, assessment := range assessments {
		cal.AddComponent(createAssessmentTodo(cal, &assessment))
	}

	return cal
}

func createEmptyAssessmentsCal() *goics.Component {
	// Create a new assessments cal with some basic properties
	cal := goics.NewComponent()
	cal.SetType("VCALENDAR")
	cal.AddProperty("CALSCAL", "GREGORIAN")
	cal.AddProperty("PRODID", "-//tmpo.io/src/goics")

	return cal
}

func createAssessmentTodo(cal *goics.Component, assessment *coursescraper.Assessment) *goics.Component {
	// Create a new assessment todo with the detaisl specified in the assessment struct
	todo := goics.NewComponent()
	todo.SetType("VTODO")
	todo.AddProperty("DUE", assessment.DueDate)
	todo.AddProperty("SUMMARY", assessment.Name)
	todo.AddProperty("DESCRIPTION", assessment.Description)

	return todo
}
