package main

import (
	"regexp"
	"strconv"
)

func (p *passport) numValid(key string, min, max int) bool {
	if v, ok := (*p)[key]; ok {
		num, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if num < min || num > max {
			return false
		}

		return true
	} else {
		return false
	}
}

// If matchValidator is nil, any match will be considered valid.
func (p *passport) regexValid(key string, regexString string, matchValidator func([]string)bool) bool {
	regex, err := regexp.Compile(regexString)
	if err != nil {
		panic(err)
	}

	if v, ok := (*p)[key]; ok {
		if matches := regex.FindStringSubmatch(v); matches != nil {
			if matchValidator == nil {
				return true
			} else {
				return matchValidator(matches)
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

