package two

import (
	"strings"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
)

type round struct {
	playerA gameElement
	playerB gameElement
}

type gameElement int64

const (
	paper gameElement = iota
	scissor
	rock
)

type outcome int64

const (
	win outcome = iota
	lose
	draw
)

type strategy struct {
	rock     outcome
	scissors outcome
	paper    outcome
}

var rockStrategy = strategy{
	rock:     draw,
	scissors: lose,
	paper:    win,
}

var paperStrategy = strategy{
	rock:     lose,
	scissors: win,
	paper:    draw,
}

var scissorsStrategy = strategy{
	rock:     win,
	scissors: draw,
	paper:    lose,
}

func PartOne(filename string) int {
	rounds := parse(filename)

	score := 0
	for _, round := range rounds {
		strategy := toStrategy(round.playerA)

		score += func() int {
			switch round.playerB {
			case rock:
				return outcomeScore(strategy.rock) + elementScore(rock)
			case paper:
				return outcomeScore(strategy.paper) + elementScore(paper)
			}
			return outcomeScore(strategy.scissors) + elementScore(scissor)
		}()
	}
	return score
}

func PartTwo(filename string) int {
	rounds := parse(filename)

	element := func(s strategy, o outcome) gameElement {
		if s.rock == o {
			return rock
		}
		if s.paper == o {
			return paper
		}
		return scissor
	}

	calculateScore := func(o outcome, s strategy) int {
		switch o {
		case win:
			return outcomeScore(win) + elementScore(element(s, win))
		case draw:
			return outcomeScore(draw) + elementScore(element(s, draw))
		}
		return outcomeScore(lose) + elementScore(element(s, lose))
	}

	expected := func(e gameElement) outcome {
		switch e {
		case rock:
			return lose
		case paper:
			return draw
		}
		return win
	}

	score := 0
	for _, round := range rounds {
		e := expected(round.playerB)
		s := toStrategy(round.playerA)

		score += calculateScore(e, s)
	}

	return score
}

func toStrategy(e gameElement) strategy {
	switch e {
	case scissor:
		return scissorsStrategy
	case paper:
		return paperStrategy
	}
	return rockStrategy
}

func outcomeScore(o outcome) int {
	switch o {
	case win:
		return 6
	case lose:
		return 0
	}
	return 3
}

func elementScore(e gameElement) int {
	switch e {
	case paper:
		return 2
	case scissor:
		return 3
	}
	return 1
}

func parse(filename string) []round {
	strslice := parser.ReadFile(filename)
	rounds := make([]round, 0)

	element := func(s string) gameElement {
		switch s {
		case "A", "X":
			return rock
		case "B", "Y":
			return paper
		case "C", "Z":
			return scissor
		}
		panic("not implemented")
	}

	for _, line := range strslice {
		split := strings.Split(line, " ")
		round := round{
			playerA: element(split[0]),
			playerB: element(split[1]),
		}
		rounds = append(rounds, round)
	}

	return rounds
}
