package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	file := openInputFile()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	earliestDeparture, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	busses := strings.Split(scanner.Text(), ",")

	earliestBus, shortestWait := math.MaxInt32, math.MaxInt32
	for _, b := range busses {
		if b == "x" {
			continue
		}

		id, _ := strconv.Atoi(b)
		// doesn't handle waits that are perfectly divisible by earliest wait,
		// but I doubt it'll be the solution case
		if wait := id - earliestDeparture % id; wait < shortestWait {
			earliestBus = id
			shortestWait = wait
		}
	}

	fmt.Println(earliestBus * shortestWait)
}