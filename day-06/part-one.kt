import java.io.File
import kotlin.system.exitProcess

/**
 * @property name Name of the object.
 * @property orbit The object this object is orbiting around.
 * @property orbits Objects orbiting around this object.
 */
class SpaceObject(
	val name: String, val orbit: SpaceObject?, val orbits: MutableList<SpaceObject>
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

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: java -jar part-one.jar <program file>")
		exitProcess(1)
	}

	val map = File(args[0]).readLines().map { it.split(')') }

	val com = SpaceObject("COM", null, mutableListOf())
	val objects = mutableMapOf(com.name to com)
	for (orbit in map) {
		val orbitee = objects[orbit[0]] ?: com
		val obj = SpaceObject(orbit[1], orbitee, mutableListOf())
		orbitee.orbits.add(obj)
		objects[orbit[1]] = obj
	}
	com.printTree()
	println()
	println(objects.values.fold(0) { acc, obj -> acc + obj.distanceToCore() })
}