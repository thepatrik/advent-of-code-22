package eleven

import (
	"sort"
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type monkey struct {
	id             int
	items          []int
	operation      func(int) int
	divisiblyBy    int
	truthyMonkeyId int
	falsyMonkeyId  int
}

func (m monkey) destination(worry int) int {
	if worry%m.divisiblyBy == 0 {
		return m.truthyMonkeyId
	}
	return m.falsyMonkeyId
}

func PartOne(filename string) int {
	monkeys := parse(filename)

	worryFunc := func(worryLevel int) int {
		return worryLevel / 3
	}

	return monkeyBusiness(monkeys, 20, worryFunc)
}

func PartTwo(filename string) int {
	monkeys := parse(filename)

	mod := func() int {
		product := 1
		for _, monkey := range monkeys {
			product *= monkey.divisiblyBy
		}
		return product
	}()

	worryFunc := func(worryLevel int) int {
		return worryLevel % mod
	}

	return monkeyBusiness(monkeys, 10000, worryFunc)
}

func monkeyBusiness(monkeys []monkey, rounds int, worryFunc func(int) int) int {
	inspections := make(map[int]int, 0)
	for round := 0; round < rounds; round++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[monkey.id]++
				worry := monkey.operation(item)
				nextWorry := worryFunc(worry)
				dest := monkey.destination(nextWorry)
				monkeys[dest].items = append(monkeys[dest].items, nextWorry)
			}

			monkeys[i].items = make([]int, 0)
		}
	}

	keys := make([]int, 0)
	for k := range inspections {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return inspections[keys[i]] > inspections[keys[j]]
	})

	return inspections[keys[0]] * inspections[keys[1]]
}

func parse(filename string) []monkey {
	monkeys := make([]monkey, 0)
	lines := parser.ReadFile(filename)
	lineId := 0
	m := monkey{}
	for _, line := range lines {
		if line == "" {
			lineId = 0
			m = monkey{}
			continue
		}

		switch lineId {
		case 0:
			replaced := strings.Replace(line, ":", "", -1)
			fields := strings.Fields(replaced)
			m.id, _ = strconv.Atoi(fields[1])
		case 1:
			replaced := strings.Replace(line, ",", "", -1)
			fields := strings.Fields(replaced)

			items := make([]int, 0)
			for _, field := range fields {
				if val, err := strconv.Atoi(field); err == nil {
					items = append(items, val)
				}
			}
			m.items = items
		case 2:
			split := strings.Split(line, "new = old ")
			operation := split[1]
			fields := strings.Fields(operation)

			if fields[1] == "old" {
				m.operation = func(worryLevel int) int {
					return worryLevel * worryLevel
				}
			} else {
				val, _ := strconv.Atoi(fields[1])
				if fields[0] == "+" {
					m.operation = func(worryLevel int) int {
						return worryLevel + val
					}
				} else {
					m.operation = func(worryLevel int) int {
						return worryLevel * val
					}
				}
			}
		case 3:
			fields := strings.Fields(line)
			val, _ := strconv.Atoi(fields[len(fields)-1])
			m.divisiblyBy = val
		case 4:
			fields := strings.Fields(line)
			m.truthyMonkeyId, _ = strconv.Atoi(fields[len(fields)-1])
		case 5:
			fields := strings.Fields(line)
			m.falsyMonkeyId, _ = strconv.Atoi(fields[len(fields)-1])
			monkeys = append(monkeys, m)
		}
		lineId++
	}

	return monkeys
}
