package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"errors"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	file, err := os.Open(filename)
	if err != nil {
		file = os.Stdin
	}
	readF := bufio.NewReader(file)
	for {
		line, err := readF.ReadString('\n')
		if err != nil {
			errors.New("Error reading")
			break
		}
		fmt.Println(line)
		hi()
		//logM := logMessage(line)
		//newLine, err2 := logM.generate("2006/01/02 15:04:05.999", "Europe/Ljubljana")
		//if err2 != nil {
		//	errors.New("Error generating new string")
		//}
		//fmt.Println(newLine)
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
