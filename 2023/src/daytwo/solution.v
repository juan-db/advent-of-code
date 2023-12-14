module daytwo

import os
import strconv
import math

pub fn part1(file string) int {
	lines := os.read_lines(file) or {
		panic('failed to read file: ${err}')
	}

	limits := Round{12, 13, 14}

	mut sum := 0

	line_loop: for l in lines {
		id, rounds := get_game_data(l)

		for r in rounds {
			round := parse_round(r)
			if round.r > limits.r || round.g > limits.g || round.b > limits.b {
				continue line_loop
			}
		}

		sum += id
	}

	return sum
}

pub fn part2(file string) int {
	lines := os.read_lines(file) or {
		panic('failed to read file: ${err}')
	}

	mut sum := 0

	for l in lines {
		_, rounds := get_game_data(l)

		mut min_round := Round{0, 0, 0}
		for r in rounds {
			new_round := parse_round(r)
			min_round = Round{
				math.max(min_round.r, new_round.r)
				math.max(min_round.g, new_round.g)
				math.max(min_round.b, new_round.b)
			}
		}

		sum += min_round.r * min_round.g * min_round.b
	}

	return sum

}

fn get_game_data(line string) (int, []string) {
	parts := line.split(':')
	if parts.len < 2 {
		panic('misformed line: ${line}')
	}

	id := strconv.atoi(parts[0].split(' ')[1]) or {
		panic('error parsing id: ${err}')
	}
	rounds := parts[1].split(';')

	return id, rounds
}

fn parse_round(round string) Round {
	cubes := round.split(',')
	mut r := Round{0, 0, 0}
	for c in cubes {
		parts := c.trim_space().split(' ')
		count := strconv.atoi(parts[0]) or { panic('error parsing count: ${err}') }
		color := parts[1]
		r = match color {
			'red'   { Round{ ...r, r: count } }
			'green' { Round{ ...r, g: count } }
			'blue'  { Round{ ...r, b: count } }
			else    { panic('unknown color "${color}"') }
		}
	}
	return r
}

struct Game {
	id int
	rounds []Round
}

struct Round {
	r int
	g int
	b int
}
