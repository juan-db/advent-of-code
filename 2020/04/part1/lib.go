package main

import (
	"bufio"
	"fmt"
	"os"
)

var UsageString = "usage: go run ./solution.go <input file name>"

func OpenInputFile() *os.File {
	if len(os.Args) < 2 {
		fmt.Println(UsageString)
		os.Exit(1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(UsageString)
		os.Exit(1)
	}

	return inputFile
}

func ReadByLine(file *os.File, handler func(string)) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		handler(line)
	}
}
