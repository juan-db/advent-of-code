package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Policy struct {
	posA   int
	posB   int
	letter byte
}

func parsePolicy(text string) Policy {
	constraints := strings.Split(text, " ")
	positions := strings.Split(constraints[0], "-")

	posA, _ := strconv.Atoi(positions[0])
	posB, _ := strconv.Atoi(positions[1])

	return Policy{
		posA,
		posB,
		constraints[1][0],
	}
}

func isValid(pass string, policy Policy) bool {
	letterA := pass[policy.posA]
	letterB := pass[policy.posB]

	return letterA != letterB && (letterA == policy.letter || letterB == policy.letter)
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