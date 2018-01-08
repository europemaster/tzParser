package tz

import (
	"time"
	"regexp"
	"errors"
	"strconv"
	"koala2/src/golang.org/x/net/html/atom"
)

type logMessage string
type timeStamp time.Time


//`[0-9]{4}\/[0-9]{2}\/[0-9]{2}\ [0-9]{2}\:+[0-9]{2}\:+[0-9]{2}\.+[0-9]{6}`

func createRegex(layout string) (string, error) {
	//regx := `[0-9]{4}\/[0-9]{2}\/[0-9]{2}\ [0-9]{2}\:+[0-9]{2}\:+[0-9]{2}\.+[0-9]{6}`
	var counters []int
	var delimiters []string
	delC := 0
	for i := 0; i < len(layout); i++ {
		curS := string(layout[i])
		if _, err := strconv.Atoi(curS); err == nil {
			delC++
		} else {
			counters = append(counters, delC)
			delimiters = append(delimiters, curS)
			delC = 0
		}
	}
	var regx []rune
	for j := 0; j< len(delimiters); j++ {
		regx = append(regx, '[0-9]')
	}
}

func (l logMessage) getTS(layout string) (timeStamp, error) {
	rgx, errR := createRegex(layout)
	if errR != nil {
		return timeStamp(time.Time{}), errors.New("could not make regex out of given layout")
	}

	tsReg, _ := regexp.Compile(rgx)
	tsStr := tsReg.FindString(string(l))
	ts, err := time.Parse(layout, tsStr)

	if err != nil {
		return timeStamp(time.Time{}), errors.New("could not convert")
	}
	return timeStamp(ts), nil
}

// location "Europe/Ljubljana"

func (t timeStamp) changeTo(location string) (timeStamp, error) {
	loc, err := time.LoadLocation(location)
	if err == nil {
		return timeStamp(t.In(loc)), nil
	} else {
		return timeStamp(time.Time{}), errors.New("could not load location")
	}


	//runes := []rune(deviation)
	//sign := string(runes[0:1])
	//num, err := strconv.Atoi(string(runes[1:2]))
	//
	//if err != nil {
	//	return time.Time{}, err
	//}
	//if sign == "+" {
	//
	//} else if sign == "-" {
	//
	//} else {
	//	return time.Time{}, err
	//}

}

