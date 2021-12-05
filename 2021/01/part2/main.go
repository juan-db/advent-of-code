package main

import (
	"strconv"
)

func main() {
	queue := make(chan int, 3)
	sum := 0
	increases := 0

	ReadInputFileByLine(func(line string) {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if len(queue) < 3 {
			queue <- num
			sum += num
			return
		}

		last := <-queue
		newSum := sum - last + num
		if sum < newSum {
			increases += 1
		}

		sum = newSum
		queue <- num
	})

	println(increases)
}
