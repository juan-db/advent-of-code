package main

import (
	"fmt"
	"math"
	"strconv"
)

const invalidNumber = 257342611

func findSet(nums *[]int, maxLen int) (found bool, min, max int) {
	for i := 0; i < len(*nums) - maxLen; i++ {
		min, max = math.MaxInt32, math.MinInt32
		total := 0
		for j := i; j - i < maxLen; j++ {
			num := (*nums)[j]

			if num < min {
				min = num
			}
			if num > max {
				max = num
			}

			total += (*nums)[j]
		}
		if total == invalidNumber {
			return true, min, max
		}
	}

	return false, 0, 0
}

func main() {
	var nums []int
	ReadInputFileByLine(func (line string) {
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	})

	for i := 2; i < 100; i++ {
		if found, min, max := findSet(&nums, i); found {
			fmt.Println(min + max)
			return
		}
	}

	panic("no set found")
}
