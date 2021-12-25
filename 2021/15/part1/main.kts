import kotlin.math.abs
import java.io.File
import java.util.*

data class Point(val x: Int, val y: Int, val risk: Int) {
	fun distance(p: Point) = abs(p.x - x) + abs(p.y - y)
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

val map = File(args[0])
	.readLines()
	.mapIndexed { i, l ->
		l.mapIndexed { j, v ->
			Point(j, i, v - '0')
		}
	}

fun findPath(): PathPoint {
	fun visitPoint(p: PathPoint): List<PathPoint> =
		listOf(
			p.p.x to p.p.y - 1, p.p.x to p.p.y + 1,
			p.p.x - 1 to p.p.y, p.p.x + 1 to p.p.y
		)
		.filter {
			it.first >= 0 && it.first < map[0].size
				&& it.second >= 0 && it.second < map.size
		}
		.map {
			val next = map[it.second][it.first]
			PathPoint(next, p, p.score + next.risk, p.end)
		}
		.filter {
			!p.visited(it)
		}

	val end = map[map.size - 1][map[0].size - 1]
	val traversed = mutableSetOf<Point>()
	val queue = PriorityQueue<PathPoint>()
	queue.add(PathPoint(map[0][0], null, 0, end))

	while (queue.isNotEmpty()) {
		val node = queue.poll()
		traversed.add(node.p)
		val connected = visitPoint(node)
		val endPath = connected.firstOrNull { it.p == end }
		if (endPath != null) {
			return endPath
		}
		queue.addAll(connected.filter { !traversed.contains(it.p) })
	}

	throw Exception("No valid path")
}

findPath().score