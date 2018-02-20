package main

import (
	"time"
	"regexp"
	"strconv"
	"bytes"
	"strings"
)

type logMessage string

func createRegex(layout string, tZone bool) ([]string) {
	var counters []int
	var delimiters []string
	delC := 0
	for i, r := range layout {
		r := string(r)
		if _, err := strconv.Atoi(r); err == nil {
			if i == len(layout) - 1 {
				delC++
				counters = append(counters, delC)
			} else {
				delC++
			}
		} else {
			counters = append(counters, delC)
			delimiters = append(delimiters, r)
			delC = 0
		}
	}
	var regx bytes.Buffer
	for ind, c := range counters {
		if ind == (len(counters) - 1) {
			regx.WriteString("[0-9]{")
			regx.WriteString(strconv.Itoa(c))
			regx.WriteString("}")
			break
		} else {
			regx.WriteString("[0-9]{")
			regx.WriteString(strconv.Itoa(c))
			regx.WriteString("}\\")
			regx.WriteString(string(delimiters[ind]))
		}
	}
	if tZone == true {
		regx.WriteString("\\s+[A-Z]{3}")
	}
	return strings.Fields(regx.String())
}

func (l logMessage) getTS(layout string) (time.Time, error) {
	rgx := createRegex(layout, false)

	tsReg, _ := regexp.Compile(strings.Join(rgx, " "))
	tsStr := tsReg.FindString(string(l))
	ts, err := time.Parse(layout, tsStr)

	if err != nil {
		return time.Time{}, err
	}
	return ts, nil
}


func changeTo(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err == nil {
		return time.Time(t).In(loc), nil
	} else {
		return time.Time{}, err
	}
}

//replace old ts with new one
func (l logMessage) generate(layout string, location string) (logMessage, error) {
	regx := createRegex(layout, true)
	regex := strings.Join(regx, " ")
	ts, err := l.getTS(layout)
	if err != nil {
		return "", err
	}
	newTs, err := changeTo(ts, location)
	r, _ := regexp.Compile(regex)
	return logMessage(r.ReplaceAllString(string(l), newTs.String())), nil
}