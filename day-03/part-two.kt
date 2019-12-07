import java.io.File
import kotlin.math.abs

// I know it's dumb that I convert from instruction of form <direction><count>,
// to line, then recalculate direction. I don't really want to do invest the
// time to fix it right now.

data class Point(val x: Int, val y: Int) {
	companion object {
		val ORIGIN = Point(0, 0)
	}

	fun distance(other: Point) = abs(x - other.x) + abs(y - other.y)
}

data class Line(val a: Point, val b: Point) {
	enum class Direction {
		UP, DOWN, LEFT, RIGHT
	}

	enum class SimpleDirection {
		VERTICAL, HORIZONTAL
	}

	public val direction =
		if (a.x == b.x) {
			if (a.y < b.y) Direction.UP else Direction.DOWN
		} else {
			if (a.x < b.x) Direction.RIGHT else Direction.LEFT
		}

	private val simpleDirection =
		if (direction == Direction.UP || direction == Direction.DOWN) {
			SimpleDirection.VERTICAL
		} else {
			SimpleDirection.HORIZONTAL
		}

	val length = abs((a.x - b.x) + (a.y - b.y))

	// Not sure if lines that run on top of each other are considered
	// "intersecting" but I doubt the input will need me to cater for that case
	// so I'm gonna pretend it's not possible.
	fun intersection(other: Line): Point? {
		if (this.simpleDirection == other.simpleDirection) {
			return null
		}

		val (a, b) = run {
			val x = this.normalize()
			val y = other.normalize()

			if (x.simpleDirection == SimpleDirection.HORIZONTAL) {
				Pair(x, y)
			} else {
				Pair(y, x)
			}
		}

		return if (
		// Horizontal line starts before and ends after the vertical line.
			a.a.x <= b.a.x && a.b.x >= b.a.x
			// Horizontal line is in between the start and end y of the vertical line.
			&& a.a.y >= b.a.y && a.a.y <= b.b.y
		) {
			Point(b.a.x, a.a.y)
		} else {
			null
		}
	}

	/**
	 * Creates a line where the direction is either UP or RIGHT.
	 * This function will always create and return a new line.
	 * For the returned Line, the following will hold true:
	 * `a.x <= b.x && a.y <= b.y`
	 */
	private fun normalize(): Line {
		return when (direction) {
			Direction.UP, Direction.RIGHT -> Line(a.copy(), b.copy())
			else -> Line(b.copy(), a.copy())
		}
	}
}

fun getNextCoordinate(original: Point, instruction: String): Point =
	when (instruction[0]) {
		// There HAS to be a better way to do this
		'R' -> Point(original.x + instruction.substring(1).toInt(), original.y)
		'U' -> Point(original.x, original.y + instruction.substring(1).toInt())
		'L' -> Point(original.x - instruction.substring(1).toInt(), original.y)
		'D' -> Point(original.x, original.y - instruction.substring(1).toInt())
		else -> throw Exception("Unrecognized direction: ${instruction[0]}")
	}


fun generatePoints(instructions: List<String>) =
	sequence {
		var last = Point.ORIGIN.copy()
		yield(last)

		for (instruction in instructions) {
			last = getNextCoordinate(last, instruction)
			yield(last)
		}
	}

fun generateLines(instructions: List<String>) =
	generatePoints(instructions)
		.zipWithNext()
		.map { Line(it.first, it.second) }

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: ./part-one <filename>")
		return
	}

	val program = File(args[0])
		.readLines()
		.map { it.split(",") }

	// This is not the optimal way of doing this but it's the simplest I can
	// think of.
	var fewestSteps = Int.MAX_VALUE
	var aDistance = 0
	for (a in generateLines(program[0])) {
		var bDistance = 0
		for (b in generateLines(program[1])) {
			val intersection = a.intersection(b)
			if (intersection != null && intersection != Point.ORIGIN) {
				val distance = aDistance + bDistance +
					intersection.distance(
						when (a.direction) {
							Line.Direction.UP, Line.Direction.RIGHT -> a.a
							Line.Direction.DOWN, Line.Direction.LEFT -> a.b
						}
					) +
					intersection.distance(
						when (b.direction) {
							Line.Direction.UP, Line.Direction.RIGHT -> b.a
							Line.Direction.DOWN, Line.Direction.LEFT -> b.b
						}
					)
				if (distance < fewestSteps) {
					fewestSteps = distance
				}
				break
			}
			bDistance += b.length
		}
		aDistance += a.length
	}
	println(fewestSteps)
}
