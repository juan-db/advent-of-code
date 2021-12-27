import kotlin.math.abs
import java.io.File
import java.util.*

data class Point(val x: Int, val y: Int, val risk: Int) {
	fun distance(p: Point) = abs(p.x - x) + abs(p.y - y)
}

class Map(val map: List<List<Int>>) {
	val width = map.first().size
	val maxWidth = width * 5

	val height = map.size
	val maxHeight = height * 5

	fun getPoint(x: Int, y: Int): Point {
		val excess = x / width + y / height
		val risk = map[y % height][x % width]
		var newRisk = risk + excess
		if (newRisk > 9) {
			newRisk -= 9
		}
		return Point(x, y, newRisk)
	}
}

data class PathPoint(val p: Point, val from: PathPoint?, val score: Int, val end: Point) : Comparable<PathPoint> {
	override fun compareTo(other: PathPoint): Int =
		(score + p.distance(end)) - (other.score + other.p.distance(end))

	fun visited(p: PathPoint): Boolean {
		var q: PathPoint? = this
		while (q != null) {
			if (p.p.x == q.p.x && p.p.y == q.p.y) {
				return true
			}
			q = q.from
		}
		return false
	}
}

fun printPath(p: PathPoint): Int {
	val depth = if (p.from != null) {
		printPath(p.from)
	} else {
		0
	}
	println("${"%3d".format(depth)}: ${p.p.x}, ${p.p.y} (${p.p.risk}): ${p.score}")
	return depth + 1
}

val map = Map(
	File(args[0])
		.readLines()
		.map { l -> l.map { v -> v - '0' } }
)

fun findPath(): PathPoint {
	fun visitPoint(p: PathPoint): List<PathPoint> =
		listOf(
			p.p.x to p.p.y - 1, p.p.x to p.p.y + 1,
			p.p.x - 1 to p.p.y, p.p.x + 1 to p.p.y
		)
			.filter {
				it.first >= 0 && it.first < map.maxWidth
				&& it.second >= 0 && it.second < map.maxHeight
			}
			.map {
				val next = map.getPoint(it.first, it.second)
				PathPoint(next, p, p.score + next.risk, p.end)
			}
			.filter {
				!p.visited(it)
			}

	val start = map.getPoint(0, 0)
	val end = map.getPoint(map.width * 5 - 1, map.height * 5 - 1)

	val frontier = mutableSetOf<Point>()
	frontier.add(start)

	val queue = PriorityQueue<PathPoint>()
	queue.add(PathPoint(start, null, 0, end))

	while (queue.isNotEmpty()) {
		val node = queue.poll()
		val connected = visitPoint(node)

		val endPath = connected.firstOrNull { it.p == end }
		if (endPath != null) {
			return endPath
		}

		val unvisited = connected.filter { !frontier.contains(it.p) }
		frontier.addAll(unvisited.map { it.p })
		queue.addAll(unvisited)
	}

	throw Exception("No valid path")
}

val path = findPath()
printPath(path)
path.score
