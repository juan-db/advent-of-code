package main

import (
	"strconv"
	"strings"
)

func main() {
	var bingoNumbers []int
	boardLines := make([]string, 0, 5)
	var bingoBoards []*BingoBoard
	ReadInputFileByLine(func(line string) {
		if bingoNumbers == nil {
			bingoNumbers = parseBingoNumbers(line)
			return
		}

		if line == "" {
			return
		}

		boardLines = append(boardLines, line)
		if len(boardLines) < 5 {
			return
		}

		bingoBoards = append(bingoBoards, NewBingoBoard(boardLines))
		boardLines = boardLines[:0]
	})

	for _, num := range bingoNumbers {
		for _, board := range bingoBoards {
			board.MarkNumber(num)
			if board.HasWon() {
				boardScore := board.GetScore()
				println(boardScore * num)
				return
			}
		}
	}
}

func parseBingoNumbers(line string) []int {
	tokens := strings.Split(line, ",")
	numbers := make([]int, 0, len(tokens))
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
	}
	return numbers
}

type BingoBoard struct {
	board  [5][5]int
	marked [5][5]bool
}

func NewBingoBoard(lines []string) *BingoBoard {
	var board [5][5]int
	for i, v := range lines {
		var line [5]int
		index := 0
		splitLine := strings.Split(v, " ")
		for _, t := range splitLine {
			if t == "" {
				continue
			}

			num, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}

			line[index] = num
			index++
		}
		board[i] = line
	}

	output := BingoBoard{}
	output.board = board
	return &output
}

func (b *BingoBoard) MarkNumber(number int) {
	for i, v := range b.board {
		for j, w := range v {
			if w == number {
				b.marked[i][j] = true
				return
			}
		}
	}
}

func (b *BingoBoard) HasWon() bool {
	checkRow := func(row [5]bool) bool {
		for _, v := range row {
			if !v {
				return false
			}
		}
		return true
	}

	checkColumn := func(column int) bool {
		for _, v := range b.marked {
			if !v[column] {
				return false
			}
		}
		return true
	}

	for _, v := range b.marked {
		if checkRow(v) {
			return true
		}
	}

	for i := 0; i < len(b.marked); i++ {
		if checkColumn(i) {
			return true
		}
	}

	return false
}

func (b *BingoBoard) GetScore() int {
	score := 0
	for i, v := range b.marked {
		for j, w := range v {
			if !w {
				score += b.board[i][j]
			}
		}
	}
	return score
}
