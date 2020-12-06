package main

import "fmt"

type group map[rune]bool

func (g *group) yesCount() int {
	total := 0
	for range *g {
		total += 1
	}
	return total
}

func main() {
	g := group{}
	total := 0
	ReadInputFileByLine(func(line string) {
		if line == "" {
			total += g.yesCount()
			g = group{}
		}

		for _, c := range line {
			g[c] = true
		}
	})
	total += g.yesCount()

	fmt.Println(total)
}
