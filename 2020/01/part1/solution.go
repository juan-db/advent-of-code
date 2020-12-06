package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	file := OpenInputFile()

	expenses := make(map[int]bool)
	scanner := bufio.NewScanner(file)
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
