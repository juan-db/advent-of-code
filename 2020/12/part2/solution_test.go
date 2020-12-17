package main

import (
	"testing"
)

func TestRotate0Degrees(t *testing.T) {
	actual := point{1, 0}

	rotate(&actual, 0)

	expected := point{1, 0}
	if actual.x != expected.x || actual.y != expected.y {
		t.Fatalf("expected: %v\nbut got: %v\n", expected, actual)
	}
}

func TestRotate90Degrees(t *testing.T) {
	actual := point{1, 0}

	rotate(&actual, 90)

	expected := point{0, -1}
	if actual.x != expected.x || actual.y != expected.y {
		t.Fatalf("expected: %v\nbut got: %v\n", expected, actual)
	}
}

func TestRotate180Degrees(t *testing.T) {
	actual := point{1, 0}

	rotate(&actual, 180)

	expected := point{-1, 0}
	if actual.x != expected.x || actual.y != expected.y {
		t.Fatalf("expected: %v\nbut got: %v\n", expected, actual)
	}
}

func TestRotate270Degrees(t *testing.T) {
	actual := point{1, 0}

	rotate(&actual, 270)

	expected := point{0, 1}
	if actual.x != expected.x || actual.y != expected.y {
		t.Fatalf("expected: %v\nbut got: %v\n", expected, actual)
	}
}

func TestRotate360Degrees(t *testing.T) {
	actual := point{1, 0}

	rotate(&actual, 360)

	expected := point{1, 0}
	if actual.x != expected.x || actual.y != expected.y {
		t.Fatalf("expected: %v\nbut got: %v\n", expected, actual)
	}
}
