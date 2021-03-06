#!/usr/bin/node

if (process.argv.length != 3) {
	console.log("Usage: node solution.js <program file>");
	process.exit(1);
}

const program = require("fs")
	.readFileSync(process.argv[2], "utf8")
	.split(",")
	.map(x => +x);

function mathOp(op) {
	return function (p, a, b, dst) {
		const args = [...arguments];

		if (args.length < 4) {
			throw new Error(`Expected three (3) arguments (two (2) operands and a destination) but only received ${args.length - 1}: ${args.slice(1).join(", ")}`);
		}

		for (const [i, v] of args.slice(1).entries()) {
			if (v < 0 || v >= program.length) {
				throw new Error(`Invalid position provided as ${["first", "second", "third"][i]} operand to operation: ${v}`);
			}
		}

		p[dst] = op(p[a], p[b]);
	}
}

const ops = {
	1: mathOp((a, b) => a + b),
	2: mathOp((a, b) => a * b),
	99: () => {
		console.log(program[0]);
		process.exit(0);
	}
};

for (let ip = 0; ip < program.length; ip += 4) {
	const opcode = program[ip];
	const op = ops[opcode];
	if (!op) {
		throw new Error(`Unrecognized opcode: ${opcode}`);
	}

	let args = program.slice(ip + 1, ip + 4);
	op.apply(null, [program, ...args]);
}
