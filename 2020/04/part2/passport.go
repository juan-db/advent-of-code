package main

import (
	"strconv"
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
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !p.numValid("byr", 1920, 2002) {
		return false
	}

	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !p.numValid("iyr", 2010, 2020) {
		return false
	}

	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !p.numValid("eyr", 2020, 2030) {
		return false
	}

	//hgt (Height) - a number followed by either cm or in:
	//  If cm, the number must be at least 150 and at most 193.
	//  If in, the number must be at least 59 and at most 76.
	heightValidator := func(matches []string) bool {
		height, _ := strconv.Atoi(matches[1])
		if matches[2] == "cm" {
			return height >= 150 && height <= 193
		} else {
			return height >= 59 && height <= 76
		}
	}
	if !p.regexValid("hgt", "(\\d+)(cm|in)", heightValidator) {
		return false
	}

	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if !p.regexValid("hcl", "#[0-9a-f]{6}", nil) {
		return false
	}

	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if !p.regexValid("ecl", "^amb|blu|brn|gry|grn|hzl|oth$", nil) {
		return false
	}

	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	if !p.regexValid("pid", "^\\d{9}$", nil) {
		return false
	}

	//cid (Country ID) - ignored, missing or not.

	return true
}
