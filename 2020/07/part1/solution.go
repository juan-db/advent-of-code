package main

import "fmt"

func (b *bag) climbTree(action func(*bag)) {
	for _, p := range b.parents {
		action(p.parent)
		p.parent.climbTree(action)
	}
}

func main() {
	bags := map[string]*bag{}
	ReadInputFileByLine(func(line string) {
		graphRule(line, &bags)
	})
	uniqueGoldParents := map[string]bool{}
	bags["shiny gold"].climbTree(func (b *bag) {
		uniqueGoldParents[b.color] = true
	})
	fmt.Println(len(uniqueGoldParents))
}
