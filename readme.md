# UQ Assessment To Calendar

A Go program to save assessment information from a UQ ECP (Electronic Course Profile) to a calendar.

## Features

- [x] Every type of assessment except for exams
- [x] Semester, Delivery Mode, Location selection
- [x] Assessments with ranges for dates and multiple dates
- [x] Essentially every type of date format
- [ ] Figuring out dates for exams
- [ ] Allow inputting a course list and creating a single calendar
- [ ] Useful error messages and error-handling

## How to Use

- Clone the repo and build using make
- To get the assessment dates for a course in the current semester, in the flexible delivery mode and for the St. Lucia campus, call `uq_a2c course_code`. e.g. for COMP3400 call `uq_a2c COMP3400`
- Import the .ics file it outputs to your favourite calendar

For more comprehensive options and usage see below:
```
Usage:

uq_a2c (-h|--help)
uq_a2c course_code [options]

Options:
-s --semester	Semester to get assessments for in the format "Semester <n>, <year>"
-l --location	Location this course was held one of {"St Lucia", "Gatton"}
-d --delivery	Delivery mode of the course {"Flexible Delivery", "Internal", "External"}
-o --output     File to output the produced calendar to (note: include the ".ics" extension)
```