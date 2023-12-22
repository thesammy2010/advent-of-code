package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func CheckSum(x int, y int) (int, int, int, error) {
	if x+y == 2020 {
		return x, y, x * y, nil
	}
	return 0, 0, 0, errors.New("x & y do not add up to 2020")
}

func CastToInt(x string) int {
	y, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	return int(y)
}

func MakePairs(data []int) [][]int {
	list := make([][]int, 0)
	// i and j are the pairs
	for _, i := range data {
		for _, j := range data {
			if i != j {
				tempList := []int{i, j}
				list = append(list, tempList)
			}
		}
	}
	return list
}

func Main(filepath string) (int, int, int, error) {
	inputData, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputDataIntegers := []int{}
	for _, row := range inputData {
		inputDataIntegers = append(inputDataIntegers, CastToInt(row))
	}
	numberPairs := MakePairs(inputDataIntegers)
	for _, row := range numberPairs {
		a, b, c, err := CheckSum(row[0], row[1])
		if err == nil {
			return a, b, c, nil
		}
	}
	return 0, 0, 0, errors.New("No pairs add up to 2020")
}
