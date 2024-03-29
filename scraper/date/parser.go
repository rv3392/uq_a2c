// Package date provides functions
package date

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

// AssessmentDate is a node of a date tree i.e. a tree to store the assessment dates
// of a particular assessment. Each node has 0 or more children with non-leaf nodes
// representing either a range or multi-date and leaf nodes representing a particular
// date/time itself.
//
// An AssessmentDate has a DateType which represents whether this is a leaf node
// (SingleDateType) or a non-leaf node (MultiDateType, RangeDateType). Only an
// AssessmentDate of SingleDateType is allowed to have a dateValue and this must
// be set using SetDateValue and retrieved using GetDateValue.
type AssessmentDate struct {
	ChildDates []*AssessmentDate
	DateType   AssessmentDateType
	dateValue  time.Time
}

type AssessmentDateType int

const (
	SingleDateType AssessmentDateType = iota
	MultiDateType  AssessmentDateType = iota
	RangeDateType  AssessmentDateType = iota
)

// ErrTime is the unix timestamp representing 1/1/0 00:00:00 and is used to represent an error
const ErrTime = -62167255200

func ParseString(dateToParse string) *AssessmentDate {
	return matchDates(dateToParse)
}

func (d *AssessmentDate) ToString() string {
	retStr := ""
	switch d.DateType {
	case SingleDateType:
		retStr = d.GetDateValue().Format("20060102T150405Z")
	case RangeDateType:
		start := d.ChildDates[0].ToString()
		end := d.ChildDates[1].ToString()
		retStr = start + " - " + end
	case MultiDateType:
		for _, child := range d.ChildDates {
			retStr += child.ToString()
			retStr += ", "
		}
		retStr = retStr[:len(retStr) - 2]

	default:
	}

	return retStr
}

func NewAssessmentDate(t AssessmentDateType, children []*AssessmentDate) *AssessmentDate {
	return &AssessmentDate{DateType: t, ChildDates: children}
}

func (d *AssessmentDate) SetDateValue(t time.Time) error {
	if d.DateType == SingleDateType {
		d.dateValue = t
		return nil
	}

	return errors.New("Warning: Can't set date for AssessmentDate with DateType not equal to singleDateType. Are you sure you want to do this?")
}

func (d *AssessmentDate) GetDateValue() time.Time {
	return d.dateValue
}

func matchDates(dateToParse string) *AssessmentDate {
	var datesFound []*AssessmentDate
	var thisDate *AssessmentDate

	dateType, splitDate := splitDateStr(dateToParse)
	if dateType != SingleDateType {
		for _, date := range splitDate {
			fmt.Print(date)
			datesFound = append(datesFound, matchDates(date))
		}

		thisDate = NewAssessmentDate(dateType, datesFound)
		return thisDate
	}

	thisDate = NewAssessmentDate(dateType, make([]*AssessmentDate, 0))
	res, _ := matchDate(splitDate[0])

	thisDate.SetDateValue(res)
	return thisDate
}

func splitDateStr(toSplit string) (AssessmentDateType, []string) {
	var dateType AssessmentDateType
	var dateParts []string

	multiDateParts := strings.Split(toSplit, ",")
	rangeDateParts := strings.Split(toSplit, "-")
	if len(multiDateParts) > 1 {
		dateType = MultiDateType
		dateParts = multiDateParts
	} else if len(rangeDateParts) > 1 {
		dateType = RangeDateType
		dateParts = rangeDateParts
	} else {
		dateType = SingleDateType
		dateParts = []string{toSplit}
	}

	for i := range dateParts {
		dateParts[i] = strings.TrimSpace(dateParts[i])
	}

	return dateType, dateParts
}

func matchDate(dateToParse string) (time.Time, bool) {
	if dateToParse == "" {
		return time.Unix(ErrTime, 0), false
	}

	tz, _ := time.LoadLocation("Australia/Brisbane")
	res, err := dateparse.ParseIn(dateToParse, tz, dateparse.PreferMonthFirst(false))
	if err != nil {
		return time.Unix(ErrTime, 0), false
	}

	return res, true
}
