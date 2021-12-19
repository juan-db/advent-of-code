package main

import (
	"github.com/juan-db/libaoc"
	"strings"
)

func main() {
	var uniqueDigitOutputs int

	libaoc.ReadInputFileByLine(func(line string) {
		outputs := strings.Split(line, "|")[1]
		tokens := strings.Split(outputs, " ")
		for _, v := range tokens {
			switch len(v) {
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 7:
				uniqueDigitOutputs++
			}
		}
	})

	println(uniqueDigitOutputs)
}
