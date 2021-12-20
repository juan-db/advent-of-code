package main

import (
	"github.com/juan-db/libaoc"
	"strconv"
	"strings"
)

type Heightmap struct {
	Map    [][]int
	Width  int
	Height int
}

func (h *Heightmap) AppendRow(row []int) {
	if len(h.Map) > 0 {
		// We already have rows, so we should already know the width of the map
		if len(row) != h.Width {
			panic("row width does not match previous row widths")
		}
	} else {
		h.Width = len(row)
	}

	h.Map = append(h.Map, row)

	h.Height = len(h.Map)
}

func (h *Heightmap) TryGetPoint(x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= h.Width || y >= h.Height {
		return 0, false
	}

	return h.Map[y][x], true
}

func (h *Heightmap) IsLowestPoint(x, y int) bool {
	neighbors := []struct {
		x int
		y int
	}{
		{x: x - 1, y: y},
		{x: x + 1, y: y},
		{x: x, y: y - 1},
		{x: x, y: y + 1},
	}

	lowest := true
	point := h.Map[y][x]
	for _, v := range neighbors {
		if w, ok := h.TryGetPoint(v.x, v.y); ok {
			if w <= point {
				lowest = false
			}
		}
	}

	return lowest
}

func main() {
	m := Heightmap{}
	libaoc.ReadInputFileByLine(func(line string) {
		row := strings.Split(line, "")
		rowHeights := make([]int, 0, len(row))
		for _, v := range row {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			rowHeights = append(rowHeights, num)
		}
		m.AppendRow(rowHeights)
	})

	var riskSum int
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.IsLowestPoint(x, y) {
				riskSum += 1 + m.Map[y][x]
			}
		}
	}
	println(riskSum)
}
