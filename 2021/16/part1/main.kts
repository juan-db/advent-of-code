import java.io.File

class Reader(val data: String) {
	var cursor = 0
		private set

	fun read(n: Int): String {
		cursor += n
		return data.slice(cursor - n until cursor)
	}
}

data class Packet(val version: Int, val type: Int, val number: Int, val children: List<Packet>) {
	companion object {
		fun parse(reader: Reader): Packet {
			val version = reader.read(3).toInt(2)
			val type = reader.read(3).toInt(2)

			return if (type == 4) {
				// literal
				var number = 0
				do {
					val last = reader.read(1) == "0"
					number = (number shl 4) + reader.read(4).toInt(2)
				} while (!last)

				Packet(version, type, number, emptyList())
			} else {
				// operator
				val lengthType = reader.read(1).first() - '0'
				val children = mutableListOf<Packet>()
				if (lengthType == 0) {
					val length = reader.read(15).toInt(2)
					val start = reader.cursor
					while (reader.cursor - start < length) {
						children.add(parse(reader))
					}
				} else {
					val subpackets = reader.read(11).toInt(2)
					while (children.size < subpackets) {
						children.add(parse(reader))
					}
				}

				Packet(version, type, 0, children)
			}
		}
	}
}

val hexToBin = mapOf(
	'0' to "0000", '1' to "0001", '2' to "0010", '3' to "0011",
	'4' to "0100", '5' to "0101", '6' to "0110", '7' to "0111",
	'8' to "1000", '9' to "1001", 'A' to "1010", 'B' to "1011",
	'C' to "1100", 'D' to "1101", 'E' to "1110", 'F' to "1111",
)

val hex = File(args[0]).readLines().first()

var binary = ""
for (c in hex) {
	binary += hexToBin[c]
}

val packet = Packet.parse(Reader(binary))
fun sumVersion(p: Packet): Int {
	var sum = 0
	for (c in p.children) {
		sum += sumVersion(c)
	}
	return sum + p.version
}
sumVersion(packet)