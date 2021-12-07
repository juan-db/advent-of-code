package main

import (
	"strconv"
)

func main() {
	var lines []string
	bitCounter := newBitCounter()
	ReadInputFileByLine(func(line string) {
		lines = append(lines, line)
		bitCounter.ProcessLine(line)
	})

	var potentialO2Indices []int
	var potentialCo2Indices []int

	o2Mask := bitCounter.GetMostCommonMask()
	co2Mask := bitCounter.GetMostCommonMask()
	for i, v := range lines {
		if v[0] == o2Mask[0] {
			potentialO2Indices = append(potentialO2Indices, i)
		}
		if v[0] != co2Mask[0] {
			potentialCo2Indices = append(potentialCo2Indices, i)
		}
	}

	binO2Rating := FilterRatings(lines, potentialO2Indices, 1, false)
	o2Rating, err := strconv.ParseUint(binO2Rating, 2, len(binO2Rating))
	if err != nil {
		panic(err)
	}

	binCo2Rating := FilterRatings(lines, potentialCo2Indices, 1, true)
	co2Rating, err := strconv.ParseUint(binCo2Rating, 2, len(binCo2Rating))
	if err != nil {
		panic(err)
	}

	println(o2Rating * co2Rating)
}

func FilterRatings(lines []string, potentialIndices []int, bitIndex int, co2 bool) string {
	bitCounter := newBitCounter()
	for _, v := range potentialIndices {
		bitCounter.ProcessLine(lines[v])
	}

	var newPotentialIndices []int
	mask := bitCounter.GetMostCommonMask()
	for _, v := range potentialIndices {
		var match bool
		if co2 {
			match = lines[v][bitIndex] != mask[bitIndex]
		} else {
			match = lines[v][bitIndex] == mask[bitIndex]
		}
		if match {
			newPotentialIndices = append(newPotentialIndices, v)
		}
	}

	if len(newPotentialIndices) == 1 {
		return lines[newPotentialIndices[0]]
	} else {
		return FilterRatings(lines, newPotentialIndices, bitIndex+1, co2)
	}
}
