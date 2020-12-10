package main

import (
	"fmt"
	"sort"
	"strconv"
)

var distances = map[int]int{}
func findDistanceToDevice(as []int) int {
	if len(as) == 1 {
		// only the socket left
		return 1
	}

	i := len(as) - 1
	if v, ok := distances[as[i]]; ok {
		return v
	}

	total := 0
	for i, j := len(as)-1, len(as)-2; j >= 0; j-- {
		if as[i]-as[j] <= 3 {
			distance := findDistanceToDevice(as[:j + 1])
			distances[as[j]] = distance
			total += distance
		}
	}
	return total
}

func main() {
	adapters := []int{0} // first value for the socket
	ReadInputFileByLine(func(line string) {
		num, _ := strconv.Atoi(line)
		adapters = append(adapters, num)
	})

	sort.Ints(adapters)

	fmt.Println(findDistanceToDevice(adapters))
}
