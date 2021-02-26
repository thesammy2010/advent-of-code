package part1

import (
	"advent-of-code/advent_of_code/day5/part1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert.Equal(t, 0, part1.Convert("L"), "The two numbers should be the same.")
	assert.Equal(t, 1, part1.Convert("R"), "The two numbers should be the same.")
	assert.Equal(t, 21, part1.Convert("LRLRLR"), "The two numbers should be the same.")
}

func TestCreatePassport(t *testing.T) {
	expectedResult := part1.Seat{44, 5, 357}
	calculatedResult := part1.CreateSeat("FBFBBFFRLR")
	assert.Equal(t, expectedResult.Row, calculatedResult.Row, "The two attributes should be the same.")
	assert.Equal(t, expectedResult.Column, calculatedResult.Column, "The two attributes should be the same.")
	assert.Equal(t, expectedResult.SeatID, calculatedResult.SeatID, "The two attributes should be the same.")
}

func TestMain(t *testing.T) {
	assert.Equal(t, part1.Main("../input_test.txt"), 820)
}
