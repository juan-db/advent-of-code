package main

import "strconv"

func main() {
	initialized := false
	var last int
	increases := 0
	ReadInputFileByLine(func(line string) {
		current, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		if !initialized {
			initialized = true
			last = current
			return
		}

		if last < current {
			increases += 1
		}
		last = current
	})
	println(increases)
}
