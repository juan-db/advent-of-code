package main

import "fmt"

type group struct {
	yeses map[rune]int
	people int
}

func newGroup() *group {
	return &group{
		map[rune]int{},
		0,
	}
}

func (g *group) yesCount() int {
	total := 0
	for _, v := range g.yeses {
		if v == g.people {
			total += 1
		}
	}
	return total
}

func main() {
	g := newGroup()
	total := 0
	ReadInputFileByLine(func(line string) {
		if line == "" {
			total += g.yesCount()
			g = newGroup()
			return
		}

		g.people += 1
		for _, c := range line {
			if v, ok := g.yeses[c]; ok {
				g.yeses[c] = v + 1
			} else {
				g.yeses[c] = 1
			}
		}
	})
	total += g.yesCount()

	fmt.Println(total)
}
