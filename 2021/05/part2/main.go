package main

import (
	"strconv"
	"strings"
)

func main() {
	vents := make(map[Point]int)
	ReadInputFileByLine(func(line string) {
		l := ParseLine(line)

		if l.start.x == l.end.x {
			// vertical line
			var start int
			var end int
			if l.start.y < l.end.y {
				start = l.start.y
				end = l.end.y
			} else {
				start = l.end.y
				end = l.start.y
			}
			for y := start; y <= end; y++ {
				p := Point{l.start.x, y}

				if v, ok := vents[p]; ok {
					vents[p] = v + 1
				} else {
					vents[p] = 1
				}
			}
		} else if l.start.y == l.end.y {
			// horizontal line
			var start int
			var end int
			if l.start.x < l.end.x {
				start = l.start.x
				end = l.end.x
			} else {
				start = l.end.x
				end = l.start.x
			}
			for x := start; x <= end; x++ {
				p := Point{x, l.start.y}

				if v, ok := vents[p]; ok {
					vents[p] = v + 1
				} else {
					vents[p] = 1
				}
			}
		} else {
			var start Point
			var end Point
			if l.start.x < l.end.x {
				start = l.start
				end = l.end
			} else {
				start = l.end
				end = l.start
			}

			var step int
			if start.y < end.y {
				step = 1
			} else {
				step = -1
			}

			for x, y := start.x, start.y; x <= end.x; x, y = x+1, y+step {
				p := Point{x, y}
				if v, ok := vents[p]; ok {
					vents[p] = v + 1
				} else {
					vents[p] = 1
				}
			}
		}
	})

	overlaps := 0
	for _, v := range vents {
		if v > 1 {
			overlaps += 1
		}
	}
	println(overlaps)
}

func ParseLine(text string) Line {
	tokens := strings.Split(text, " ")
	return Line{
		start: ParsePoint(tokens[0]),
		end:   ParsePoint(tokens[2]),
	}
}

func ParsePoint(text string) Point {
	tokens := strings.Split(text, ",")

	x, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	return Point{x, y}
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}
