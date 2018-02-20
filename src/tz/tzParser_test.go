package main

import (
	"testing"
	"strings"
	"fmt"
)

func TestGetTs(t *testing.T) {
	log := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	if ts, err := log.getTS("2006/01/02 15:04:05.999"); err != nil {
		t.Fatal("couldn't get timestamp", err)
	} else {
		fmt.Println("get ts: ", ts)
	}
}

func TestChangeTo(t *testing.T) {
	log := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	ts, _ := log.getTS("2006/01/02 15:04:05.999")
	if tsRes, err := changeTo(ts, "Europe/Ljubljana"); err != nil {
		t.Fatal("couldn't convert", err)
	} else {
		fmt.Println("change to: ", tsRes)
	}
}

func TestCreateRegex(t *testing.T) {
	rgx := createRegex("2006/01/02 15:04:05.999", true)
	rgxP := strings.Join(rgx, " ")
	if rgxP != "[0-9]{4}\\/[0-9]{2}\\/[0-9]{2}\\ [0-9]{2}\\:[0-9]{2}\\:[0-9]{2}\\.[0-9]{3}\\s+[A-Z]{3}" {
		fmt.Println("rgx: ", rgxP)
		t.Fatal("couldn't create regexp")
	} else {
		fmt.Println("OK ", rgxP)
	}
}

func TestGenerate(t *testing.T) {
	logM := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	chLine, err := logM.generate("2006/01/02 15:04:05.999", "Europe/Ljubljana")
	if err != nil {
		t.Fatal("couldn't generate new log")
	} else if chLine != "2017/12/28 18:10:57.194 CET"{
		fmt.Println("Changed line: ", chLine)
	} else {
		fmt.Println("OK ", chLine)
	}
}
