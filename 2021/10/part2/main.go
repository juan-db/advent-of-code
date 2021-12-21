package main

import (
	"github.com/juan-db/libaoc"
	"sort"
)

func IsOpeningBracket(c int32) bool {
	switch c {
	case '(':
		return true

	case '[':
		return true

	case '{':
		return true

	case '<':
		return true

	default:
		return false
	}
}

func GetBracketScore(b int32) int {
	switch b {
	case '(':
		return 1

	case '[':
		return 2

	case '{':
		return 3

	case '<':
		return 4
	}

	panic("illegal character")
}

func GetOpeningBracket(closingBracket int32) int32 {
	switch closingBracket {
	case ')':
		return '('

	case ']':
		return '['

	case '}':
		return '{'

	case '>':
		return '<'
	}

	panic("illegal closing bracket")
}

func main() {
	var completionScores []int
	libaoc.ReadInputFileByLine(func(line string) {
		var bracketStack []int32

		for _, v := range line {
			if IsOpeningBracket(v) {
				bracketStack = append(bracketStack, v)
			} else {
				if bracketStack[len(bracketStack)-1] != GetOpeningBracket(v) {
					return
				}

				bracketStack = bracketStack[:len(bracketStack)-1]
			}
		}

		var score int
		for i := len(bracketStack) - 1; i >= 0; i-- {
			score = score*5 + GetBracketScore(bracketStack[i])
		}
		completionScores = append(completionScores, score)
	})
	sort.Ints(completionScores)
	println(completionScores[len(completionScores)/2])
}
