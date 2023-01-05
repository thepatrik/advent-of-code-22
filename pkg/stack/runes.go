package stack

import (
	"fmt"
)

var ErrEmptyStack = fmt.Errorf("empty stack")

type RuneStack []rune

func New() RuneStack {
	return make([]rune, 0)
}

func (s RuneStack) Push(v rune) RuneStack {
	return append(s, v)
}

func (s RuneStack) Pop() (RuneStack, rune, error) {
	l := len(s)
	if l == 0 {
		return s, 0, ErrEmptyStack
	}
	return s[:l-1], s[l-1], nil
}

func (s RuneStack) Empty() bool {
	return len(s) == 0
}
