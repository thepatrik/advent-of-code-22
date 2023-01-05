package six

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	answer := PartOne("./data")
	if answer != 1848 {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo("./data")
	if answer != 2308 {
		t.Fail()
	}
}
