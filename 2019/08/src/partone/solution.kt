package partone

import java.io.File
import java.lang.IllegalArgumentException
import kotlin.system.exitProcess

class Image(val width: Int, val height: Int, val layers: Array<Layer>) {
	class Layer(val data: Array<Array<Int>>)

	companion object {
		fun read(filename: String) = with(File(filename)) {
			val reader = this.bufferedReader()

			val dimensions = reader
				.readLine()
				?.trim()
				?.split("*")
				?.map(String::trim)
				?.map(String::toIntOrNull) ?: emptyList()
			if (dimensions.size != 2) {
				error("First line must contain dimensions in form: <width>*<height>")
			}
			val (width, height) =
				try {
					dimensions.requireNoNulls()
				} catch (e: IllegalArgumentException) {
					error("Width and height must be integers.")
				}

			val dataLine = reader
				.readLine() ?: error("Second line must contain image data.")

			val data =
				try {
					dataLine
						.trim()
						.map { it.toString().toIntOrNull() }
						.requireNoNulls()
				} catch (e: IllegalArgumentException) {
					error("Image data must be all integers.")
				}
			if (data.size % (width * height) != 0) {
				error("Malformed image data. Image data should be any amount of layers of size `width * height`.")
			}

			val layers = data
				.chunked(width * height)
				.map { layer ->
					layer
						.chunked(width)
						.map { row ->
							row.toTypedArray()
						}
						.toTypedArray()
				}
				.map { Layer(it) }
				.toTypedArray()

			Image(width, height, layers)
		}
	}
}

fun main(args: Array<String>) {
	if (args.size != 1) {
		println("Usage: java -jar part-one.jar <program file>")
		exitProcess(1)
	}

	val flattenedLayerData = Image
		.read(args[0])
		.layers
		.map { it.data.flatten() }
		.minBy { layer -> layer.count { it == 0 } }!!
	val ones = flattenedLayerData.count { it == 1 }
	val twos = flattenedLayerData.count { it == 2 }
	println(ones * twos)
}
