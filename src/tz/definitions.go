package main

import (
	"time"
	"regexp"
	"strconv"
	"bytes"
	"strings"
)

type logMessage string
type timeStamp time.Time


//`[0-9]{4}\/[0-9]{2}\/[0-9]{2}\ [0-9]{2}\:+[0-9]{2}\:+[0-9]{2}\.+[0-9]{6}`

func createRegex(layout string) ([]string) {
	//regx := `[0-9]{4}\/[0-9]{2}\/[0-9]{2}\ [0-9]{2}\:+[0-9]{2}\:+[0-9]{2}\.+[0-9]{6}`
	var counters []int
	var delimiters []string
	delC := 0
	for _, r := range layout {
		r := string(r)
		if _, err := strconv.Atoi(r); err == nil {
			delC++
		} else {
			counters = append(counters, delC)
			delimiters = append(delimiters, r)
			delC = 0
		}
	}
	//var regx []rune
	var regx bytes.Buffer
	for j := 0; j < len(counters); j++ {
		if j == (len(counters) - 1) {
			regx.WriteString("[0-9]{")
			regx.WriteString(string(counters[j]))
			regx.WriteString("}")
		} else {
			regx.WriteString("[0-9]{")
			regx.WriteString(string(counters[j]))
			regx.WriteString("}\\")
			regx.WriteString(string(delimiters[j]))
		}
	}
	return strings.Fields(regx.String())
}

func (l logMessage) getTS(layout string) (timeStamp, error) {
	rgx := createRegex(layout)
	//if errR != nil {
	//	return timeStamp(time.Time{}), errors.New("could not make regex out of given layout")
	//}

	tsReg, _ := regexp.Compile(strings.Join(rgx, " "))
	tsStr := tsReg.FindString(string(l))
	ts, err := time.Parse(layout, tsStr)

	if err != nil {
		return timeStamp(time.Time{}), err
	}
	return timeStamp(ts), nil
}

// location "Europe/Ljubljana"

func (t timeStamp) changeTo(location string) (timeStamp, error) {
	loc, err := time.LoadLocation(location)
	if err == nil {
		return timeStamp(time.Time(t).In(loc)), nil
	} else {
		return timeStamp(time.Time{}), err
	}
}

