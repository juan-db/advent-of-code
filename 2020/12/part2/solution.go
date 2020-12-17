package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

// Rotates the given point clockwise, angle degrees.
// Only works for 90 degree intervals since the input doesn't seem to require more complex rotations.
func rotate(waypoint *point, angle int) {
	// couldn't get the fancy matrix thing to work
	switch angle % 360 {
	case 0:
		break

	case 90:
		waypoint.x, waypoint.y = waypoint.y, -waypoint.x

	case 180:
		waypoint.x, waypoint.y = -waypoint.x, -waypoint.y

	case 270:
		waypoint.x, waypoint.y = -waypoint.y, waypoint.x

	default:
		panic(fmt.Errorf("can't rotate by %v degrees", angle))
	}
}

func execute(ship, waypoint *point, instruction string) {
	d := instruction[0]
	v, _ := strconv.Atoi(instruction[1:])

	switch d {
	case 'N':
		waypoint.y += v

	case 'E':
		waypoint.x += v

	case 'S':
		waypoint.y -= v

	case 'W':
		waypoint.x -= v

	case 'L':
		rotate(waypoint, 360 - v)

	case 'R':
		rotate(waypoint, v)

	case 'F':
		ship.x += waypoint.x * v
		ship.y += waypoint.y * v
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func main() {
	ship, waypoint := point{0, 0}, point{10, 1}
	ReadInputFileByLine(func(line string) {
		execute(&ship, &waypoint, line)
	})
	fmt.Println(abs(ship.x) + abs(ship.y))
}
