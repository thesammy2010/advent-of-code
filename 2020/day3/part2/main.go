package part2

import (
	"advent-of-code/advent_of_code/day3/part1"
	"advent-of-code/advent_of_code/helpers"
	"fmt"
)

func Slope(rows []string, slide int, lineSkip int) int {

	count := 0
	index := 0
	lineNumber := 0

	// go's while loops are weird
	for lineNumber < len(rows) {

		row := rows[lineNumber]
		res := part1.Loop(len(row), index, row)
		if res {
			count++
		}
		index += slide
		lineNumber += lineSkip

	}

	return count

}

func Main(filepath string) int {
	inputFile, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		fmt.Println("%n", err)
	}

	conditions := [][]int{
		{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1},
	}
	var counts []int

	for _, list := range conditions {
		counts = append(counts, Slope(inputFile, list[1], list[0]))
	}

	// go doesn't have generics, much sad
	returnValue := 1
	for _, val := range counts {
		returnValue *= val
	}

	return returnValue
}
