package main

import (
	"fmt"
	"strings"
)

type passport map[string]string

func (p *passport) parsePassportFields(line string) {
	fields := strings.Split(line, " ")
	for _, f := range fields {
		x := strings.Split(f, ":")
		(*p)[x[0]] = x[1]
	}
}

func (p *passport) isValid() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, v := range requiredFields {
		if _, ok := (*p)[v]; !ok {
			return false
		}
	}
	return true
}

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
