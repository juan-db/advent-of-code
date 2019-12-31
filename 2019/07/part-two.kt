package parttwo

import java.io.File
import java.lang.Integer.max
import java.util.*
import kotlin.collections.ArrayList
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
			if (args[2].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("Last argument to ADD must be a position argument.")
			}

			program[args[2].position as Int] = args[0].value + args[1].value
			false
		}),
		MUL(2, 3, { program, args ->
			if (args[2].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("Last argument to MUL must be a position argument.")
			}

			program[args[2].position as Int] = args[0].value * args[1].value
			false
		}),
		READ(3, 1, { program, args ->
			if (args[0].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("Only position argument mode is accepted for the READ instruction.")
			}

			program[args[0].position as Int] = program.input.poll()
			false
		}),
		PRINT(4, 1, { program, args ->
			program.output.add(args[0].value)
			false
		}),
		JUMP_IF_TRUE(5, 2, { program, args ->
			if (args[0].value != 0) {
				program.ip = args[1].value
			}
			false
		}),
		JUMP_IF_FALSE(6, 2, { program, args ->
			if (args[0].value == 0) {
				program.ip = args[1].value
			}
			false
		}),
		LESS_THAN(7, 3, { program, args ->
			if (args[2].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("Third argument to less than must be positional.")
			}

			program[args[2].position as Int] = if (args[0].value < args[1].value) 1 else 0
			false
		}),
		EQUAL(8, 3, { program, args ->
			if (args[2].mode != Parameter.Mode.POSITION) {
				throw InvalidArgumentMode("Third argument to equal must be positional.")
			}

			program[args[2].position as Int] = if (args[0].value == args[1].value) 1 else 0
			false
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
			val explicitModes: Array<Parameter.Mode> = (code / 100)
				.takeIf { it > 0 }
				?.toString()
				?.chunked(1)
				?.map { Parameter.Mode.values()[it.toInt()] }
				?.toTypedArray() ?: arrayOf()
			val paddingModes = Array(
				instruction.argCount - explicitModes.size
			) { Parameter.Mode.values()[0] }

			return Instruction(
				instruction, explicitModes.reversedArray() + paddingModes
			)
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
	enum class State {
		INITIALIZED,
		RUNNING,
		WAITING_FOR_INPUT,
		HALTED
	}

	var ip = 0

	val input = LinkedList<Int>()
	val output = object : LinkedList<Int>() {
		override fun add(element: Int): Boolean {
			largestOutput = max(largestOutput ?: 0, element)
			return super.add(element)
		}
	}

	var largestOutput: Int? = null
		private set

	private var instruction: Instruction? = null

	var state: State = State.INITIALIZED
		private set

	fun poll(): Int = this[ip++]

	fun execute() {
		state = State.RUNNING
		do {
			if (instruction == null) {
				instruction = Instruction.parse(poll())
			}
			if (
				instruction!!.instruction == Instruction.KnownInstruction.READ
				&& input.size < 1
			) {
				state = State.WAITING_FOR_INPUT
				return
			}
			val done = instruction!!(this)
			instruction = null
		} while (!done)
		state = State.HALTED
	}
}

fun List<Int>.toProgram(): Program = Program(this)

fun Array<Int>.permutations(): Sequence<Array<Int>> = sequence {
	fun Array<Int>.swap(aIndex: Int, bIndex: Int) {
		val tmp = this[aIndex]
		this[aIndex] = this[bIndex]
		this[bIndex] = tmp
	}

	// Shamelessly stolen from RosettaCode
	val array = this@permutations
	val c = Array(array.size) { 0 }
	var plus = false

	yield(array)

	var i = 0
	while (i < array.size) {
		if (c[i] < i) {
			array.swap(
				if (i % 2 == 0) 0 else c[i],
				i
			)
			yield(array)
			plus = !plus
			c[i] += 1
			i = 0
		} else {
			c[i] = 0
			i += 1
		}
	}
}

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


	arrayOf(5, 6, 7, 8, 9)
		.permutations()
		.map { amplify(program, it) to it.clone() }
		.maxBy { it.first }!!
		.let { println("${it.first} with ${it.second.contentToString()}") }
}

fun amplify(originalProgram: Program, phaseSettings: Array<Int>): Int {
	val programs = phaseSettings
		.map { phaseSetting ->
			Program(originalProgram).also {
				it.input.add(phaseSetting)
			}
		}
	programs[0].input.push(0)

	var i = 0
	while (programs.any { it.state != Program.State.HALTED }) {
		val program = programs[i]
		val previous = programs[(i - 1).let { if (it >= 0) it else 4 }]
		while (previous.output.size > 0) {
			program.input.add(previous.output.poll())
		}
		program.execute()
		i = (i + 1).let { if (it < programs.size) it else 0 }
	}

	return programs.last().largestOutput ?: error("Last amp never output anything.")
}
