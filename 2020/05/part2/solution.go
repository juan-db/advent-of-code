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
	seats := map[int]bool{}
	ReadInputFileByLine(func(line string) {
		seat := seatFromBinaryPartition(line)
		seats[seat.row * (MaxCol + 1)  + seat.col] = true
	})
	for x, max := 1, MaxRow * (MaxCol + 1) + MaxCol - 1; x < max; x++ {
		if _, ok := seats[x]; ok {
			continue
		}
		if _, ok := seats[x - 1]; !ok {
			continue
		}
		if _, ok := seats[x + 1]; ok {
			fmt.Println(x)
			return
		}
	}
}
