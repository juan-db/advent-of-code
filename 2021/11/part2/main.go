package main

import (
	"github.com/juan-db/libaoc"
)

func flash(octopi [][]int, x, y int) int {
	flashes := 1
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if y+i < 0 || y+i >= len(octopi) || x+j < 0 || x+j >= len(octopi[0]) {
				continue
			}

			octopi[y+i][x+j] += 1

			if octopi[y+i][x+j] == 10 {
				flashes += flash(octopi, x+j, y+i)
			}
		}
	}
	return flashes
}

func simulate(octopi [][]int) int {
	flashes := 0
	for i := 0; i < len(octopi); i++ {
		for j := 0; j < len(octopi[i]); j++ {
			octopi[i][j] += 1
			if octopi[i][j] == 10 {
				flashes += flash(octopi, j, i)
			}
		}
	}
	return flashes
}

func resetFlashed(octopi [][]int) {
	for y := 0; y < len(octopi); y++ {
		for x := 0; x < len(octopi[0]); x++ {
			if octopi[y][x] > 9 {
				octopi[y][x] = 0
			}
		}
	}
}

func main() {
	var octopi [][]int
	libaoc.ReadInputFileByLine(func(line string) {
		var row []int
		for _, v := range line {
			row = append(row, int(v-'0'))
		}
		octopi = append(octopi, row)
	})

	for step := 1; ; step++ {
		var buffer [][]int
		for _, v := range octopi {
			row := make([]int, len(v))
			copy(row, v)
			buffer = append(buffer, row)
		}

		flashes := simulate(buffer)
		if flashes == 100 {
			println(step)
			break
		}

		resetFlashed(buffer)
		octopi = buffer
	}
}
