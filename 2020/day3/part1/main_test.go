package part1

import (
	"advent-of-code/advent_of_code/day3/part1"
	"testing"
)

func TestLoop(t *testing.T) {
	input_data1 := part1.Loop(11, 0, "..##.......")
	expected_result1 := false
	if input_data1 != expected_result1 {
		t.Errorf("Mismatch! Expected %t, got %t", input_data1, expected_result1)
	}

	input_data2 := part1.Loop(11, 3, "..##.......")
	expected_result2 := true
	if input_data2 != expected_result2 {
		t.Errorf("Mismatch! Expected %t, got %t", input_data2, expected_result2)
	}
}

func TestMain(t *testing.T) {
	input_data := part1.Main("../input_test.txt")
	if input_data != 7 {
		t.Errorf("Mismatch! Expected %d, got %d", 7, input_data)
	}
}
