module main

import os
// import dayone
import daytwo


fn main() {
	if os.args.len < 2 {
		eprintln('Usage: ${os.args[0]} <input filepath>')
		exit(1)
	}

	file := os.args[1]
	// println(dayone.part_1(file))
	// println(dayone.part_2(file))
	println(daytwo.part2(file))
}
