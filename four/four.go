package four

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type section struct {
	min int
	max int
}

func PartOne(filename string) int {
	sections := parse(filename)
	overlapping := 0

	for i := 0; i < len(sections); i = i + 2 {
		section1 := sections[i]
		section2 := sections[i+1]

		if section1.min >= section2.min && section1.max <= section2.max {
			overlapping += 1
		} else if section2.min >= section1.min && section2.max <= section1.max {
			overlapping += 1
		}
	}

	return overlapping
}

func PartTwo(filename string) int {
	sections := parse(filename)
	contained := 0

	isWithin := func(number int, min int, max int) bool {
		return number >= min && number <= max
	}

	for i := 0; i < len(sections); i = i + 2 {
		section1 := sections[i]
		section2 := sections[i+1]

		contained += func() int {
			if isWithin(section1.min, section2.min, section2.max) {
				return 1
			}
			if isWithin(section1.max, section2.min, section2.max) {
				return 1
			}
			if isWithin(section2.min, section1.min, section1.max) {
				return 1
			}
			if isWithin(section2.max, section1.min, section1.max) {
				return 1
			}
			return 0
		}()
	}

	return contained
}

func parse(filename string) []section {
	strslice := parser.ReadFile(filename)
	sections := make([]section, 0)

	for _, line := range strslice {
		split1 := strings.Split(line, ",")
		split2 := strings.Split(split1[0], "-")
		split3 := strings.Split(split1[1], "-")

		min, _ := strconv.Atoi(split2[0])
		max, _ := strconv.Atoi(split2[1])

		section1 := section{
			min: min,
			max: max,
		}
		sections = append(sections, section1)

		min, _ = strconv.Atoi(split3[0])
		max, _ = strconv.Atoi(split3[1])
		section2 := section{
			min: min,
			max: max,
		}
		sections = append(sections, section2)
	}

	return sections
}
