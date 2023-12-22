package part1

import (
	"advent-of-code/advent_of_code/day1/part1"
	"testing"
)

func TestCheckSum(t *testing.T) {
	input1, input2 := 2019, 1
	expectedOutput := []interface{}{2019, 1, 2019, nil}
	calculatedOutput1, calculatedOutput2, calculatedOutput3, calculatedOutput4 := part1.CheckSum(input1, input2)
	calculatedOutput := []interface{}{calculatedOutput1, calculatedOutput2, calculatedOutput3, calculatedOutput4}

	for i := range calculatedOutput {
		if calculatedOutput[i] != expectedOutput[i] {
			t.Errorf("Output mismatch. Expected %s. Got %v", expectedOutput, calculatedOutput)
			break
		}
	}
}

func TestMakePairs(t *testing.T) {
	// function doesn't currently remove duplicates
	input := []int{1, 2, 3}
	expectedOutput := [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}}
	calculatedOutput := part1.MakePairs(input)

	for i, val := range calculatedOutput {
		for j := range val {
			if calculatedOutput[i][j] != expectedOutput[i][j] {
				t.Error("Output mismatch", "Expected: ", expectedOutput, "Calculated: ", calculatedOutput)
				break
			}
		}
	}
}

func TestMain(t *testing.T) {
	expectedOutput := []int{1721, 299, 514579}
	co1, co2, co3, calculatedErr := part1.Main("../input_test.txt")
	calculatedOutput := []int{co1, co2, co3}

	if calculatedErr != nil {
		t.Errorf("%s", calculatedErr)
	}
	for i := range calculatedOutput {
		if calculatedOutput[i] != expectedOutput[i] {
			t.Errorf("Output mismatch. Expected %v. Got %v", expectedOutput, calculatedOutput)
			break
		}
	}
}
