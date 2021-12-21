package main

import (
	"github.com/juan-db/libaoc"
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
	case ')':
		return 3

	case ']':
		return 57

	case '}':
		return 1197

	case '>':
		return 25137
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
	errorScore := 0
	libaoc.ReadInputFileByLine(func(line string) {
		var bracketStack []int32
		for _, v := range line {
			if IsOpeningBracket(v) {
				bracketStack = append(bracketStack, v)
			} else {
				if bracketStack[len(bracketStack)-1] != GetOpeningBracket(v) {
					errorScore += GetBracketScore(v)
					return
				}

				bracketStack = bracketStack[:len(bracketStack)-1]
			}
		}
	})
	println(errorScore)
}
