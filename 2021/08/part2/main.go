package main

import (
	"github.com/juan-db/libaoc"
	"strings"
)

type Segment int

const (
	Top Segment = iota
	TopLeft
	TopRight
	Middle
	BottomLeft
	BottomRight
	Bottom
)

var AllSegments = []Segment{
	Top,
	TopLeft,
	TopRight,
	Middle,
	BottomLeft,
	BottomRight,
	Bottom,
}

var Digits = map[int][]Segment{
	0: []Segment{Top, TopLeft, TopRight, BottomLeft, BottomRight, Bottom},
	1: []Segment{TopRight, BottomRight},
	2: []Segment{Top, TopRight, Middle, BottomLeft, Bottom},
	3: []Segment{Top, TopRight, Middle, BottomRight, Bottom},
	4: []Segment{TopLeft, TopRight, Middle, BottomRight},
	5: []Segment{Top, TopLeft, Middle, BottomRight, Bottom},
	6: []Segment{Top, TopLeft, Middle, BottomLeft, BottomRight, Bottom},
	7: []Segment{Top, TopRight, BottomRight},
	8: []Segment{Top, TopLeft, TopRight, Middle, BottomLeft, BottomRight, Bottom},
	9: []Segment{Top, TopLeft, TopRight, Middle, BottomRight, Bottom},
}

type DisplayDecoder struct {
	PotentialSignals  map[Segment]map[string]bool
	ImpossibleSignals map[Segment]map[string]bool
	RecordedDigits    map[int]bool
}

func CreateDisplayDecoder() *DisplayDecoder {
	potentialSignals := map[Segment]map[string]bool{}
	impossibleSignals := map[Segment]map[string]bool{}
	for _, v := range AllSegments {
		potentialSignals[v] = map[string]bool{}
		impossibleSignals[v] = map[string]bool{}
	}
	return &DisplayDecoder{
		PotentialSignals:  potentialSignals,
		ImpossibleSignals: impossibleSignals,
		RecordedDigits:    map[int]bool{},
	}
}

func ContainsSegment(s []Segment, v Segment) bool {
	for _, w := range s {
		if w == v {
			return true
		}
	}
	return false
}

func SortSegments(s []Segment) {
	// Yes, bubble sort
	end := len(s)
	for swapped := true; swapped; {
		swapped = false

		for i := 1; i < end; i++ {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
				swapped = true
			}
		}

		end -= 1
	}
}

func (d *DisplayDecoder) RecordSignal(signal string) {
	signal = strings.TrimSpace(signal)

	recordDigit := func(digit int, segments []Segment) {
		if _, ok := d.RecordedDigits[digit]; ok {
			// We've already recorded this digit
			return
		}
		d.RecordedDigits[digit] = true

		signals := strings.Split(signal, "")
		for _, v := range segments {
			for _, w := range signals {
				d.PotentialSignals[v][w] = true
			}
		}
		for _, v := range AllSegments {
			if ContainsSegment(segments, v) {
				continue
			}

			for _, w := range signals {
				d.ImpossibleSignals[v][w] = true
			}
		}
	}

	switch len(signal) {
	case 2: // 1
		recordDigit(1, []Segment{TopRight, BottomRight})

	case 3: // 7
		recordDigit(7, []Segment{Top, TopRight, BottomRight})

	case 4: // 4
		recordDigit(4, []Segment{TopLeft, TopRight, Middle, BottomRight})

	case 7: // 8
		recordDigit(8, []Segment{Top, TopLeft, TopRight, Middle, BottomLeft, BottomRight, Bottom})
	}
}

func (d *DisplayDecoder) EliminateImpossibleSignals() {
	for i, v := range d.PotentialSignals {
		for j := range v {
			if _, ok := d.ImpossibleSignals[i][j]; ok {
				delete(v, j)
			}
		}
	}
}

type Ssantehu struct {
	s Segment
	v []string
}

func GeneratePermutation(k chan<- *Display, m map[Segment]string, s []Ssantehu) {
	if len(s) == 0 {
		k <- CreateDisplay(m)
		return
	}

	p := s[0]
	for _, v := range p.v {
		m[p.s] = v
		GeneratePermutation(k, m, s[1:])
		delete(m, p.s)
	}
}

func (d *DisplayDecoder) GetAllPossiblePermutations() chan *Display {
	var p []Ssantehu
	for i, v := range d.PotentialSignals {
		q := make([]string, 0, len(v))
		for k := range v {
			q = append(q, k)
		}
		p = append(p, Ssantehu{s: i, v: q})
	}
	k := make(chan *Display)
	go GeneratePermutation(k, map[Segment]string{}, p)
	return k
}

type Display struct {
	Signals map[Segment]string
}

func CreateDisplay(signals map[Segment]string) *Display {
	s := map[Segment]string{}
	for _, v := range AllSegments {
		s[v] = signals[v]
	}
	return &Display{s}
}

func (d *Display) GetMatchingDigit(signal string) int {
	var segments []Segment
	for _, v := range strings.Split(signal, "") {
		for k, w := range d.Signals {
			if v == w {
				segments = append(segments, k)
				break
			}
		}
	}

	SortSegments(segments)

	for i, v := range Digits {
		if len(segments) != len(v) {
			continue
		}

		matches := true

		for j, w := range v {
			if segments[j] != w {
				matches = false
				break
			}
		}

		if matches {
			return i
		}
	}

	return -1
}

func main() {
	var totalOutput int
	libaoc.ReadInputFileByLine(func(line string) {
		decoder := CreateDisplayDecoder()
		split := strings.Split(line, "|")

		uniqueSignals := strings.Split(split[0], " ")
		for _, v := range uniqueSignals {
			decoder.RecordSignal(v)

			if len(decoder.RecordedDigits) >= 4 {
				break
			}
		}

		decoder.EliminateImpossibleSignals()

		k := decoder.GetAllPossiblePermutations()
		var display *Display
		for v := range k {
			match := true
			for _, w := range uniqueSignals {
				if len(w) == 0 {
					continue
				}

				if v.GetMatchingDigit(w) < 0 {
					match = false
					break
				}
			}
			if match {
				display = v
				break
			}
		}

		outputs := strings.Split(split[1], " ")
		var output int
		for _, v := range outputs {
			if len(v) == 0 {
				continue
			}

			output = output*10 + display.GetMatchingDigit(v)
		}
		totalOutput += output
	})
	println(totalOutput)
}
