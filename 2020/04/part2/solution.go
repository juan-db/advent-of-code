package main

import (
	"fmt"
)


func main() {
	file := OpenInputFile()
	p := &passport{}
	valid := 0
	ReadByLine(file, func(line string) {
		if line == "" {
			if p.isValid() {
				valid += 1
			}
			p = &passport{}
		} else {
			p.parsePassportFields(line)
		}
	})
	if p.isValid() {
		valid += 1
	}

	fmt.Println(valid)
}
