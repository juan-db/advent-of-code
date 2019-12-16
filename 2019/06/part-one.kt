package partone

import java.io.File
import kotlin.system.exitProcess

/**
 * @property name Name of the object.
 * @property orbit The object this object is orbiting around.
 * @property orbits Objects orbiting around this object.
 */
class SpaceObject(
	val name: String, var orbit: SpaceObject?, val orbits: MutableList<SpaceObject>
) {
	fun distanceToCore(): Int {
		var obj = orbit
		var total = 0
		while (obj != null) {
			total += 1
			obj = obj.orbit
		}
		return total
	}

	fun printTree(depth: Int = 0, printIndent: Boolean = false) {
		if (printIndent) {
			print(" ".repeat(depth))
		}
		if (depth != 0) {
			print(" - ")
		}

		print(name)
		for (i in 0 until orbits.size) {
			val child = orbits[i]
			val newDepth = (if (depth != 0) 3 else 0) + depth + name.length
			if (i != 0) {
				println()
				child.printTree(newDepth, true)
			} else {
				child.printTree(newDepth)
			}
		}
	}
}

fun createObjects(map: Iterable<String>): Map<String, SpaceObject> {
	val objects = mutableMapOf<String, SpaceObject>()
	for (obj in map) {
		objects.computeIfAbsent(obj) { SpaceObject(it, null, mutableListOf()) }
	}
	return objects
}

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: java -jar part-one.jar <program file>")
		exitProcess(1)
	}

	val map = File(args[0]).readLines().map { it.split(')') }
	val objects = createObjects(map.flatten())
	for (orbit in map) {
		// Not sure why `?: error` is preferred to `!!` but ok IDEA.
		val orbitee = objects[orbit[0]] ?: error("")
		val orbiter = objects[orbit[1]] ?: error("")
		orbitee.orbits.add(orbiter)
		orbiter.orbit = orbitee
	}

	val orbits = objects.values.sumBy { it.distanceToCore() }
	println(orbits)
}