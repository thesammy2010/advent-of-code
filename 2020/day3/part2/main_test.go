package part1

import (
	"advent-of-code/advent_of_code/day3/part2"
	"advent-of-code/advent_of_code/helpers"
	"testing"
)

func TestSlope(t *testing.T) {

	rows, err := helpers.ReadFileLineByLine("../input_test.txt")
	if err != nil {
		panic(err)
	}

	expectedResult1 := 2
	calculatedResult1 := part2.Slope(rows, 1, 1)
	if expectedResult1 != calculatedResult1 {
		t.Errorf("Mismatch! Expected %d, got %d", expectedResult1, calculatedResult1)
	}

	expectedResult2 := 7
	calculatedResult2 := part2.Slope(rows, 3, 1)
	if expectedResult2 != calculatedResult2 {
		t.Errorf("Mismatch! Expected %d, got %d", expectedResult2, calculatedResult2)
	}

	expectedResult3 := 3
	calculatedResult3 := part2.Slope(rows, 5, 1)
	if expectedResult3 != calculatedResult3 {
		t.Errorf("Mismatch! Expected %d, got %d", expectedResult3, calculatedResult3)
	}

	expectedResult4 := 4
	calculatedResult4 := part2.Slope(rows, 7, 1)
	if expectedResult4 != calculatedResult4 {
		t.Errorf("Mismatch! Expected %d, got %d", expectedResult4, calculatedResult4)
	}

	expectedResult5 := 2
	calculatedResult5 := part2.Slope(rows, 1, 2)
	if expectedResult5 != calculatedResult5 {
		t.Errorf("Mismatch! Expected %d, got %d", expectedResult5, calculatedResult5)
	}

}

func TestMain(t *testing.T) {
	input_data := part2.Main("../input_test.txt")
	if input_data != 336 {
		t.Errorf("Mismatch! Expected %d, got %d", 336, input_data)
	}
}
