package thirteen

import (
	"encoding/json"
	"reflect"
	"sort"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type pkt []any

type pair struct {
	left  pkt
	right pkt
}

func PartOne(filename string) int {
	pkts := parse(filename)
	sum := 0

	for i, pkt := range pkts {
		if cmp(pkt.left, pkt.right) < 1 {
			sum += i + 1
		}
	}

	return sum
}

func PartTwo(filename string) int {
	parsed := parse(filename)

	pkts := [][]any{}
	for _, pkt := range parsed {
		pkts = append(pkts, pkt.left, pkt.right)
	}

	divOne, divTwo := newPkt("[[2]]"), newPkt("[[6]]")

	pkts = append(pkts, divOne, divTwo)

	sort.Slice(
		pkts,
		func(i, j int) bool {
			return cmp(pkts[i], pkts[j]) < 1
		},
	)

	divOneIx, divTwoIx := 0, 0
	for i, pkt := range pkts {
		if cmp(pkt, divOne) == 0 {
			divOneIx = i + 1
		} else if cmp(pkt, divTwo) == 0 {
			divTwoIx = i + 1
		}
	}

	return divOneIx * divTwoIx
}

func cmp(left []any, right []any) int {
	for i := 0; i < len(left) && i < len(right); i++ {
		nextLeft := left[i]
		nextRight := right[i]

		leftIsNum := reflect.TypeOf(nextLeft).Name() == "float64"
		rightIsNum := reflect.TypeOf(nextRight).Name() == "float64"

		if leftIsNum && rightIsNum {
			if nextLeft.(float64) < nextRight.(float64) {
				return -1
			}
			if nextLeft.(float64) > nextRight.(float64) {
				return 1
			}
		} else {
			var leftLeaves []any
			var rightLeaves []any

			if leftIsNum {
				leftLeaves = []any{nextLeft}
			} else {
				leftLeaves = nextLeft.([]any)
			}

			if rightIsNum {
				rightLeaves = []any{nextRight}
			} else {
				rightLeaves = nextRight.([]any)
			}

			res := cmp(leftLeaves, rightLeaves)

			if res != 0 {
				return res
			}
		}
	}

	if len(left) < len(right) {
		return -1
	}

	if len(left) > len(right) {
		return 1
	}

	return 0
}

func newPkt(s string) pkt {
	var div2 []any
	err := json.Unmarshal([]byte(s), &div2)
	if err != nil {
		panic(err)
	}

	return div2
}

func parse(filename string) []pair {
	pkts := make([]pair, 0)
	strslice := parser.ReadFile(filename)
	lineIx := 0
	p := pair{}
	for _, line := range strslice {
		if line == "" {
			lineIx = 0
			p = pair{}
			continue
		}

		if lineIx == 0 {
			p.left = newPkt(line)
		}

		if lineIx == 1 {
			p.right = newPkt(line)
			pkts = append(pkts, p)
		}

		lineIx++
	}

	return pkts
}
