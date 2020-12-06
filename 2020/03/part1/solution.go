package main

import (
	"bufio"
	"fmt"
)

func main() {
	file := OpenInputFile()

	scanner := bufio.NewScanner(file)
	x := 0
	trees := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line[x % len(line)] == '#' {
			trees += 1
		}
		x += 3
	}

	fmt.Println(trees)
}
