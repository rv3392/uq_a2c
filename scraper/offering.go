package scraper

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type offering struct {
	semester string
	location string
	mode     string
}

func getCourseOfferingsPageURL(courseCode string) string {
	return courseOfferingsPage + courseCode
}

func cleanData(input string) string {
	return strings.Replace(strings.Replace(input, "\n", "", -1), "\t", "", -1)
}

func courseOfferingRowHTMLToOffering(tr *goquery.Selection) (offering, string) {
	data := tr.ChildrenFiltered("td")
	sem := cleanData(data.Slice(0, 1).Text())
	loc := cleanData(data.Slice(1, 2).Text())
	mode := cleanData(data.Slice(2, 3).Text())
	url, _ := data.Slice(3, 4).Children().Attr("href")

	return offering{sem, loc, mode}, cleanData(url)
}

func addHTMLOfferingsTableToMap(table *colly.HTMLElement, offerings map[offering]string) {
	table.DOM.Find("tbody").Find("tr").Each(func(i int, tr *goquery.Selection) {
		thisOffering, url := courseOfferingRowHTMLToOffering(tr)
		offerings[thisOffering] = url
	})
}

func initCourseOfferingsCollector(offerings map[offering]string) colly.Collector {
	c := colly.NewCollector()

	c.OnHTML("table[id=course-current-offerings]", func(e *colly.HTMLElement) {
		addHTMLOfferingsTableToMap(e, offerings)
	})

	c.OnHTML("table[id=course-archived-offerings]", func(e *colly.HTMLElement) {
		addHTMLOfferingsTableToMap(e, offerings)
	})

	return *c
}

func getCourseOfferings(courseCode string) map[offering]string {
	offerings := make(map[offering]string)
	c := initCourseOfferingsCollector(offerings)
	c.Visit(getCourseOfferingsPageURL(courseCode))

	return offerings
}
