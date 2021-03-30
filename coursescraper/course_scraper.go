// Package coursescraper provides methods to scrape course profiles and return
// data.
package coursescraper

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

const courseOfferingsPage = "https://my.uq.edu.au/programs-courses/course.html?course_code="

// Assessment represents a single Assessment and all associated data.
type Assessment struct {
	name        string
	format      string
	dueDate     time.Time
	weight      int
	description string
}

func initAssessmentCollector() colly.Collector {
	c := colly.NewCollector()
	return *c
}

// ScrapeAssessments scrapes the assessment data into a slice of Assessment
// structs. Scraping is done using "colly".
func ScrapeAssessments(courseCode string) []Assessment {
	offerings := getCourseOfferings(courseCode)

	fmt.Println("Data:")
	for k, v := range offerings {
		fmt.Fprintf(os.Stdout, "%s, %s, %s = %s\n", k.semester, k.location, k.mode, v)
	}

	return []Assessment{} // Placeholder return value
}
