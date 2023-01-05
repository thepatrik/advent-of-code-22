package six

import "github.com/thepatrik/advent-of-code-22/pkg/parser"

func PartOne(filename string) int {
	buffer := parser.ReadFile(filename)[0]
	return findMarker(buffer, 4)
}

func PartTwo(filename string) int {
	buffer := parser.ReadFile(filename)[0]
	return findMarker(buffer, 14)
}

func findMarker(buffer string, length int) int {
	findChar := func(char byte, buffer string) bool {
		for i := 0; i < len(buffer); i++ {
			if char == buffer[i] {
				return true
			}
		}

		return false
	}

outer:
	for i := 0; i < len(buffer); i++ {
		for j := 0; j < length-1; j++ {
			if findChar(buffer[i+j], buffer[i+j+1:i+length]) {
				continue outer
			}
		}

		return i + length
	}

	return 0
}
