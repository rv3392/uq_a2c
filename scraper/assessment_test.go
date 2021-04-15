package scraper

import (
	"fmt"
	"testing"
)

func TestScrapeAssessments_GivenS12021CourseWithExam_GetAllAssessmentsCorrectly(t *testing.T) {
	// ScrapeAssessments isn't a great to test, surely there's a better way of designing this
	assessments := ScrapeAssessments("COMP3400", "Semester 1, 2021", "St Lucia", "Flexible Delivery")
	for _, a := range assessments {
		fmt.Printf("%s", (&a).ToString())
	}
}
