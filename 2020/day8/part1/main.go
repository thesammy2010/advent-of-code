package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"strconv"
	"strings"
)

func Action(row string) (int, int) {

	split := strings.Split(row, " ")
	act := split[0]
	val, _ := strconv.ParseInt(split[1], 10, 64)

	if act == "nop" {
		return 1, 0
	} else if act == "acc" {
		return 1, int(val)
	} else {
		return int(val), 0
	}

}

func Main(filepath string) int {

	rows, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		panic(err)
	}
	indexes := make(map[int]bool, 0)
	ind := 0
	acc := 0

	for true {
		newIndex, newAcc := Action(rows[ind])
		ind += newIndex

		if _, ok := indexes[ind]; ok {
			break
		}

		indexes[ind] = true
		acc += newAcc
	}

	return acc

}
