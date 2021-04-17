package calendar

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
	"uq_a2c/scraper"
	"uq_a2c/scraper/date"

	"github.com/jordic/goics"
)

func generateTestCaseOne() []scraper.Assessment {
	var err error

	date1 := date.NewAssessmentDate(date.SingleDateType, nil)
	err = date1.SetDateValue(time.Date(2021, time.March, 15, 13, 0, 0, 0, time.Local))

	date2 := date.NewAssessmentDate(date.SingleDateType, nil)
	err = date2.SetDateValue(time.Date(2021, time.April, 18, 13, 0, 0, 0, time.Local))

	if err != nil {
		fmt.Errorf("%s", error.Error(err))
		os.Exit(0)
	}

	testAssessment1 := scraper.Assessment{
		Name:    "Assignment 1",
		Format:  "Problem Set/s",
		DueDate: date1,
		Weight:  "",
		Description: "The assesment is broken into two parts, equally weighted:Written submission. " +
			"Mainly proofs and answering theory questions.  Human evaluated.      Code submission.  Haskell " +
			"code automatically graded.Students will have one weekto complete assessments at home.",
	}


	testAssessment2 := scraper.Assessment{
		Name:    "Assignment 2",
		Format:  "Problem Set/s",
		DueDate: date2,
		Weight:  "",
		Description: "The assesment is broken into two parts, equally weighted:Written submission. " +
			"Mainly proofs and answering theory questions.  Human evaluated.      Code submission.  Haskell " +
			"code automatically graded.Students will have one weekto complete assessments at home.",
	}

	return []scraper.Assessment{testAssessment1, testAssessment2}
}

func TestCreateAssessmentsCal_GivenSingleDateType_CorrectValues(t *testing.T) {
	testAssessments := generateTestCaseOne()
	cal := CreateAssessmentsCal(testAssessments)

	for i := 0; i < len(testAssessments); i++ {
		calComp, isComponent := cal.Elements[i].(*goics.Component)
		if !isComponent {
			t.Errorf("A component of cal is not a *goics.Component")
		}

		if calComp.Properties["SUMMARY"] != testAssessments[i].Name {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect SUMMARY property:\nExpected: " + testAssessments[i].Name +
				"\nWas: " + calComp.Properties["SUMMARY"])
		}

		if calComp.Properties["DESCRIPTION"] != testAssessments[i].Description {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DESCRIPTION property:\nExpected: " + testAssessments[i].Description +
				"\nWas: " + calComp.Properties["DESCRIPTION"])
		}

		if calComp.Properties["DTSTART"] != testAssessments[i].DueDate.ToString() {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DUE property:\nExpected: " + testAssessments[i].DueDate.ToString() +
				"\nWas: " + calComp.Properties["DTSTART"])
		}

		if calComp.Properties["DTEND"] != testAssessments[i].DueDate.ToString() {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DUE property:\nExpected: " + testAssessments[i].DueDate.ToString() +
				"\nWas: " + calComp.Properties["DTEND"])
		}
	}
}

func generateTestCaseTwo() []scraper.Assessment {
	var err error

	startDate1 := date.NewAssessmentDate(date.SingleDateType, nil)
	endDate1 := date.NewAssessmentDate(date.SingleDateType, nil)
	date1 := date.NewAssessmentDate(date.RangeDateType, []*date.AssessmentDate{startDate1, endDate1})

	err = startDate1.SetDateValue(time.Date(2002, time.March, 15, 13, 0, 0, 0, time.Local))
	err = endDate1.SetDateValue(time.Date(2002, time.March, 15, 16, 35, 40, 0, time.Local))

	startDate2 := date.NewAssessmentDate(date.SingleDateType, nil)
	endDate2 := date.NewAssessmentDate(date.SingleDateType, nil)
	date2 := date.NewAssessmentDate(date.RangeDateType, []*date.AssessmentDate{startDate2, endDate2})

	err = startDate1.SetDateValue(time.Date(2021, time.March, 10, 12, 0, 0, 0, time.Local))
	err = startDate2.SetDateValue(time.Date(2002, time.March, 15, 12, 0, 0, 0, time.Local))

	if err != nil {
		fmt.Errorf("%s", error.Error(err))
		os.Exit(0)
	}

	testAssessment1 := scraper.Assessment{
		Name:    "Assignment 1",
		Format:  "Problem Set/s",
		DueDate: date1,
		Weight:  "",
		Description: "The assesment is broken into two parts, equally weighted:Written submission. " +
			"Mainly proofs and answering theory questions.  Human evaluated.      Code submission.  Haskell " +
			"code automatically graded.Students will have one weekto complete assessments at home.",
	}


	testAssessment2 := scraper.Assessment{
		Name:    "Assignment 2",
		Format:  "Problem Set/s",
		DueDate: date2,
		Weight:  "",
		Description: "The assesment is broken into two parts, equally weighted:Written submission. " +
			"Mainly proofs and answering theory questions.  Human evaluated.      Code submission.  Haskell " +
			"code automatically graded.Students will have one weekto complete assessments at home.",
	}

	return []scraper.Assessment{testAssessment1, testAssessment2}
}

func TestCreateAssessmentsCal_GivenRangeDateType_CorrectValues(t *testing.T) {
	testAssessments := generateTestCaseTwo()
	cal := CreateAssessmentsCal(testAssessments)

	for i := 0; i < len(testAssessments); i++ {
		calComp, isComponent := cal.Elements[i].(*goics.Component)
		if !isComponent {
			t.Errorf("A component of cal is not a *goics.Component")
		}

		if calComp.Properties["SUMMARY"] != testAssessments[i].Name {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect SUMMARY property:\nExpected: " + testAssessments[i].Name +
				"\nWas: " + calComp.Properties["SUMMARY"])
		}

		if calComp.Properties["DESCRIPTION"] != testAssessments[i].Description {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DESCRIPTION property:\nExpected: " + testAssessments[i].Description +
				"\nWas: " + calComp.Properties["DESCRIPTION"])
		}

		if calComp.Properties["DTSTART"] != testAssessments[i].DueDate.ChildDates[0].ToString() {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DUE property:\nExpected: " + testAssessments[i].DueDate.ToString() +
				"\nWas: " + calComp.Properties["DTSTART"])
		}

		if calComp.Properties["DTEND"] != testAssessments[i].DueDate.ChildDates[0].ToString() {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DUE property:\nExpected: " + testAssessments[i].DueDate.ToString() +
				"\nWas: " + calComp.Properties["DTEND"])
		}
	}
}
