import java.io.File
import kotlin.system.exitProcess

data class Parameter(val mode: Mode, val value: Int, val position: Int?) {
	enum class Mode {
		POSITION, IMMEDIATE
	}
}

/**
 * @return True if execution of the program is completed as a result of the
 * execution of this function.
 */
typealias InstructionFunction = (program: Program, args: List<Parameter>) -> Boolean

data class Instruction(
	val instruction: KnownInstruction, val parameterModes: Array<Parameter.Mode>
) {
	enum class KnownInstruction(
		val opcode: Int, val argCount: Int, val op: InstructionFunction
	) {
		ADD(1, 3, { program, args ->
			if (args[0].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("First argument to ADD must be a position argument.")
			}

			program[args[0].value] = args[1].value + args[2].value
			true
		}),
		MUL(2, 3, { program, args ->
			if (args[0].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("First argument to MUL must be a position argument.")
			}

			program[args[0].value] = args[1].value * args[2].value
			true
		}),
		READ(3, 1, { program, args ->
			if (args[0].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("Only position argument mode is accepted for the READ instruction.")
			}


			var num: Int? = null
			while (num == null) {
				try {
					print("Please enter a number: ")
					num = readLine()?.toInt()
				} catch (e: java.lang.NumberFormatException) {
					println("That is not a valid number!")
				}
			}
			program[args[0].value] = num as Int
			true
		}),
		PRINT(4, 1, { program, args ->
			// Yes, it is useless to pass an immediate value to print, but why
			// not?
			// if (args[0].mode != Parameter.Mode.POSITION) {
			//  throw InvalidArgumentMode("Only position argument mode is accepted for the PRINT instruction.")
			// }

			println(program[args[0].value])
			true
		}),
		HALT(99, 0, { _, _ -> true });

		operator fun invoke(program: Program, args: List<Parameter>) = op(program, args)

		companion object {
			operator fun get(opcode: Int) = values().first { it.opcode == opcode }
		}
	}

	companion object Companion {
		fun parse(code: Int): Instruction {
			val instruction = KnownInstruction[code % 100]
			val explicitModes = (code / 100)
				.toString()
				.chunked(1)
				.map { Parameter.Mode.values()[it.toInt()] }
				.toTypedArray()
			val paddingModes = List(
				instruction.argCount - explicitModes.size
			) { Parameter.Mode.values()[0] }
			return Instruction(instruction, explicitModes + paddingModes)
		}
	}

	operator fun invoke(program: Program): Boolean {
		val args = parameterModes.map {
			val int = program.poll()
			if (it == Parameter.Mode.IMMEDIATE) {
				Parameter(it, int, null)
			} else {
				Parameter(it, program[int], int)
			}
		}
		return instruction(program, args)
	}

	override fun equals(other: Any?): Boolean {
		if (this === other) return true
		if (javaClass != other?.javaClass) return false

		other as Instruction

		if (instruction != other.instruction) return false
		if (!parameterModes.contentEquals(other.parameterModes)) return false

		return true
	}

	override fun hashCode(): Int {
		var result = instruction.hashCode()
		result = 31 * result + parameterModes.contentHashCode()
		return result
	}
}

class InvalidArgumentMode(message: String) : Exception(message)

class Program(collection: Collection<Int>) : ArrayList<Int>(collection) {
	private var ip = 0

	fun poll(): Int = this[ip++]
}

fun List<Int>.toProgram(): Program = Program(this)

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: java -jar part-one.jar <program file>")
		exitProcess(1)
	}

	val program = File(args[0])
		.readText()
		.trim()
		.split(",")
		.map { it.toInt() }
		.toProgram()

	do {
		val instruction = Instruction.parse(program.poll())
		val done = instruction(program)
	} while (!done)
}


//const ops = {
//	1: mathOp((a, b) => a + b),
//	2: mathOp((a, b) => a * b),
//	99: () => true
//};
//
//for (let a = 0; a <= 99; ++a) {
//	for (let b = 0; b <= 99; ++b) {
//		const p = [program[0], a, b, ...program.slice(3)];
//		if (execute(p) === 19690720) {
//			console.log(100 * a + b);
//			process.exit(0);
//		}
//	}
//}
//
//function mathOp(op) {
//	return function (memory, a, b, dst) {
//		const args = [...arguments];
//
//		if (args.length < 4) {
//			throw new Error(`Expected three (3) arguments (two (2) operands and a destination) but only received ${args.length - 1}: ${args.slice(1).join(", ")}`);
//		}
//
//		for (const [i, v] of args.slice(1).entries()) {
//		if (v < 0 || v >= memory.length) {
//			throw new Error(`Invalid address provided as ${["first", "second", "third"][i]} argument to instruction: ${v}`);
//		}
//	}
//
//		memory[dst] = op(memory[a], memory[b]);
//	}
//}
//
//function execute(memory) {
//	for (let ip = 0; ip < memory.length; ip += 4) {
//		const opcode = memory[ip];
//		const op = ops[opcode];
//		if (!op) {
//			throw new Error(`Unrecognized opcode: ${opcode}`);
//		}
//
//		const args = memory.slice(ip + 1, ip + 4);
//		const done = op.apply(null, [memory, ...args]);
//		if (done) {
//			return memory[0];
//		}
//	}
//
//	throw new Error("No halt instruction in code.");
//}
