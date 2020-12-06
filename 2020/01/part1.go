package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./part1.go <input file name>")
		os.Exit(1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("usage: go run ./part1.go <input file name>")
		os.Exit(1)
	}

	expenses := make(map[int]bool)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if _, ok := expenses[2020 - num]; ok {
			fmt.Println(num * (2020 - num))
			return
		}

		expenses[num] = true
	}

	panic("couldn't find two expenses that add up to 2020")
}
