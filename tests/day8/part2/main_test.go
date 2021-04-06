package part2

import (
	"advent-of-code/advent_of_code/day8/part2"
	"advent-of-code/advent_of_code/helpers"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	rawResult, _ := helpers.ReadFileLineByLine("../input_test.txt")
	expectedResult := []string{"jmp +0", "acc +1", "jmp +4", "acc +3", "jmp -3", "acc -99", "acc +1", "jmp -4", "acc +6"}
	assert.Equal(t, part2.Change(rawResult, 0), expectedResult)
}

func TestMain(t *testing.T) {
	assert.Equal(t, part2.Main("../input_test.txt"), 8)
}
