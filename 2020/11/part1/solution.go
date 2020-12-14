package main

import "fmt"

const (
	Floor byte = '.'
	Occupied byte = '#'
	Empty byte = 'L'
)

func countNeighbors(x, y int, layout [][]byte) int {
	neighbors := 0
	row := layout[y]
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}

			a := x + i
			b := y + j

			// Make sure the index is in range
			if a < 0 || a >= len(row) || b < 0 || b >= len(layout) {
				continue
			}

			if layout[b][a] == Occupied {
				neighbors++
			}
		}
	}
	return neighbors
}

func step(old, new [][]byte) (anyChange bool) {
	for i, v := range old {
		copy(new[i], v)
	}

	anyChange = false
	for y, row := range old {
		for x, col := range row {
			neighbors := countNeighbors(x, y, old)
			// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
			// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
			// Otherwise, the seat's state does not change.
			var nextState byte
			if col == Empty && neighbors == 0 {
				nextState = Occupied
				anyChange = true
			} else if col == Occupied && neighbors >= 4 {
				nextState = Empty
				anyChange = true
			} else {
				nextState = col
			}
			new[y][x] = nextState
		}
	}
	return
}

func countOccupiedSeats(layout [][]byte) int {
	total := 0
	for _, row := range layout {
		for _, col := range row {
			if col == Occupied {
				total++
			}
		}
	}
	return total
}

func main() {
	var a, b [][]byte
	ReadInputFileByLine(func(line string) {
		row := make([]byte, len(line))
		for i, c := range line {
			row[i] = byte(c)
		}
		a = append(a, row)
	})
	b = make([][]byte, len(a))
	for i, row := range a {
		b[i] = make([]byte, len(row))
	}

	const MAX_ITERS = 1000
	for i := 0; i < MAX_ITERS; i++ {
		var changed bool
		if i % 2 == 0 {
			changed = step(a, b)
		} else {
			changed = step(b, a)
		}
		if !changed {
			fmt.Println(countOccupiedSeats(a))
			return
		}
	}
	panic("Layout kept changing")
}
