package dateparser

import (
	"testing"
	"time"
)

func TestParseString_GivenSingleDateTime_ParsesIntoStartAndEndDate(t *testing.T) {
	singleDate := "15 Mar 21 13:00"
	date := ParseString(singleDate)

	if date.DateType != SingleDateType {
		var wasStr string
		if date.DateType == MultiDateType {
			wasStr = "MultiDateType"
		} else {
			wasStr = "RangeDateType"
		}
		t.Errorf("Incorrect parsed type at Root Node\nExpected:SingleDateType\nWas:%s", wasStr)
	}

	expectedDate, _ := time.Parse(time.RFC3339, "2021-03-15T13:00:00+10:00")
	if date.GetDateValue() != expectedDate {
		t.Errorf("Incorrect parsed date \nExpected:%s\nWas:%s", "2021-03-15T13:00:00+10:00", date.GetDateValue())
	}
}

func TestParseString_GivenDateRange_ParsesIntoStartAndEndDate(t *testing.T) {
	dualDate := "15 Jul 21 09:00 - 15 Jul 21 12:00"
	date := ParseString(dualDate)

	if date.DateType != RangeDateType {
		var wasStr string
		if date.DateType == SingleDateType {
			wasStr = "SingleDateType"
		} else {
			wasStr = "MultiDateType"
		}
		t.Errorf("Incorrect parsed type at Root Node\nExpected:RangeDateType\nWas:%s", wasStr)
	}

	if len(date.ChildDates) != 2 {
		t.Errorf("Range date parsed incorrectly - incorrect number of children \nExpected:%s\nWas:%d", "2", len(date.ChildDates))
	}

	startDate := date.ChildDates[0]
	if startDate.DateType != SingleDateType {
		var wasStr string
		if startDate.DateType == MultiDateType {
			wasStr = "MultiDateType"
		} else {
			wasStr = "RangeDateType"
		}
		t.Errorf("Incorrect parsed type at Root Node\nExpected:SingleDateType\nWas:%s", wasStr)
	}

	expectedDate, _ := time.Parse(time.RFC3339, "2021-07-15T09:00:00+10:00")
	if startDate.GetDateValue() != expectedDate {
		t.Errorf("Incorrect parsed date \nExpected:%s\nWas:%s", "2021-07-15T09:00:00+10:00", startDate.GetDateValue())
	}

	endDate := date.ChildDates[1]
	if endDate.DateType != SingleDateType {
		var wasStr string
		if endDate.DateType == MultiDateType {
			wasStr = "MultiDateType"
		} else {
			wasStr = "RangeDateType"
		}
		t.Errorf("Incorrect parsed type at Root Node\nExpected:SingleDateType\nWas:%s", wasStr)
	}

	expectedDate, _ = time.Parse(time.RFC3339, "2021-07-15T12:00:00+10:00")
	if endDate.GetDateValue() != expectedDate {
		t.Errorf("Incorrect parsed date \nExpected:%s\nWas:%s", "2021-07-15T12:00:00+10:00", endDate.GetDateValue())
	}
}

func TestParseString_GivenMultipleDates_ParsesIntoStartAndEndDate(t *testing.T) {
	multiDate := "23/08/20, 07/09/20, 25/09/20, 28/10/20"
	date := ParseString(multiDate)

	if date.DateType != MultiDateType {
		var wasStr string
		if date.DateType == SingleDateType {
			wasStr = "SingleDateType"
		} else {
			wasStr = "RangeDateType"
		}
		t.Errorf("Incorrect parsed type at Root Node\nExpected:MultiDateType\nWas:%s", wasStr)
	}

	if len(date.ChildDates) != 4 {
		t.Errorf("Range date parsed incorrectly - incorrect number of children \nExpected:%s\nWas:%d", "4", len(date.ChildDates))
	}

	expectedDates := []string{"2020-08-23T00:00:00+10:00", "2020-09-07T00:00:00+10:00", "2020-09-25T00:00:00+10:00", "2020-10-28T00:00:00+10:00"}
	for i, d := range date.ChildDates {
		if d.DateType != SingleDateType {
			var wasStr string
			if d.DateType == MultiDateType {
				wasStr = "MultiDateType"
			} else {
				wasStr = "RangeDateType"
			}
			t.Errorf("Incorrect parsed type at Root Node\nExpected:SingleDateType\nWas:%s", wasStr)
		}

		expectedDate, _ := time.Parse(time.RFC3339, expectedDates[i])
		if d.GetDateValue() != expectedDate {
			t.Errorf("Incorrect parsed date \nExpected:%s\nWas:%s", expectedDates[i], d.GetDateValue())
		}
	}
}
