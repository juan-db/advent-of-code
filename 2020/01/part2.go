package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run ./part1 <input file name>")
		os.Exit(1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("usage: go run ./part1 <input file name>")
		os.Exit(1)
	}

	expenses := make(map[int]bool)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		for k := range expenses {
			diff := 2020 - k - num
			if _, ok := expenses[diff]; ok {
				fmt.Println(diff * num * k)
				return
			}
		}

		expenses[num] = true
	}

	panic("couldn't find two expenses that add up to 2020")
}
