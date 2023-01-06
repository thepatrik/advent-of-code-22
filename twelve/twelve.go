package twelve

import (
	"errors"
	"math"

	localmath "github.com/thepatrik/advent-of-code-22/pkg/math"
	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type Pos struct {
	x int
	y int
}

type heightmap [][]rune

var directions = []Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

var errNotFound = errors.New("not found")

const (
	endRune     = 'E'
	startRune   = 'S'
	lowestRune  = 'a'
	highestRune = 'z'
)

func PartOne(filename string) int {
	hmap := parse(filename)
	starts, end := findStartsAndEnd(hmap, false)

	return minSteps(hmap, starts, end)
}

func PartTwo(filename string) int {
	hmap := parse(filename)
	starts, end := findStartsAndEnd(hmap, true)

	return minSteps(hmap, starts, end)
}

func findStartsAndEnd(hmap heightmap, includeLowestRune bool) ([]Pos, Pos) {
	starts := []Pos{}
	end := Pos{}

	for i := range hmap {
		for j := range hmap[i] {
			switch hmap[i][j] {
			case startRune:
				starts = append(starts, Pos{i, j})
			case 'a':
				if includeLowestRune {
					starts = append(starts, Pos{i, j})
				}
			case endRune:
				end = Pos{i, j}
			}
		}
	}

	return starts, end
}

func minSteps(hmap heightmap, starts []Pos, end Pos) int {
	min := math.MaxInt

	hmap[end.x][end.y] = highestRune
	for _, start := range starts {
		hmap[start.x][start.y] = lowestRune
		steps, err := traverse(hmap, start, end)
		if err != nil {
			continue
		}

		min = localmath.Min(min, steps)
	}

	return min
}

func parse(filename string) heightmap {
	strslice := parser.ReadFile(filename)
	hmap := make(heightmap, 0)

	for _, line := range strslice {
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		hmap = append(hmap, row)
	}

	return hmap
}

func traverse(hmap heightmap, start, end Pos) (int, error) {
	visits := make(map[Pos]bool)
	visits[start] = true
	queue := []Pos{start}
	steps := 0
	found := false

outer:
	for len(queue) != 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			pos := queue[0]
			queue = queue[1:]

			if pos == end {
				found = true
				break outer
			}

			for _, dir := range directions {
				nxtPos := Pos{pos.x + dir.x, pos.y + dir.y}

				if visits[nxtPos] {
					continue
				}

				if nxtPos.x < 0 || nxtPos.x >= len(hmap) || nxtPos.y < 0 || nxtPos.y >= len(hmap[0]) {
					continue
				}

				if hmap[nxtPos.x][nxtPos.y] > hmap[pos.x][pos.y] && hmap[nxtPos.x][nxtPos.y]-hmap[pos.x][pos.y] > 1 {
					continue
				}

				visits[nxtPos] = true
				queue = append(queue, nxtPos)
			}
		}

		steps++
	}

	if !found {
		return 0, errNotFound
	}

	return steps, nil
}
