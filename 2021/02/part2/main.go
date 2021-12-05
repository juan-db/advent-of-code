package main

import (
	"strconv"
	"strings"
)

func main() {
	position := 0
	depth := 0
	aim := 0

	ReadInputFileByLine(func(line string) {
		tokens := strings.Split(line, " ")

		units := MustParse(tokens[1])
		switch tokens[0] {
		case "forward":
			position += units
			depth += aim * units
			break

		case "up":
			aim -= units
			break

		case "down":
			aim += units
			break
		}
	})

	println(position * depth)
}

func MustParse(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
