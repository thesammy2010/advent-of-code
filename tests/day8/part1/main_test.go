package part1

import (
	"advent-of-code/advent_of_code/day8/part1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAction(t *testing.T) {
	// will use a struct next time
	er11, er12 := 1, 0
	cr11, cr12 := part1.Action("nop +0")
	er21, er22 := -3, 0
	cr21, cr22 := part1.Action("jmp -3")
	er31, er32 := 1, 6
	cr31, cr32 := part1.Action("acc +6")

	assert.Equal(t, cr11, er11)
	assert.Equal(t, cr12, er12)
	assert.Equal(t, cr21, er21)
	assert.Equal(t, cr22, er22)
	assert.Equal(t, cr31, er31)
	assert.Equal(t, cr32, er32)
}

func TestMain(t *testing.T) {
	assert.Equal(t, part1.Main("../input_test.txt"), 5)
}
