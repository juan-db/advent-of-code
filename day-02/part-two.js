// Why do I always overengineer *everything*?

if (process.argv.length != 3) {
	console.log("Usage: node solution.js <program file>");
	process.exit(1);
}

const ops = {
	1: mathOp((a, b) => a + b),
	2: mathOp((a, b) => a * b),
	99: () => true
};

const result = execute(
	require("fs")
		.readFileSync(process.argv[2], "utf8")
		.split(",")
		.map(x => +x)
);
console.log(result);

function mathOp(op) {
	return function (memory, a, b, dst) {
		const args = [...arguments];

		if (args.length < 4) {
			throw new Error(`Expected three (3) arguments (two (2) operands and a destination) but only received ${args.length - 1}: ${args.slice(1).join(", ")}`);
		}

		for (const [i, v] of args.slice(1).entries()) {
			if (v < 0 || v >= memory.length) {
				throw new Error(`Invalid address provided as ${["first", "second", "third"][i]} argument to instruction: ${v}`);
			}
		}

		memory[dst] = op(memory[a], memory[b]);
	}
}

function execute(memory) {
	for (let ip = 0; ip < memory.length; ip += 4) {
		const opcode = memory[ip];
		const op = ops[opcode];
		if (!op) {
			throw new Error(`Unrecognized opcode: ${opcode}`);
		}

		const args = memory.slice(ip + 1, ip + 4);
		const done = op.apply(null, [memory, ...args]);
		if (done) {
			return memory[0];
		}
	}

	throw new Error("No halt instruction in code.");
}
