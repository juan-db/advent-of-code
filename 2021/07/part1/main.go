package main

import (
	"github.com/juan-db/libaoc"
	"math"
	"strconv"
	"strings"
)

func main() {
	var positions []int
	min := math.MaxInt
	max := math.MinInt

	libaoc.ReadInputFileByLine(func(line string) {
		tokens := strings.Split(line, ",")
		positions = make([]int, 0, len(tokens))

		for _, v := range tokens {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			positions = append(positions, num)

			if num > max {
				max = num
			}

			if num < min {
				min = num
			}
		}
	})

	var fuelSpend int
	minFuelSpend := math.MaxInt
	for i := min; i <= max; i++ {
		fuelSpend = 0

		for _, v := range positions {
			distance := v - i
			if distance < 0 {
				distance = -distance
			}

			fuelSpend += distance
		}

		if fuelSpend < minFuelSpend {
			minFuelSpend = fuelSpend
		}
	}

	println(minFuelSpend)
}
