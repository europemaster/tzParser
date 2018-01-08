package tz

import (
	"testing"
)

func TestTs(t *testing.T) {
	log := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	if _, err := log.getTS("2017/01/02 10:04:15.000000"); err != nil {
		t.Fatal("couldn't get timestamp")
	}
}
