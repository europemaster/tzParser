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


//`[0-9]{4}\/[0-9]{2}\/[0-9]{2}\ [0-9]{2}\:+[0-9]{2}\:+[0-9]{2}\.+[0-9]{6}`

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

func (l logMessage) getTS(layout string) (timeStamp, error) {
	rgx := createRegex(layout)
	//if errR != nil {
	//	return timeStamp(time.Time{}), errors.New("could not make regex out of given layout")
	//}

	tsReg, _ := regexp.Compile(strings.Join(rgx, " "))
	fmt.Println("tsReg: ", tsReg)
	tsStr := tsReg.FindString(string(l))
	fmt.Println("tsStr: ",tsStr)
	ts, err := time.Parse(layout, tsStr)
	fmt.Println("ts: ",ts)

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

