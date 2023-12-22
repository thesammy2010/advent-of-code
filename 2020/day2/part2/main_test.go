package part2

import (
	"advent-of-code/advent_of_code/day2/part1"
	"advent-of-code/advent_of_code/day2/part2"
	"testing"
)

func TestIsPolicyValid(t *testing.T) {
	p1 := part1.Policy{1, 4, "a", "abcde"}
	if true != part2.IsPolicyValid(p1) {
		t.Errorf("Mismatch! Expected %t, got %t", true, part2.IsPolicyValid(p1))
	}
	p2 := part1.Policy{1, 4, "b", "cdefg"}
	if false != part2.IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, part2.IsPolicyValid(p2))
	}
	p3 := part1.Policy{1, 4, "c", "ccccccccc"}
	if false != part2.IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, part2.IsPolicyValid(p3))
	}
}

func TestMain(t *testing.T) {
	input_data := part2.Main("../input_test.txt")
	if input_data != 1 {
		t.Errorf("Mismatch! Expected %d, got %d", 2, input_data)
	}
}
