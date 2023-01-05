package nine

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-22/pkg/math"
	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type move struct {
	direction string
	distance  int
}

type pos struct {
	x int
	y int
}

func (p1 pos) distance(p2 pos) int {
	xDist := math.Abs(p1.x - p2.x)
	yDist := math.Abs(p1.y - p2.y)
	return math.Max(xDist, yDist)
}

func moveTail(head pos, tail pos) pos {
	deltaX := head.x - tail.x
	deltaY := head.y - tail.y

	if math.Abs(deltaX) > 1 || math.Abs(deltaY) > 1 {
		return pos{tail.x + math.Sign(deltaX), tail.y + math.Sign(deltaY)}
	}

	return tail
}

func PartOne(filename string) int {
	moves := parse(filename)
	mPos := make(map[pos]int)
	head, tail := pos{
		x: 0,
		y: 0,
	}, pos{
		x: 0,
		y: 0,
	}

	mPos[tail] += 1

	for _, move := range moves {
		switch move.direction {
		case "U":
			for i := 0; i < move.distance; i++ {
				head.y += 1
				if head.distance(tail) > 1 {
					tail.y = head.y - 1
					tail.x = head.x
					mPos[tail] += 1
				}
			}
		case "D":
			for i := 0; i < move.distance; i++ {
				head.y -= 1
				if head.distance(tail) > 1 {
					tail.y = head.y + 1
					tail.x = head.x
					mPos[tail] += 1
				}
			}
		case "R":
			for i := 0; i < move.distance; i++ {
				head.x += 1
				if head.distance(tail) > 1 {
					tail.x = head.x - 1
					tail.y = head.y
					mPos[tail] += 1
				}
			}
		case "L":
			for i := 0; i < move.distance; i++ {
				head.x -= 1
				if head.distance(tail) > 1 {
					tail.x = head.x + 1
					tail.y = head.y
					mPos[tail] += 1
				}
			}
		}
	}
	return len(mPos)
}

func PartTwo(filename string) int {
	moves := parse(filename)
	mPos := make(map[pos]int)

	length := 10

	positions := make([]pos, 0)
	for i := 0; i < length; i++ {
		positions = append(positions, pos{x: 0, y: 0})
	}

	mPos[positions[9]] += 1

	for _, move := range moves {
		switch move.direction {
		case "U":
			for i := 0; i < move.distance; i++ {
				positions[0].y += 1

				for j := 1; j < length; j++ {
					positions[j] = moveTail(positions[j-1], positions[j])
					if j == length-1 {
						mPos[positions[j]] += 1
					}
				}
			}
		case "D":
			for i := 0; i < move.distance; i++ {
				positions[0].y -= 1
				for j := 1; j < length; j++ {
					if positions[j-1].distance(positions[j]) > 1 {
						positions[j] = moveTail(positions[j-1], positions[j])
					}
					if j == length-1 {
						mPos[positions[j]] += 1
					}
				}
			}
		case "R":
			for i := 0; i < move.distance; i++ {
				positions[0].x += 1
				for j := 1; j < length; j++ {
					if positions[j-1].distance(positions[j]) > 1 {
						positions[j] = moveTail(positions[j-1], positions[j])
					}
					if j == length-1 {
						mPos[positions[j]] += 1
					}
				}
			}
		case "L":
			for i := 0; i < move.distance; i++ {
				positions[0].x -= 1
				for j := 1; j < length; j++ {
					if positions[j-1].distance(positions[j]) > 1 {
						positions[j] = moveTail(positions[j-1], positions[j])
					}
					if j == length-1 {
						mPos[positions[j]] += 1
					}
				}
			}
		}
	}

	return len(mPos)
}

func parse(filename string) []move {
	moves := make([]move, 0)

	lines := parser.ReadFile(filename)
	for _, line := range lines {
		fields := strings.Fields(line)
		distance, _ := strconv.Atoi(fields[1])
		moves = append(moves, move{
			direction: fields[0],
			distance:  distance,
		})
	}

	return moves
}
