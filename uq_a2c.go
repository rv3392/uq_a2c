package main

import (
	"fmt"
	"os"

	calendar "uq_a2c/cal"
	coursescraper "uq_a2c/scraper"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Semester string `short:"s" long:"semester" description:"semester of the course" default:"Semester 1, 2021"`
	Location string `short:"l" long:"location" description:"delivery location of the course" default:"St Lucia"`
	Delivery string `short:"d" long:"delivery" description:"delivery mode of the course" default:"Flexible Delivery"`
	FileName string `short:"o" long:"output" description:"the name of the output file"`
}

var usage = `Usage: 

uq_a2c (-h|--help)
uq_a2c course_code [options]

Options:
-s --semester	Semester to get assessments for in the format "Semester <n>, <year>"
-l --location	Location this course was held one of {"St Lucia", "Gatton"}
-d --delivery	Delivery mode of the course {"Flexible Delivery", "Internal", "External"}
-o --output		File to output the produced calendar to (note: include the ".ics" extension)
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
	course := os.Args[1]

	_, err := flags.ParseArgs(&opts, os.Args[2:])
	if err != nil {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

	assessments := coursescraper.ScrapeAssessments(course, opts.Semester, opts.Location, opts.Delivery)
	assessmentCal := calendar.CreateAssessmentsCal(assessments)

	var outputFile string
	if opts.FileName == "" {
		outputFile = course + ".ics"
	} else {
		outputFile = opts.FileName
	}

	calendar.Save(outputFile, assessmentCal)
}
