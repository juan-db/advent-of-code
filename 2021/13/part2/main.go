package main

import (
	"github.com/juan-db/libaoc"
	"strconv"
	"strings"
)

func MustParse(v string) int {
	n, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return n
}

type Point struct {
	x int
	y int
}

func CreatePaper(points []Point) [][]bool {
	var largestX, largestY int
	for _, v := range points {
		if v.x > largestX {
			largestX = v.x
		}
		if v.y > largestY {
			largestY = v.y
		}
	}

	var paper [][]bool
	for i := 0; i <= largestY; i++ {
		paper = append(paper, make([]bool, largestX+1))
	}

	for _, v := range points {
		paper[v.y][v.x] = true
	}

	return paper
}

func Fold(paper [][]bool, x, y int) [][]bool {
	if x >= 0 {
		// Fold along x
		d := len(paper[0]) - x
		for i := 0; i < d; i++ {
			for j := 0; j < len(paper); j++ {
				paper[j][x-i] = paper[j][x+i] || paper[j][x-i]
			}
		}
		for i := range paper {
			paper[i] = paper[i][:x]
		}
		return paper
	} else {
		// Fold along y
		d := len(paper) - y
		for i := 0; i < d; i++ {
			for j := 0; j < len(paper[0]); j++ {
				paper[y-i][j] = paper[y+i][j] || paper[y-i][j]
			}
		}
		return paper[:y]
	}
}

func PrintPaper(paper [][]bool) {
	for _, v := range paper {
		for _, w := range v {
			if w {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func main() {
	var points []Point
	var paper [][]bool
	var pointsFinished bool
	libaoc.ReadInputFileByLine(func(line string) {
		if line == "" {
			pointsFinished = true
			paper = CreatePaper(points)
			return
		}

		if !pointsFinished {
			// Read new dots
			tokens := strings.Split(line, ",")
			x := MustParse(tokens[0])
			y := MustParse(tokens[1])
			points = append(points, Point{x, y})
		} else {
			// Read and perform folds
			fold := strings.Split(line, " ")[2]
			tokens := strings.Split(fold, "=")
			axis := tokens[0]
			coord := MustParse(tokens[1])
			if axis == "x" {
				paper = Fold(paper, coord, -1)
			} else {
				paper = Fold(paper, -1, coord)
			}
		}
	})
	PrintPaper(paper)
}
