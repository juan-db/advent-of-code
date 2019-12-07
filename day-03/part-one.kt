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

	private val direction =
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
			Point(a.a.x, b.a.y)
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

fun findNextCoordinate(original: Point, instruction: String): Point =
	when (instruction[0]) {
		// There HAS to be a better way to do this
		'R' -> Point(original.x + instruction.substring(1).toInt(), original.y)
		'U' -> Point(original.x, original.y + instruction.substring(1).toInt())
		'L' -> Point(original.x - instruction.substring(1).toInt(), original.y)
		'D' -> Point(original.x, original.y - instruction.substring(1).toInt())
		else -> throw Exception("Unrecognized direction: ${instruction[0]}")
	}

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: ./part-one <filename>")
		return
	}

	// TODO: I want to try a "functional" approach once it's done.
	// might use quite a bit more memory.
	val program = File(args[0])
		.readLines()
		.map { it.split(",") }

	val intersections = mutableListOf<Point>()
	var a = Point.ORIGIN.copy()
	for (aInstruction in program[0]) {
		val aLine = Line(a, findNextCoordinate(a, aInstruction))
		var b = Point.ORIGIN.copy()
		for (bInstruction in program[1]) {
			val bLine = Line(b, findNextCoordinate(b, bInstruction))
			val intersection = aLine.intersection(bLine)
			if (intersection != null) {
				intersections.add(intersection)
			}
			b = bLine.b
		}
		a = aLine.b
	}
	println(
		intersections
			.filterNot { it == Point.ORIGIN }
			.map {
				println(it)
				Point.ORIGIN.distance(it).let {
					println(it)
					it
				} }
			.min()
	)
}
