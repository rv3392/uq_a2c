package calendar

import (
	"strconv"
	"testing"

	"uq_a2c/coursescraper"

	"github.com/jordic/goics"
)

func generateTestCaseOne() []coursescraper.Assessment {
	testAssessment1 := coursescraper.Assessment{
		Name:    "Assignment 1",
		Format:  "Problem Set/s",
		DueDate: "15 Mar 21 13:00",
		Weight:  "",
		Description: "The assesment is broken into two parts, equally weighted:Written submission. " +
			"Mainly proofs and answering theory questions.  Human evaluated.      Code submission.  Haskell " +
			"code automatically graded.Students will have one weekto complete assessments at home.",
	}

	testAssessment2 := coursescraper.Assessment{
		Name:    "Assignment 2",
		Format:  "Problem Set/s",
		DueDate: "18 Apr 21 13:00",
		Weight:  "",
		Description: "The assesment is broken into two parts, equally weighted:Written submission. " +
			"Mainly proofs and answering theory questions.  Human evaluated.      Code submission.  Haskell " +
			"code automatically graded.Students will have one weekto complete assessments at home.",
	}

	testAssessments := []coursescraper.Assessment{testAssessment1, testAssessment2}

	return testAssessments
}

func TestCreateAssessmentsCal(t *testing.T) {
	testAssessments := generateTestCaseOne()
	cal := CreateAssessmentsCal(testAssessments)

	for i := 0; i < 2; i++ {
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

		if calComp.Properties["DUE"] != testAssessments[i].DueDate {
			t.Errorf("Component " + strconv.Itoa(i) + " of type " + calComp.Tipo +
				" had the incorrect DUE property:\nExpected: " + testAssessments[i].DueDate +
				"\nWas: " + calComp.Properties["DUE"])
		}
	}
}
