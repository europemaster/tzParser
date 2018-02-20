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
		line, err := readF.ReadString('\n')
		if err != nil {
			errors.New("error reading")
			break
		}
		fmt.Println("Original line: ", line)
		logM := logMessage(line)
		newLine, err2 := logM.generate(layoutFlag, locationFlag)
		if err2 != nil {
			errors.New("error generating new string")
		}
		fmt.Println("Converted line: ", newLine)
	}
}
