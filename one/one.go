package one

import (
	"sort"
	"strconv"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
	"github.com/thepatrik/advent-of-code-22/pkg/slice"
)

type elf struct {
	id   int
	cals []int
}

func PartOne(filename string) int {
	elves := parse(filename)

	max := 0
	maxElf := 0

	for _, elf := range elves {
		cals := slice.Sum(elf.cals)

		if cals > max {
			max = cals
			maxElf = elf.id
		}
	}

	return maxElf
}

func PartTwo(filename string) int {
	elves := parse(filename)

	sort.SliceStable(elves, func(i, j int) bool {
		return slice.Sum(elves[j].cals) < slice.Sum(elves[i].cals)
	})

	return slice.Sum(elves[0].cals) + slice.Sum(elves[1].cals) + slice.Sum(elves[2].cals)
}

func parse(filename string) []elf {
	strslice := parser.ReadFile(filename)

	id := 1

	elves := make([]elf, 0)
	cals := make([]int, 0)

	for i := 0; i <= len(strslice); i++ {
		line := func() string {
			if i == len(strslice) {
				return ""
			}
			return strslice[i]
		}()

		if line != "" {
			intcals, _ := strconv.Atoi(line)
			cals = append(cals, intcals)
			continue
		}

		elf := elf{
			id:   id,
			cals: cals,
		}
		elves = append(elves, elf)
		id++

		cals = make([]int, 0)
	}

	return elves
}
