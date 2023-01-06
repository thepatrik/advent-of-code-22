package thirteen

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	answer := PartOne("./data")
	if answer != 5659 {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo("./data")
	if answer != 22110 {
		t.Fail()
	}
}
