package main

import (
	"bufio"
	"fmt"
)

type Slope struct {
	x int
	y int
	trees int
}

func main() {
	file := OpenInputFile()

	scanner := bufio.NewScanner(file)
	slopes := []Slope{
		{1, 1, 0},
		{3, 1, 0},
		{5, 1, 0},
		{7, 1, 0},
		{1, 2, 0},
	}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		width := len(line)
		for i, s := range slopes {
			if y % s.y == 0 {
				x := s.x * (y / s.y) % width
				if line[x] == '#' {
					slopes[i].trees += 1
				}
			}
		}
		y += 1
	}

	total := 1
	for _, s := range slopes {
		total *= s.trees
	}

	fmt.Println(total)
}
