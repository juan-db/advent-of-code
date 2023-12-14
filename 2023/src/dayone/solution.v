module dayone

import os
import strconv

pub fn part_1(file string) int {
	lines := os.read_lines(file) or {
		panic('failed to read file: ${err}')
	}

	mut sum := 0
	for l in lines {
		mut first := ?rune(none)
		mut last := rune(0)
		for c in l {
			if !c.is_digit() {
				continue
			}

			if first == none {
				first = c
			}
			last = c
		}
		value := '${first or { panic('failed to find first number') }}${last}'
		sum += strconv.atoi(value) or {
			panic('${value} is not a valid number')
		}
	}
	return sum
}

pub fn part_2(file string) int {
	lines := os.read_lines(file) or {
		panic('couldn\'t read file: ${err}')
	}

	words := [
		'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine',
		'1', '2', '3', '4', '5', '6', '7', '8', '9',
	]

	mut sum := 0
	for l in lines {
		mut matchers := words.map(Matcher.new(it))
		s := find_match(l, mut matchers) or {
			panic('couldn\'t find a number from the front')
		}

		mut reverse_matchers := words.map(Matcher.new(it.reverse()))
		e := find_match(l.reverse(), mut reverse_matchers) or {
			panic('couldn\'t find a number from the end')
		}

		value := '${s}${e}'
		sum += strconv.atoi(value) or {
			panic('${value} is not a valid number')
		}
	}
	return sum
}

fn find_match(line string, mut matchers []Matcher) ?string {
	for i, v in line {
		for mut m in matchers {
			if _ := m.check(i, v) {
				return m.get_value()
			}
		}
	}
	return none
}

struct Match {
	start int
mut:
	count int
}

struct Matcher {
	pattern string
mut:
	matches []Match = []
}

fn Matcher.new(pattern string) Matcher {
	return Matcher { pattern, [] }
}

fn (matcher Matcher) get_value() string {
	return match matcher.pattern {
		 'one', 'eno'     { '1' }
		 'two', 'owt'     { '2' }
		 'three', 'eerht' { '3' }
		 'four', 'ruof'   { '4' }
		 'five', 'evif'   { '5' }
		 'six', 'xis'     { '6' }
		 'seven', 'neves' { '7' }
		 'eight', 'thgie' { '8' }
		 'nine', 'enin'   { '9' }
		 else    { matcher.pattern }
	}
}

fn (mut matcher Matcher) check(index int, value u8) ?int {
	matcher.matches << Match{start: index, count: 0}

	for i := 0; i < matcher.matches.len; i++ {
		mut m := &matcher.matches[i]
		if matcher.pattern[m.count] != value {
			matcher.matches.delete(i)
			i -= 1
			continue
		}

		m.count += 1
		if m.count == matcher.pattern.len {
			return m.start
		}
	}

	return none
}
