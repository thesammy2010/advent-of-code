package part2

import (
	"advent-of-code/advent_of_code/day1/part1"
	"advent-of-code/advent_of_code/helpers"
	"errors"
	"fmt"
	"os"
)

func CheckSumV2(x int, y int, z int) (int, int, int, int, error) {
	if x+y+z == 2020 {
		return x, y, z, x * y * z, nil
	}
	return 0, 0, 0, 0, errors.New("x & y do not add up to 2020")
}

func MakeTriples(data []int) [][]int {
	list := make([][]int, 0)
	// i, j and k are the triples
	for _, i := range data {
		for _, j := range data {
			for _, k := range data {
				if i != j && i != k && j != k {
					tempList := []int{i, j, k}
					list = append(list, tempList)
				}
			}
		}
	}
	return list
}

func Main(filepath string) (int, int, int, int, error) {
	inputData, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputDataIntegers := []int{}
	for _, row := range inputData {
		inputDataIntegers = append(inputDataIntegers, part1.CastToInt(row))
	}
	numberPairs := MakeTriples(inputDataIntegers)
	for _, row := range numberPairs {
		a, b, c, d, err := CheckSumV2(row[0], row[1], row[2])
		if err == nil {
			return a, b, c, d, nil
		}
	}
	return 0, 0, 0, 0, errors.New("No pairs add up to 2020")

}
