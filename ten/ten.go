package ten

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type operation string

const (
	noop operation = "noop"
	addx operation = "addx"
)

type cycle struct {
	register int
	op       operation
}

func PartOne(filename string) int {
	cycles := parse(filename)

	strength := 0
	for _, cycle := range []int{20, 60, 100, 140, 180, 220} {
		strength += cycle * cycles[cycle-1].register
	}

	return strength
}

func PartTwo(filename string) string {
	cycles := parse(filename)
	var img strings.Builder
	w, h, register, rest := 40, 6, cycles[0].register, 0

	withinStripe := func(pos, stripePos int) bool {
		return pos == stripePos-1 || pos == stripePos || pos == stripePos+1
	}

	pixel := func(pos, stripePos int) rune {
		if withinStripe(pos, stripePos) {
			return '#'
		}
		return '.'
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pxl := pixel(x, register)
			img.WriteRune(pxl)

			cycle := x + 1 + (y * w)

			if cycles[cycle].op == noop {
				rest = func(n int) int {
					if n == 0 {
						return 1
					}
					return 0
				}(rest)
			}

			if cycle%2 == rest {
				register = cycles[cycle].register
			}
		}

		img.WriteString("\n")
	}

	return img.String()
}

func parse(filename string) []cycle {
	x := 1
	cycles := make([]cycle, 0)
	cycles = append(cycles, cycle{register: x, op: noop})

	lines := parser.ReadFile(filename)
	for _, line := range lines {
		fields := strings.Fields(line)
		value := 0
		if len(fields) == 2 {
			value, _ = strconv.Atoi(fields[1])
		}

		operation := fields[0]
		switch operation {
		case "noop":
			cycles = append(cycles, cycle{register: x, op: noop})
		case "addx":
			for i := 0; i < 2; i++ {
				switch i {
				case 1:
					x += value
				}
				cycles = append(cycles, cycle{register: x, op: addx})
			}
		}
	}

	return cycles
}
