package part2

import (
	"advent-of-code/advent_of_code/day8/part1"
	"advent-of-code/advent_of_code/helpers"
	"strings"
)

func Change(rs []string, indexToChange int) []string {
	switch rs[indexToChange][:3] {
	case "nop":
		rs[indexToChange] = strings.ReplaceAll(rs[indexToChange], "nop", "jmp")
	case "jmp":
		rs[indexToChange] = strings.ReplaceAll(rs[indexToChange], "jmp", "nop")
	}
	return rs
}

func Main(filepath string) int {

	rows, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		panic(err)
	}
	indexes := make(map[int]bool)
	ind := 0
	acc := 0

	// make map of changable keys
	changableIndexes := make([]int, 0)
	for k, v := range rows {
		if v[:3] == "nop" || v[:3] == "jmp" {
			changableIndexes = append(changableIndexes, k)
		}
	}
	nextIndex := -1
	// main run
	for {

		newIndex, newAcc := part1.Action(rows[ind])
		ind += newIndex

		// swap condition
		if _, ok := indexes[ind]; ok {

			nextIndex += 1
			orig_rows, _ := helpers.ReadFileLineByLine(filepath)
			rows = Change(orig_rows, changableIndexes[nextIndex])
			ind = 0
			acc = 0
			indexes = make(map[int]bool)
			continue
		}

		// terminated successfully
		if ind >= len(rows) {
			acc += newAcc
			break
		}

		indexes[ind] = true
		acc += newAcc

	}

	return acc
}
