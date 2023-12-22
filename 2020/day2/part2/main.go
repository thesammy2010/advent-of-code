package part2

import (
	"advent-of-code/advent_of_code/day2/part1"
	"advent-of-code/advent_of_code/helpers"
	"fmt"
)

func IsPolicyValid(p part1.Policy) bool {
	i := []int{p.LowerLimit - 1, p.UpperLimit - 1} // indexes
	char1 := string(p.Password[i[0]])
	char2 := string(p.Password[i[1]])

	if char1 != p.Character && char2 == p.Character {
		return true
	}
	if char1 == p.Character && char2 != p.Character {
		return true
	}
	return false
}

func Main(filepath string) int {
	lines, err := helpers.ReadFileLineByLine(filepath)
	if err != nil {
		fmt.Println(err)
	}

	var results int

	for _, line := range lines {
		policy := part1.MakePolicy(line)
		if IsPolicyValid(*policy) {
			results++
		}
	}

	return results

}
