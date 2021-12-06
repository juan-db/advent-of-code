package main

func main() {
	var bitCounts []int
	lineCount := 0

	ReadInputFileByLine(func(line string) {
		lineCount += 1

		if bitCounts == nil {
			bitCounts = make([]int, len(line))
		}

		for i, v := range line {
			if v == '1' {
				bitCounts[i] += 1
			}
		}
	})

	var gammaRate uint
	var mask uint
	for i, v := range bitCounts {
		gammaRate <<= 1

		if lineCount-v < v {
			gammaRate |= 1
		}

		mask |= 1 << i
	}

	epsilonRate := (^gammaRate) & mask

	println(gammaRate * epsilonRate)
}
