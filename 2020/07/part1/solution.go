package main

func main() {
	bags := map[string]*bag{}
	ReadInputFileByLine(func(line string) {
		graphRule(line, &bags)
	})
}
