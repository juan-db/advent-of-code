package main

import "fmt"

func (b *bag) countChildrenRecursively() int {
	total := 1
	for _, c := range b.children {
		total += c.count * c.child.countChildrenRecursively()
	}
	return total
}

func main() {
	bags := map[string]*bag{}
	ReadInputFileByLine(func(line string) {
		graphRule(line, &bags)
	})
	total := bags["shiny gold"].countChildrenRecursively()
	// - 1 because we want to exclude the gold bag itself.
	fmt.Println(total - 1)
}
