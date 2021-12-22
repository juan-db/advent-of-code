package main

import (
	"github.com/juan-db/libaoc"
	"strings"
	"unicode"
)

type Cave struct {
	Name           string
	ConnectedCaves []*Cave
}

func (c *Cave) IsBig() bool {
	return unicode.IsUpper([]rune(c.Name)[0])
}

func (c *Cave) IsEnd() bool {
	return c.Name == "end"
}

func PathContains(path []*Cave, c *Cave) bool {
	for _, v := range path {
		if v == c {
			return true
		}
	}
	return false
}

func (c *Cave) Traverse(path []*Cave) int {
	if c.IsEnd() {
		return 1
	}

	if !c.IsBig() && PathContains(path, c) {
		return 0
	}

	paths := 0
	for _, v := range c.ConnectedCaves {
		paths += v.Traverse(append(path, c))
	}
	return paths
}

func main() {
	caves := map[string]*Cave{}
	libaoc.ReadInputFileByLine(func(line string) {
		ts := strings.Split(line, "-")

		fn := ts[0]
		var fc *Cave
		var ok bool
		if fc, ok = caves[fn]; !ok {
			fc = &Cave{Name: fn}
			caves[fn] = fc
		}

		tn := ts[1]
		var tc *Cave
		if tc, ok = caves[tn]; !ok {
			tc = &Cave{Name: tn}
			caves[tn] = tc
		}

		fc.ConnectedCaves = append(fc.ConnectedCaves, tc)
		tc.ConnectedCaves = append(tc.ConnectedCaves, fc)
	})

	println(caves["start"].Traverse(nil))
}
