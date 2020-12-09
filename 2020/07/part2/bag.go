package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color    string
	parents  []*link
	children []*link
}

type link struct {
	parent *bag
	child  *bag
	count  int
}

var ruleRegex *regexp.Regexp = regexp.MustCompile(`(.*) bags contain (.*)\.`)
var bagRegex = regexp.MustCompile(`(no other bags)|((\d+) (.+) bags?)`)

func parseChild(str string) *struct {
	color string
	count int
} {
	m := bagRegex.FindStringSubmatch(str)
	if m == nil {
		panic(fmt.Errorf("invalid child bag string: '%v'", str))
	}

	if m[1] != "" {
		return nil
	}

	countStr := m[3]
	count, err := strconv.Atoi(countStr)
	if err != nil {
		panic(fmt.Errorf("failed to parse child bag count: %v: %v", countStr, err))
	}

	color := m[4]

	return &struct {
		color string
		count int
	}{
		color: color,
		count: count,
	}
}

func getBag(color string, bags *map[string]*bag) *bag {
	if b, ok := (*bags)[color]; ok {
		return b
	} else {
		b = &bag{
			color,
			[]*link{},
			[]*link{},
		}
		(*bags)[color] = b
		return b
	}
}

func (b *bag) link(child *bag, count int) {
	l := &link{b, child, count}
	b.children = append(b.children, l)
	child.parents = append(child.parents, l)
}

func graphRule(rule string, bags *map[string]*bag) {
	matches := ruleRegex.FindStringSubmatch(rule)
	if matches == nil {
		panic(fmt.Errorf("invalid rule: '%v'", rule))
	}

	parent := getBag(matches[1], bags)

	for _, i := range strings.Split(matches[2], ",") {
		c := parseChild(i)
		if c == nil {
			return
		}

		child := getBag(c.color, bags)
		parent.link(child, c.count)
	}
}

