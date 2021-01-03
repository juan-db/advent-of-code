package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const MaskWidth = 36

type Mask string

func ApplyMask(num int64, mask Mask) int64 {
	result := int64(0)
	for i, c := range mask {
		result <<= 1

		if c == 'X' {
			result |= num >> (MaskWidth - 1 - i) & 1
		} else if c == '1' {
			result |= 1
		}
	}
	return result
}

func main() {
	memory := map[int]int64{}
	mask := Mask("")
	ReadInputFileByLine(func(line string) {
		tokens := strings.Split(line, "=")
		target := strings.TrimSpace(tokens[0])
		value := strings.TrimSpace(tokens[1])

		if target == "mask" {
			mask = Mask(value)
		} else {
			target = strings.TrimPrefix(target, "mem[")
			target = strings.TrimSuffix(target, "]")

			address, err := strconv.Atoi(target)
			if err != nil {
				log.Fatalf("could not parse memory address: %v", target)
			}
			num, err := strconv.ParseInt(value, 10, 36)
			if err != nil {
				log.Fatalf("could not parse memory value: %v", value)
			}
			memory[address] = ApplyMask(num, mask)
		}
	})

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}
