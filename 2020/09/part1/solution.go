package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func isSumPresent(num int, nums *[]int) bool {
	for i, x := range *nums {
		diff := num - x
		for j, y := range *nums {
			if i == j {
				continue
			}

			if y == diff {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := make([]int, 25)
	i := 0

	file := openInputFile()
	scanner := bufio.NewScanner(file)
	for ; i < 25; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		nums[i] = num
	}

	i = 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if !isSumPresent(num, &nums) {
			fmt.Println(num)
			return
		}

		nums[i] = num
		if i++; i > 24 {
			i = 0
		}
	}
}
