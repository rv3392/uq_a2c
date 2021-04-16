package main

import (
	"fmt"
	"os"

	calendar "uq_a2c/cal"
	coursescraper "uq_a2c/scraper"
)

const currentSemester = "Semester 1, 2021"

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Error: Incorrect number of arguments")
	}

	courseCode := args[0]

	assessments := coursescraper.ScrapeAssessments(courseCode, currentSemester, "St Lucia", "Flexible Delivery")
	assessmentCal := calendar.CreateAssessmentsCal(assessments)
	calendar.Save(courseCode+".ics", assessmentCal)
}
