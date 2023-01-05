package ten

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	answer := PartOne("./data")
	if answer != 14920 {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	answer := PartTwo("./data")
	fmt.Println(answer)
}
