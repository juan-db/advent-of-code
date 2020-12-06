package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Policy struct {
	min int
	max int
	letter string
}

func parsePolicy(text string) Policy {
	constraints := strings.Split(text, " ")
	counts := strings.Split(constraints[0], "-")

	min, _ := strconv.Atoi(counts[0])
	max, _ := strconv.Atoi(counts[1])

	return Policy{
		min,
		max,
		constraints[1],
	}
}

func isValid(pass string, policy Policy) bool {
	letterCount := strings.Count(pass, policy.letter)
	return letterCount >= policy.min && letterCount <= policy.max
}

func main() {
	inputFile := OpenInputFile()

	valid := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, ":")
		policy := parsePolicy(strs[0])
		pass := strs[1]

		if isValid(pass, policy) {
			valid += 1
		}
	}

	fmt.Println(valid)
}