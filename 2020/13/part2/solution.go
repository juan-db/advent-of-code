package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type bus struct {
	id    int64
	index int64
}

func parseBusses(line string) []bus {
	tokens := strings.Split(line, ",")
	busses := make([]bus, 0, len(tokens))
	for i, b := range tokens {
		if b == "x" {
			continue
		}

		id, _ := strconv.Atoi(b)
		busses = append(busses, bus{int64(id), int64(i)})
	}
	return busses
}

/*
                                  1                             2
    0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5  6  7  8  9  0  1  2  3
	*--*--*--\--*--*--*--*--\--*--*--*--*--\--*--*--*--*--\--*--*--*--*--*-
	*----*----\----*----*----\----*----*----\----*----*----\----*----*----*
	*-*-*-*-*-*-*-*-*-*-*-*-*-\-*-*-*-*-*-*-*-*-*-*-*-*-*-*-\-*-*-*-*-*-*-*
 */

func findNext(start, current bus, startFactor, step int64) int64 {
	for i := startFactor; ; i += step {
		start := start.id * i
		if (start + current.index) % current.id == 0 {
			return i
		}
	}
}

func getBusInterval(start, current bus, startAt, step int64) (startsAt, interval int64) {
	first := findNext(start, current, startAt, step)
	second := findNext(start, current, first + step, step)
	return first, second - first
}

func main() {
	file := openInputFile()
	scanner := bufio.NewScanner(file)

	scanner.Scan() // discard first line

	scanner.Scan()
	busses := parseBusses(scanner.Text())

	interval := int64(1)
	startAt := int64(1)
	for _, b := range busses[1:] {
		startAt, interval = getBusInterval(busses[0], b, startAt, interval)
		fmt.Printf("%v %v\n", startAt, interval)
	}

	for _, b := range busses {
		x := b.id * ((busses[0].id * startAt + b.index) / b.id)
		fmt.Printf("%+v: %v\n", b, x)
	}
}
