package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoChanges(t *testing.T) {
	mask := Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	num := int64(0b101010101)

	expected := int64(0b101010101)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestZeroToOne(t *testing.T) {
	mask := Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XX")
	num := int64(0)

	expected := int64(0b100)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestOneToZero(t *testing.T) {
	mask := Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX0XX")
	num := int64(0b100)

	expected := int64(0)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestManyZeroesToOnes(t *testing.T) {
	mask := Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1X1X1")
	num := int64(0)

	expected := int64(0b10101)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestManyOnesToZero(t *testing.T) {
	mask := Mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX0X0X0")
	num := int64(0b10101)

	expected := int64(0)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestAllZeroesToOnes(t *testing.T) {
	mask := Mask("111111111111111111111111111111111111")
	num := int64(0)

	expected := int64(0b111111111111111111111111111111111111)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestAllOnesToZeroes(t *testing.T) {
	mask := Mask("000000000000000000000000000000000000")
	num := int64(0b111111111111111111111111111111111111)

	expected := int64(0)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestMSBToZero(t *testing.T) {
	mask := Mask("0XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	num := int64(0b100000000000000000000000000000000000)

	expected := int64(0)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}

func TestMSBToOne(t *testing.T) {
	mask := Mask("1XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	num := int64(0)

	expected := int64(0b100000000000000000000000000000000000)
	actual := ApplyMask(num, mask)

	assert.Equal(t, expected, actual)
}
