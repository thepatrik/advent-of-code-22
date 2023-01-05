package seven

import (
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

func PartOne(filename string) int {
	filem := parse(filename)

	sum := 0
	for _, size := range filem.getDirSize() {
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func PartTwo(filename string) int {
	filem := parse(filename)

	dirsize := filem.getDirSize()
	smallest := dirsize["/"]
	for _, size := range dirsize {
		if size+70000000-dirsize["/"] >= 30000000 && size < smallest {
			smallest = size
		}
	}

	return smallest
}

type FileMap map[string]int

func (filemap FileMap) getDirSize() FileMap {
	m := map[string]int{}
	for f, s := range filemap {
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			m[d] += s
		}
		m["/"] += s
	}

	return m
}

func parse(filename string) FileMap {
	strslice := parser.ReadFile(filename)
	filemap := make(FileMap)
	currentDir := ""

	for _, line := range strslice {
		if strings.HasPrefix(line, "$ cd") {
			currentDir = path.Join(currentDir, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			fields := strings.Fields(line)
			size, _ := strconv.Atoi(fields[0])
			name := fields[1]
			filemap[currentDir+"/"+name] = size
		}
	}

	return filemap
}
