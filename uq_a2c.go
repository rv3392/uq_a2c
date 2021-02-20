package main

import (
	"fmt"
	"os"

	"uq_a2c/coursescraper"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Error: Incorrect number of arguments")
	}

	assessments := coursescraper.ScrapeAssessments(args[0])
}
