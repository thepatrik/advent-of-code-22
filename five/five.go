package five

import (
	"strconv"
	"strings"

	"github.com/thepatrik/advent-of-code-22/pkg/parser"
	"github.com/thepatrik/advent-of-code-22/pkg/stack"
	intstrings "github.com/thepatrik/advent-of-code-22/pkg/strings"
)

type instruction struct {
	count int
	from  int
	to    int
}

func PartOne(filename string) string {
	stacks := getStacks()
	instructions := parse(filename, 10)

	return move(stacks, instructions, false)
}

func PartTwo(filename string) string {
	stacks := getStacks()
	instructions := parse(filename, 10)

	return move(stacks, instructions, true)
}

func move(stacks []stack.RuneStack, instructions []instruction, retainOrder bool) string {
	for _, instruction := range instructions {
		count := instruction.count
		from := instruction.from
		to := instruction.to

		runesToPush := make([]rune, 0)
		for i := 0; i < count; i++ {
			var r rune
			stacks[from-1], r, _ = stacks[from-1].Pop()
			runesToPush = append(runesToPush, r)
		}

		if retainOrder {
			for i := len(runesToPush) - 1; i >= 0; i-- {
				stacks[to-1] = stacks[to-1].Push(runesToPush[i])
			}
		} else {
			for _, r := range runesToPush {
				stacks[to-1] = stacks[to-1].Push(r)
			}
		}
	}

	s := ""
	for _, stack := range stacks {
		if !stack.Empty() {
			var r rune
			_, r, _ = stack.Pop()
			s = string(r) + s
		}
	}

	return intstrings.Reverse(s)
}

func parse(filename string, skip int) []instruction {
	strslice := parser.ReadFile(filename)

	instructions := make([]instruction, 0)

	for i, line := range strslice {
		if i < skip {
			continue
		}

		splitted := strings.Split(line, " ")
		count, _ := strconv.Atoi(splitted[1])
		from, _ := strconv.Atoi(splitted[3])
		to, _ := strconv.Atoi(splitted[5])

		instruction := instruction{
			count: count,
			from:  from,
			to:    to,
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func getStacks() []stack.RuneStack {
	stacks := make([]stack.RuneStack, 0)
	stackOne := stack.New()
	stackOne = stackOne.Push('M')
	stackOne = stackOne.Push('J')
	stackOne = stackOne.Push('C')
	stackOne = stackOne.Push('B')
	stackOne = stackOne.Push('F')
	stackOne = stackOne.Push('R')
	stackOne = stackOne.Push('L')
	stackOne = stackOne.Push('H')
	stacks = append(stacks, stackOne)

	stackTwo := stack.New()
	stackTwo = stackTwo.Push('Z')
	stackTwo = stackTwo.Push('C')
	stackTwo = stackTwo.Push('D')
	stacks = append(stacks, stackTwo)

	stackThree := stack.New()
	stackThree = stackThree.Push('H')
	stackThree = stackThree.Push('J')
	stackThree = stackThree.Push('F')
	stackThree = stackThree.Push('C')
	stackThree = stackThree.Push('N')
	stackThree = stackThree.Push('G')
	stackThree = stackThree.Push('D')
	stacks = append(stacks, stackThree)

	stackFour := stack.New()
	stackFour = stackFour.Push('P')
	stackFour = stackFour.Push('J')
	stackFour = stackFour.Push('D')
	stackFour = stackFour.Push('M')
	stackFour = stackFour.Push('T')
	stackFour = stackFour.Push('S')
	stackFour = stackFour.Push('B')
	stacks = append(stacks, stackFour)

	five := stack.New()
	five = five.Push('N')
	five = five.Push('C')
	five = five.Push('D')
	five = five.Push('R')
	five = five.Push('J')
	stacks = append(stacks, five)

	six := stack.New()
	six = six.Push('W')
	six = six.Push('L')
	six = six.Push('D')
	six = six.Push('Q')
	six = six.Push('P')
	six = six.Push('J')
	six = six.Push('G')
	six = six.Push('Z')
	stacks = append(stacks, six)

	seven := stack.New()
	seven = seven.Push('P')
	seven = seven.Push('Z')
	seven = seven.Push('T')
	seven = seven.Push('F')
	seven = seven.Push('R')
	seven = seven.Push('H')
	stacks = append(stacks, seven)

	eight := stack.New()
	eight = eight.Push('L')
	eight = eight.Push('V')
	eight = eight.Push('M')
	eight = eight.Push('G')
	stacks = append(stacks, eight)

	nine := stack.New()
	nine = nine.Push('C')
	nine = nine.Push('B')
	nine = nine.Push('G')
	nine = nine.Push('P')
	nine = nine.Push('F')
	nine = nine.Push('Q')
	nine = nine.Push('R')
	nine = nine.Push('J')
	stacks = append(stacks, nine)

	return stacks
}
