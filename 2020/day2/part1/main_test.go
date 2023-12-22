package part1

import (
	"advent-of-code/advent_of_code/day2/part1"
	"testing"
)

func TestMakePolicy(t *testing.T) {
	input := "1-4 q: abcqrs"
	expected := part1.Policy{1, 4, "q", "abcqrs"}
	calculated := part1.MakePolicy(input)

	if expected != *calculated {
		t.Errorf("Struct mismatch! Expected: %+v, got: %+v", expected, calculated)
	}
}

func TestIsPolicyValid(t *testing.T) {
	p1 := part1.Policy{1, 4, "q", "q"}
	if true != part1.IsPolicyValid(p1) {
		t.Errorf("Mismatch! Expected %t, got %t", true, part1.IsPolicyValid(p1))
	}
	p2 := part1.Policy{1, 4, "q", ""}
	if false != part1.IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, part1.IsPolicyValid(p2))
	}
	p3 := part1.Policy{1, 4, "q", "qqqqq"}
	if false != part1.IsPolicyValid(p2) {
		t.Errorf("Mismatch! Expected %t, got %t", false, part1.IsPolicyValid(p3))
	}
}

func TestMain(t *testing.T) {
	input_data := part1.Main("../input_test.txt")
	if input_data != 2 {
		t.Errorf("Mismatch! Expected %d, got %d", 2, input_data)
	}
}
