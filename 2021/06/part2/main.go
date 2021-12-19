package main

import "github.com/juan-db/libaoc"

func main() {
	fish := map[int]int{}
	for i := 0; i < 9; i++ {
		fish[i] = 0
	}

	libaoc.ReadInputFileByLine(func(line string) {
		for _, v := range line {
			if v == ',' {
				continue
			}

			num := int(v - '0')
			fish[num] += 1
		}
	})

	for i := 0; i < 256; i++ {
		newFish := map[int]int{}
		for j := 0; j < 8; j++ {
			newFish[j] = fish[j+1]
		}
		newFish[6] += fish[0]
		newFish[8] = fish[0]

		fish = newFish
	}

	count := 0
	for _, v := range fish {
		count += v
	}
	println(count)
}
