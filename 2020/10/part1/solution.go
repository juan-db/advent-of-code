package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	adapters := []int{0} // first value for the socket
	ReadInputFileByLine(func(line string) {
		num, _ := strconv.Atoi(line)
		adapters = append(adapters, num)
	})

	sort.Ints(adapters)

	ones := 0
	threes := 1 // one for the device itself
	for i, max := 0, len(adapters) - 1; i < max; i++ {
		switch diff := adapters[i + 1] - adapters[i]; diff {
		case 1:
			ones += 1
		case 3:
			threes += 1
		}
	}
	fmt.Println(ones * threes)
}
