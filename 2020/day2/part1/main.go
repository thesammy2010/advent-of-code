package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Policy struct {
	LowerLimit int
	UpperLimit int
	Character  string
	Password   string
}

//  tbd
func MakePolicy(line string) *Policy {
	// 1-3 a: abcde
	p := Policy{}
	splitLine := strings.Split(line, " ")
	lowerLimit := strings.Split(splitLine[0], "-")[0]
	upperLimit := strings.Split(splitLine[0], "-")[1]
	character := strings.Split(splitLine[1], ":")[0]
	password := splitLine[2]

	aa, _ := strconv.ParseInt(lowerLimit, 10, 64)
	bb, _ := strconv.ParseInt(upperLimit, 10, 64)

	p.LowerLimit = int(aa)
	p.UpperLimit = int(bb)
	p.Character = character
	p.Password = password

	return &p
}

func IsPolicyValid(p Policy) bool {
	if p.LowerLimit <= strings.Count(p.Password, p.Character) && strings.Count(p.Password, p.Character) <= p.UpperLimit {
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
		policy := MakePolicy(line)
		if IsPolicyValid(*policy) {
			results++
		}
	}

	return results

}
