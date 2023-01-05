package eight

import (
	"strconv"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type trees [][]int

func (t trees) dimens() (int, int) {
	return len(t[0]), len(t)
}

func (t trees) String() string {
	s := ""
	for _, row := range t {
		for _, v := range row {
			s += strconv.Itoa(v)
		}
		s += "/n"
	}
	return s
}

func (t trees) scenicScoreMap() map[string]int {
	m := make(map[string]int)
	w, h := t.dimens()

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m[strconv.Itoa(x)+":"+strconv.Itoa(y)] = t.scenicScore(x, y)
		}
	}

	return m
}

func (t trees) sightMap() map[string]bool {
	m := make(map[string]bool)

	w, h := t.dimens()
	for y := 0; y < h; y++ {
		max := -1
		for x := 0; x < w; x++ {
			if t[y][x] > max {
				max = t[y][x]
				m[strconv.Itoa(x)+":"+strconv.Itoa(y)] = true
			}
		}

		max = -1
		for x := w - 1; x >= 0; x-- {
			if t[y][x] > max {
				max = t[y][x]
				m[strconv.Itoa(x)+":"+strconv.Itoa(y)] = true
			}
		}
	}

	for x := 0; x < w; x++ {
		max := -1
		for y := 0; y < h; y++ {
			if t[y][x] > max {
				max = t[y][x]
				m[strconv.Itoa(x)+":"+strconv.Itoa(y)] = true
			}
		}

		max = -1
		for y := h - 1; y >= 0; y-- {
			if t[y][x] > max {
				max = t[y][x]
				m[strconv.Itoa(x)+":"+strconv.Itoa(y)] = true
			}
		}
	}

	return m
}

func (t trees) scenicScore(treeX, treeY int) int {
	w, h := t.dimens()

	treeScore := t[treeY][treeX]

	right := 0

	y := treeY
	for x := treeX + 1; x < w; x++ {
		right++
		if treeScore <= t[y][x] {
			break
		}
	}

	left := 0
	for x := treeX - 1; x >= 0; x-- {
		left++
		if treeScore <= t[y][x] {
			break
		}
	}

	up := 0
	for y := treeY + 1; y < h; y++ {
		up++
		if treeScore <= t[y][treeX] {
			break
		}
	}

	down := 0
	for y := treeY - 1; y >= 0; y-- {
		down++
		if treeScore <= t[y][treeX] {
			break
		}
	}

	return right * left * up * down
}

func PartOne(filename string) int {
	trees := parse(filename)
	m := trees.sightMap()

	return len(m)
}

func PartTwo(filename string) int {
	trees := parse(filename)
	m := trees.scenicScoreMap()

	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}

func parse(filename string) trees {
	lines := parser.ReadFile(filename)
	rows := make([][]int, 0)
	for _, line := range lines {
		row := make([]int, 0)
		for _, c := range line {
			i, _ := strconv.Atoi(string(c))
			row = append(row, i)
		}
		rows = append(rows, row)
	}

	return rows
}
