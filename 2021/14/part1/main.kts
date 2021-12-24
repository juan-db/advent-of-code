import java.io.File

data class Rule(val a: Char, val b: Char, val toInsert: Char)

fun applyRules(rules: List<Rule>, polymer: MutableList<Char>) {
	var i = 0
	while (i < polymer.size - 1) {
		val a = polymer[i]
		val b = polymer[i + 1]
		for (r in rules) {
			if (a == r.a && b == r.b) {
				polymer.add(i + 1, r.toInsert)
				i += 2
				break
			}
		}
	}
}

val filename = args[0]
val lines = File(filename).readLines()
val polymer = lines[0].toCharArray().toMutableList()
val rules = lines
	.slice(IntRange(2, lines.size - 1))
	.map{
		val (pair, toInsert) = it.split(" -> ")
		val (a, b) = pair.toCharArray()
		Rule(a, b, toInsert.first())
	}
for (i in 0..9) {
	applyRules(rules, polymer)
}
val elementCounts = polymer
	.groupBy { it }
	.map { it.key to it.value.size }
	.sortedBy { it.second }
println(elementCounts.last().second - elementCounts.first().second)