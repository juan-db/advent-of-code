package main

import "fmt"

const (
	Floor byte = '.'
	Occupied byte = '#'
	Empty byte = 'L'
)

type Map struct {
	width int
	height int
	// Determines which array to retrieve, a when true, b otherwise.
	// The value will be flipped when swap is called.
	curr bool
	a [][]byte
	b [][]byte
}

func (m *Map) current() [][]byte {
	if m.curr {
		return m.a
	} else {
		return m.b
	}
}

func (m *Map) next() [][]byte {
	if m.curr {
		return m.b
	} else {
		return m.a
	}
}

func (m *Map) swap() {
	m.curr = !m.curr
}

func findSeat(x, y, xSlope, ySlope int, m *Map) byte {
	x += xSlope
	y += ySlope
	for ; x >= 0 && x < m.width && y >= 0 && y < m.height; x, y = x + xSlope, y + ySlope{
		tmp := m.current()[y][x]
		if tmp != Floor {
			return tmp
		}
	}
	return Floor
}

func countNeighbors(x, y int, m *Map) int {
	layout := m.current()

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

			if findSeat(x, y, i, j, m) == Occupied {
				neighbors++
			}
		}
	}
	return neighbors
}

func step(m *Map) (anyChange bool) {
	old, new := m.current(), m.next()
	for i, v := range old {
		copy(new[i], v)
	}

	anyChange = false
	for y, row := range old {
		for x, col := range row {
			neighbors := countNeighbors(x, y, m)
			// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
			// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
			// Otherwise, the seat's state does not change.
			var nextState byte
			if col == Empty && neighbors == 0 {
				nextState = Occupied
				anyChange = true
			} else if col == Occupied && neighbors >= 5 {
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
	m := Map{len(a[0]), len(a), true, a, b}

	const MAX_ITERS = 1000
	for i := 0; i < MAX_ITERS; i++ {
		changed := step(&m)
		m.swap()
		if !changed {
			fmt.Println(countOccupiedSeats(a))
			return
		}
	}
	panic("Layout kept changing")
}
