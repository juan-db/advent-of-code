package main

import (
	"fmt"
	"strconv"
)

type direction struct {
	character byte
	x int
	y int
}

var directions = []direction{
	{'E', 1, 0},
	{'S', 0, -1},
	{'W', -1, 0},
	{'N', 0, 1},
}

type ship struct {
	x int
	y int
	d int
}

func (s *ship) direction() direction {
	return directions[s.d]
}

func (s *ship) execute(instruction string) {
	d := instruction[0]
	v, _ := strconv.Atoi(instruction[1:])

	switch d {
	case 'N':
		s.y += v

	case 'E':
		s.x += v

	case 'S':
		s.y -= v

	case 'W':
		s.x -= v

	case 'F':
		s.x += s.direction().x * v
		s.y += s.direction().y * v

	case 'L':
		i := (s.d - v / 90) % len(directions)
		if i < 0 {
			i = len(directions) + i
		}
		s.d = i

	case 'R':
		s.d = (s.d + v / 90) % len(directions)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	s := ship{0, 0, 0}
	ReadInputFileByLine(func(line string) {
		s.execute(line)
	})
	fmt.Println(abs(s.x) + abs(s.y))
}
