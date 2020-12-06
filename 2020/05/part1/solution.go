package main

import "fmt"

const MaxRow = 127
const MaxCol = 7

type seat struct {
	row int
	col int
}

func binaryPartition(commands string, downCommand, upCommand rune, min, max int) int {
	for _, x := range commands {
		move := (max - min) / 2 + 1
		if x == downCommand {
			max -= move
		} else if x == upCommand { // 'B'
			min +=  move
		} else {
			panic(fmt.Sprintf("unrecognized partition command: %v", x))
		}
	}

	if min != max {
		panic(fmt.Sprintf("min and max don't match after partitioning: %v %v", min, max))
	}

	return min
}

func seatFromBinaryPartition(bpStr string) seat {
	row := binaryPartition(bpStr[0:7], 'F', 'B', 0, MaxRow)
	col := binaryPartition(bpStr[7:], 'L', 'R', 0, MaxCol)
	return seat{row, col}
}

func main() {
	max := 0
	ReadInputFileByLine(func(line string) {
		seat := seatFromBinaryPartition(line)
		cur := seat.row * (MaxCol + 1)  + seat.col
		if cur > max {
			max = cur
		}
	})
	fmt.Println(max)
}
