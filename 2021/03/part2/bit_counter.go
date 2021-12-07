package main

type BitCounter interface {
	ProcessLine(line string)
	GetMostCommonMask() string
}

type bitCounterImpl struct {
	lineCount int
	bitCounts []int
}

func newBitCounter() BitCounter {
	return &bitCounterImpl{0, nil}
}

func (c *bitCounterImpl) ProcessLine(line string) {
	c.lineCount += 1

	if c.bitCounts == nil {
		c.bitCounts = make([]int, len(line))
	}

	for i, v := range line {
		if v == '1' {
			c.bitCounts[i] += 1
		}
	}
}

func (c *bitCounterImpl) GetMostCommonMask() string {
	var mask []rune
	for _, v := range c.bitCounts {
		ones := v
		zeroes := c.lineCount - v
		var char rune
		if ones >= zeroes {
			char = '1'
		} else {
			char = '0'
		}
		mask = append(mask, char)
	}
	return string(mask)
}
