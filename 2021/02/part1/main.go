package main

import (
	"strconv"
	"strings"
)

func main() {
	position := 0
	depth := 0

	ReadInputFileByLine(func(line string) {
		tokens := strings.Split(line, " ")

		units := MustParse(tokens[1])
		switch tokens[0] {
		case "forward":
			position += units
			break

		case "up":
			depth -= units
			break

		case "down":
			depth += units
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
