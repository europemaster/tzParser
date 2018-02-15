package main

import (
	"time"
	"regexp"
	"strconv"
	"bytes"
	"strings"
	"fmt"
)

type logMessage string
type timeStamp time.Time

func hi() string {
	return "Hello"
}
func createRegex(layout string) ([]string) {
	//regx := `[0-9]{4}\/[0-9]{2}\/[0-9]{2}\ [0-9]{2}\:+[0-9]{2}\:+[0-9]{2}\.+[0-9]{6}`
	var counters []int
	var delimiters []string
	delC := 0
	for i, r := range layout {
		r := string(r)
		if _, err := strconv.Atoi(r); err == nil {
			//fmt.Println(r + " is number")
			if i == len(layout) - 1 {
				delC++
				counters = append(counters, delC)
			} else {
				delC++
			}
		} else {
			//fmt.Println(r + " is delimiter")
			counters = append(counters, delC)
			delimiters = append(delimiters, r)
			delC = 0
		}
	}
	var regx bytes.Buffer
	//fmt.Println(counters)
	//fmt.Println(delimiters)
	for ind, c := range counters {
		//fmt.Println("index: ", ind)
		if ind == (len(counters) - 1) {
			regx.WriteString("[0-9]{")
			regx.WriteString(strconv.Itoa(c))
			regx.WriteString("}")
			break
		} else {
			//fmt.Println("ELSE")
			//fmt.Println(delimiters[ind])
			regx.WriteString("[0-9]{")
			regx.WriteString(strconv.Itoa(c))
			regx.WriteString("}\\")
			regx.WriteString(string(delimiters[ind]))
			//fmt.Println(strings.Fields(regx.String()))
		}
	}
	return strings.Fields(regx.String())
}

func (l logMessage) getTS(layout string) (time.Time, error) {
	rgx := createRegex(layout)
	//if errR != nil {
	//	return timeStamp(time.Time{}), errors.New("could not make regex out of given layout")
	//}

	tsReg, _ := regexp.Compile(strings.Join(rgx, " "))
	//fmt.Println("tsReg: ", tsReg)
	tsStr := tsReg.FindString(string(l))
	//fmt.Println("tsStr: ",tsStr)
	ts, err := time.Parse(layout, tsStr)
	//fmt.Println("ts: ",reflect.TypeOf(ts))

	if err != nil {
		return time.Time{}, err
	}
	return ts, nil
}

// location "Europe/Ljubljana"

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
	regx := createRegex(layout)
	regex := strings.Join(regx, " ")
	//fmt.Println("regex", regex)
	//fmt.Println("logline", l)
	ts, err := l.getTS(layout)
	//fmt.Println("ts", ts)
	if err != nil {
		return "", err
	}
	newTs, err := changeTo(ts, location)
	r, _ := regexp.Compile(regex)
	fmt.Println(layout, location)
	return logMessage(r.ReplaceAllString(string(l), newTs.String())), nil
}