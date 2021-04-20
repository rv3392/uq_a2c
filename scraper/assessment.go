// Package scraper provides methods to scrape course profiles and return
// data.
package scraper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"

	"uq_a2c/scraper/date"
)

const courseOfferingsPage = "https://my.uq.edu.au/programs-courses/course.html?course_code="

// Assessment represents a single Assessment and all associated data.
type Assessment struct {
	Name        string
	Format      string
	DueDate     *date.AssessmentDate
	Weight      string
	Description string
}

type parameterMap map[string]string

func parameterMapToAssessment(parameterValuePairs parameterMap) (Assessment, error) {
	if _, ok := parameterValuePairs["due_date"]; !ok {
		return Assessment{}, errors.New("assessment requires a date to be valid")
	}

	a := Assessment{Name: "", Format: "", Weight: "", Description: ""}
	for parameter, value := range parameterValuePairs {
		switch parameter {
		case "name":
			a.Name = value
		case "type":
			a.Format = value
		case "due_date":
			a.DueDate = date.ParseString(value)
		case "weight":
			a.Weight = value
		case "task_description":
			a.Description = value
		default:
			fmt.Fprintf(os.Stderr, "Warning: Found an invalid parameter while parsing an assessment: "+parameter+"\n")
		}
	}

	return a, nil
}

func removeWhitespace(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

func cleanDetailValueHTML(detHTML *goquery.Document) string {
	var retStr = ""
	retStr = strings.Replace(detHTML.Text(), detHTML.Find("strong").First().Text(), "", -1)
	return removeWhitespace(retStr)
}

func cleanDetailParamHTML(parHTML *goquery.Document) string {
	var retStr = ""
	retStr = strings.TrimSpace(parHTML.Find("strong").First().Text())
	retStr = strings.ToLower(retStr)
	retStr = strings.Replace(retStr, " ", "_", -1)
	retStr = strings.Replace(retStr, ":", "", -1)
	return removeWhitespace(retStr)
}

func parseHTMLToAssessment(assessmentHTML string) (Assessment, error) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(assessmentHTML))

	details := make(parameterMap)
	details["name"] = doc.Find("h4").First().Text()

	fmt.Fprintf(os.Stderr, "Info: Parsing details for "+details["name"]+"...")

	for _, line := range strings.Split(assessmentHTML, "<strong>") {
		fullyFormedLine := "<strong>" + line
		parameterDoc, _ := goquery.NewDocumentFromReader(strings.NewReader(fullyFormedLine))
		parameter := cleanDetailParamHTML(parameterDoc)
		value := cleanDetailValueHTML(parameterDoc)

		details[parameter] = value
	}

	fmt.Fprintf(os.Stderr, "\033[2D Finished!\n")

	a, err := parameterMapToAssessment(details)
	if err != nil {
		return Assessment{}, err
	}

	return a, nil
}

func initAssessmentCollector(assessments *[]Assessment) colly.Collector {
	c := colly.NewCollector()

	c.OnHTML("div[id=assessmentDetail]", func(details *colly.HTMLElement) {
		fmt.Printf("Found assessments div! Parsing assessments...\n")
		assessmentDetailsHTML, _ := details.DOM.Html()

		assessmentDetails := strings.Split(assessmentDetailsHTML, "<hr/>")
		for _, assessment := range assessmentDetails {
			a, err := parseHTMLToAssessment(assessment)
			if err != nil {
				continue
			}
			tempAssessments := append(*assessments, a)
			*assessments = tempAssessments
		}

		fmt.Printf("Parsing complete!\n")
	})

	return *c
}

func getAssessments(assessmentSectionURL string) []Assessment {
	assessments := make([]Assessment, 0)

	c := initAssessmentCollector(&assessments)
	c.Visit(assessmentSectionURL)

	fmt.Fprintf(os.Stderr, "\n\033[1mFound the following assessments: \033[0m\n")
	for _, a := range assessments {
		fmt.Print(a.ToString())
	}

	return assessments
}

func displayOfferingsLogging(courseCode string, offerings map[offering]string) {
	fmt.Printf("Offerings of %s: \n", courseCode)
	for k, v := range offerings {
		fmt.Fprintf(os.Stdout, "%s : %s\n", k.toString(), v)
	}
	fmt.Printf("\n")
}

// ScrapeAssessments scrapes the assessment data into a slice of Assessment
// structs. Scraping is done using "colly".
func ScrapeAssessments(courseCode string, semester string, location string, mode string) []Assessment {
	fmt.Printf("Fetching offerings...\n")
	offerings := getCourseOfferings(courseCode)
	displayOfferingsLogging(courseCode, offerings)

	fmt.Printf("Fetching assessments...\n")
	ecpURL := offerings[offering{semester: semester, location: location, mode: mode}]
	assessmentSectionURL := strings.Replace(ecpURL, "section_1", "section_5", 1)

	return getAssessments(assessmentSectionURL)
}

func (a *Assessment) ToString() string {
	return fmt.Sprintf("Task: %s\nType: %s\nDate: %s\nDescription: %s\n\n",
		a.Name, a.Format, a.DueDate.ToString(), a.Description)
}
