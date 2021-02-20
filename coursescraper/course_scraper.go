// Package coursescraper provides methods to scrape course profiles and return
// data.
package coursescraper

import (
	"time"

	"github.com/gocolly/colly/v2"
)

// Assessment represents a single Assessment and all associated data.
type Assessment struct {
	name        string
	format      string
	dueDate     time.Time
	weight      int
	description string
}

// ScrapeAssessments scrapes the assessment data into a slice of Assessment
// structs. Scraping is done using "colly".
func ScrapeAssessments(courseName string) []Assessment {
	c := colly.NewCollector()

	return []Assessment{} // Placeholder return value
}
