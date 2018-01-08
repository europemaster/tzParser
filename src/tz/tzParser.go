package tz

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"errors"
)

func main()  {
	flag.Parse()
	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		file = os.Stdin
	}
	readF := bufio.NewReader(file)
	for {
		singleByt, err := readF.ReadBytes('\n')
		if err != nil {
			errors.New("Error reading")
			break
		}
		logM := logMessage(singleByt)
		//change format to go standard
		ts, err := logM.getTS("2006/01/02 15:04:05.999999")
		if err != nil {
			fmt.Println(err)
		}
		tsLocal, err2 := ts.changeTo("Europe/Ljubljana")
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println("Original log: %s", logM)
		fmt.Println("Extracted timestamp: %s", ts)
		fmt.Println("Localized timestamp: ", tsLocal)
	}



	//log := logMessage("2017/12/28 17:10:57.194 UTC (3fcfdf5f.gateway.release_4.2.1.89+cg01961) [info] Didn't find bootloader for '58003000-1351-3432-3434-373300000000', [0 4 0], [3 18 1868], 5")
	//ts, err := log.getTS("2017/01/02 10:04:15.000000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//tsLocal, err2 := ts.changeTo("Europe/Ljubljana")
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//fmt.Println("Original log: %s", log)
	//fmt.Println("Extracted timestamp: %s", ts)
	//fmt.Println("Localized timestamp: ", tsLocal)

	//flag.Parse()
	//
	//timeFormat := flag.Arg(0)
	//from := flag.Arg(1)
	//to := flag.Arg(2)
	//fmt.Println(timeFormat, from, to)
}
