package parttwo

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
	fun find(name: String): SpaceObject? =
		if (this.name == name) {
			this
		} else {
			orbits
				.asSequence()
				.map { it.find(name) }
				.firstOrNull { it != null }
		}

	fun distance(name: String): Int? =
		if (this.name == name) {
			0
		} else {
			orbits
				.asSequence()
				.map { it.distance(name) }
				.firstOrNull { it != null }
				?.plus(1)
		}
}

fun createObjects(map: Iterable<String>): Map<String, SpaceObject> {
	val objects = mutableMapOf<String, SpaceObject>()
	for (obj in map) {
		objects.computeIfAbsent(obj) { SpaceObject(it, null, mutableListOf()) }
	}
	return objects
}

fun constructOrbits(map: List<List<String>>, objects: Map<String, SpaceObject>) {
	for (orbit in map) {
		// Not sure why `?: error` is preferred to `!!` but ok IDEA.
		val orbitee = objects[orbit[0]] ?: error("")
		val orbiter = objects[orbit[1]] ?: error("")
		orbitee.orbits.add(orbiter)
		orbiter.orbit = orbitee
	}
}

fun findDistance(root: SpaceObject): Int {
	val you = root.find("YOU") ?: error("Couldn't find YOU.")
	var current = you.orbit
	while (current != null) {
		val distance = current.distance("SAN")
		if (distance != null) {
			return distance + current.distance("YOU")!! - 2
		}
		current = current.orbit
	}
	error("Couldn't find SAN.")
}

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: java -jar part-one.jar <program file>")
		exitProcess(1)
	}

	val map = File(args[0]).readLines().map { it.split(')') }
	val objects = createObjects(map.flatten())
	constructOrbits(map, objects)
	val root = objects.values.first { it.orbit == null }
	println(findDistance(root))
}