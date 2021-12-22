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

type Path struct {
	Path                []*Cave
	HasDoubleSmallVisit bool
}

func (c *Cave) IsBig() bool {
	return unicode.IsUpper([]rune(c.Name)[0])
}

func (c *Cave) IsStart() bool {
	return c.Name == "start"
}

func (c *Cave) IsEnd() bool {
	return c.Name == "end"
}

func (p *Path) Contains(c *Cave) bool {
	for _, v := range p.Path {
		if v == c {
			return true
		}
	}
	return false
}

func (c *Cave) Traverse(path Path) int {
	if c.IsEnd() {
		return 1
	}

	if !c.IsBig() && path.Contains(c) {
		if !path.HasDoubleSmallVisit && !c.IsStart() {
			path.HasDoubleSmallVisit = true
		} else {
			return 0
		}
	}

	paths := 0
	for _, v := range c.ConnectedCaves {
		paths += v.Traverse(Path{
			append(path.Path, c),
			path.HasDoubleSmallVisit,
		})
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

	println(caves["start"].Traverse(Path{}))
}
