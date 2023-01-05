package three

import (
	"github.com/thepatrik/advent-of-code-22/pkg/parser"
	"github.com/thepatrik/advent-of-code-22/pkg/slice"
)

type rucksack struct {
	one []int
	two []int
}

func PartOne(filename string) int {
	rucksacks := parse(filename)
	score := 0

	cmp := func(a, b []int) int {
		matches := make([]int, 0)

		for _, char := range a {
			if find(char, matches) {
				continue
			}
			for _, char2 := range b {
				if char == char2 {
					matches = append(matches, char)
					break
				}
			}
		}

		return slice.Sum(matches)
	}

	for _, rucksack := range rucksacks {
		score += cmp(rucksack.one, rucksack.two)
	}

	return score
}

func PartTwo(filename string) int {
	rucksacks := parse(filename)
	score := 0

	cmp := func(a, b, c []int) int {
		matches := make([]int, 0)
		for _, char := range a {
			if find(char, matches) {
				continue
			}

		out:
			for _, char2 := range b {
				for _, char3 := range c {
					if char == char2 && char == char3 {
						matches = append(matches, char)
						break out
					}
				}
			}
		}

		return slice.Sum(matches)
	}

	sacks := make([]rucksack, 0)

	for i, sack := range rucksacks {
		sacks = append(sacks, sack)
		if (i+1)%3 == 0 {
			one := append(sacks[0].one, sacks[0].two...)
			two := append(sacks[1].one, sacks[1].two...)
			three := append(sacks[2].one, sacks[2].two...)

			score += cmp(one, two, three)
			sacks = make([]rucksack, 0)
		}
	}

	return score
}

func find(r int, runes []int) bool {
	for _, v := range runes {
		if v == r {
			return true
		}
	}
	return false
}

func parse(filename string) []rucksack {
	strslice := parser.ReadFile(filename)

	aToZ := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	charToInt := func(char rune) int {
		for i, v := range aToZ {
			if v == char {
				return i + 1
			}
		}
		return 0
	}

	rucksacks := make([]rucksack, 0)

	for _, line := range strslice {
		intSlice := make([]int, 0)
		for _, char := range line {
			intSlice = append(intSlice, charToInt(char))
		}

		rucksack := rucksack{
			one: intSlice[:len(intSlice)/2],
			two: intSlice[len(intSlice)/2:],
		}

		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}
