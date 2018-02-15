package main

import (
	"flag"
	"os"
	"bufio"
	"errors"
	"fmt"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	layoutFlag := flag.Arg(1)
	locationFlag := flag.Arg(2)
	file, err := os.Open(filename)
	if err != nil {
		file = os.Stdin
	}
	readF := bufio.NewReader(file)

	for {
		fmt.Println(locationFlag, layoutFlag)
		line, err := readF.ReadString('\n')
		if err != nil {
			errors.New("Error reading")
			break
		}
		fmt.Println("original line", line)
		logM := logMessage(line)
		//logM.generate("2006/01/02 15:04:05.999", "Europe/Ljubljana")
		newLine, err2 := logM.generate(layoutFlag, locationFlag)
		if err2 != nil {
			errors.New("Error generating new string")
		}
		fmt.Println("converted line", newLine)
	}


	//	//change format to go standard
	//	ts, err := logM.getTS("2006/01/02 15:04:05.999999")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	tsLocal, err2 := ts.changeTo("Europe/Ljubljana")
	//	if err2 != nil {
	//		fmt.Println(err2)
	//	}
	//	fmt.Println("Original log: %s", logM)
	//	fmt.Println("Extracted timestamp: %s", ts)
	//	fmt.Println("Localized timestamp: ", tsLocal)
	//}
}
