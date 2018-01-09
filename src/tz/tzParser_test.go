package main

import (
	"testing"
	"strings"
)

func TestGetTs(t *testing.T) {
	log := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	if _, err := log.getTS("2006/01/02 10:04:15.999999"); err != nil {
		t.Fatal("couldn't get timestamp", err)
	}
}

func TestChangeTo(t *testing.T) {
	log := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	ts, _ := log.getTS("2006/01/02 10:04:15.999999")
	if _, err := ts.changeTo("Europe/Ljubljana"); err != nil {
		t.Fatal("couldn't convert", err)
	}
}

func TestCreateRegex(t *testing.T) {
	rgx := createRegex("2006/01/02 10:04:15.999999")
	rgxP := strings.Join(rgx, " ")
	if rgxP != "[0-9]{4}\\/[0-9]{2}\\/[0-9]{2}\\ [0-9]{2}\\:+[0-9]{2}\\:+[0-9]{2}\\.+[0-9]{6}" {
		t.Fatal("couldn't create regexp")
	}

}