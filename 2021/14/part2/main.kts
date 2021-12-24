import java.io.File

data class Rule(val a: Char, val b: Char, val toInsert: Char)

data class PolymerPair(val a: Char, val b: Char) {
	fun applyRule(rule: Rule): Pair<PolymerPair, PolymerPair>? {
		if (rule.a != a || rule.b != b) {
			return null
		}

		return PolymerPair(rule.a, rule.toInsert) to PolymerPair(rule.toInsert, rule.b)
	}
}

val filename = args[0]
val lines = File(filename).readLines()

var polymerPairs = lines[0]
	.toCharArray()
	.let { lineChars ->
		lineChars
			.zip(lineChars.slice(1 until lineChars.size))
			.map { PolymerPair(it.first, it.second) }
	}
	.groupBy { it }
	.map { it.key to it.value.size.toLong() }

val rules = lines
	.slice(IntRange(2, lines.size - 1))
	.map{
		val (pair, toInsert) = it.split(" -> ")
		val (a, b) = pair.toCharArray()
		Rule(a, b, toInsert.first())
	}

for (i in 0..39) {
	polymerPairs = polymerPairs
		.flatMap {
			rules
				.firstNotNullOfOrNull { r -> it.first.applyRule(r) }
				?.let { p -> listOf(p.first to it.second, p.second to it.second) }
				?: listOf(it)
		}
		.groupBy { it.first }
		.map { e -> e.key to e.value.sumOf { it.second.toLong() } }
}

polymerPairs
	.flatMap { listOf(it.first.a to it.second, it.first.b to it.second) }
	.groupBy { it.first }
	.map { e ->
		e.key to e.value
			.sumOf { it.second }
			.let { if (it % 2L == 0L) it / 2 else it / 2 + 1 }
	}
	.sortedBy { it.second }
	.let { println(it.last().second - it.first().second) }
