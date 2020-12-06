package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func main() {
	input := OpenInputFile()

	expenses := make(map[int]bool)
	scanner := bufio.NewScanner(input)
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
