package main

import (
	"github.com/juan-db/libaoc"
	"sort"
	"strconv"
	"strings"
)

const NoneTag = 0

type Heightmap struct {
	Map      [][]int
	BasinMap [][]int

	Width  int
	Height int

	lastTag int
}

func (h *Heightmap) GetNextTag() int {
	h.lastTag++
	return h.lastTag
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
	h.BasinMap = append(h.BasinMap, make([]int, h.Width))

	h.Height = len(h.Map)
}

func (h *Heightmap) TryGetPoint(x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= h.Width || y >= h.Height {
		return 0, false
	}

	return h.Map[y][x], true
}

func (h *Heightmap) TryGetTag(x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= h.Width || y >= h.Height {
		return 0, false
	}

	return h.BasinMap[y][x], true
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

func (h *Heightmap) TagBasin(tag, x, y int) int {
	// for each point:
	// - go up until we hit a 9 or a tag, then
	// - go left until we hit a 9 or a tag, then
	// - right until 9 or tag, then
	// - down until 9 or tag

	if v, ok := h.TryGetTag(x, y); !ok || v != NoneTag {
		// We've hit a map edge or a tag
		return 0
	}

	// We don't have to check ok again because we already did map edge detection
	// above when checking for a tag
	if v, _ := h.TryGetPoint(x, y); v == 9 {
		// We've hit the edge of the basin
		return 0
	}

	h.BasinMap[y][x] = tag

	basinSize := 1
	neighbors := []struct {
		x int
		y int
	}{
		{x: x - 1, y: y},
		{x: x + 1, y: y},
		{x: x, y: y - 1},
		{x: x, y: y + 1},
	}
	for _, v := range neighbors {
		basinSize += h.TagBasin(tag, v.x, v.y)
	}
	return basinSize
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

	var basinSizes []int
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.IsLowestPoint(x, y) {
				basinSizes = append(basinSizes, m.TagBasin(m.GetNextTag(), x, y))
			}
		}
	}

	sort.Ints(basinSizes)

	answer := 1
	for _, v := range basinSizes[len(basinSizes)-3:] {
		answer *= v
	}
	println(answer)
}
