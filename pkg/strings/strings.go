package strings

import (
	"strings"
)

func Intersection(s1, s2 string) int {
	sum := 0
	for _, r := range s1 {
		if strings.ContainsRune(s2, r) {
			sum += 1
		}
	}
	return sum
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
