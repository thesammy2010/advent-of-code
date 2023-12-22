package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"fmt"
)

func Loop(lengthOfRow int, index int, row string) bool {
	pos := row[index%lengthOfRow]

	if string(pos) == "#" {
		return true
	}
	return false
}

func Main(file string) int {

	inputFile, err := helpers.ReadFileLineByLine(file)
	if err != nil {
		fmt.Println("%n", err)
	}
	var count int
	var index int

	for _, row := range inputFile {
		res := Loop(len(row), index, row)
		if res {
			count++
		}
		index += 3
	}

	return count

}
