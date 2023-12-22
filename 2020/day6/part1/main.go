package part1

import (
	"advent-of-code/advent_of_code/helpers"
	"strings"
)

type HashsetMap map[string]struct{}

func CompileGroups(line string) []string {
	groups := strings.Split(line, "\n\n")
	formattedGroup := make([]string, 0)
	for _, v := range groups {
		formattedGroup = append(formattedGroup, strings.Replace(v, "\n", "", -1))
	}
	return formattedGroup
}

func Main(filepath string) int {

	data, err := helpers.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	groups := CompileGroups(data)
	count := 0

	for _, i := range groups {
		set := make(map[string]bool)
		for _, j := range strings.Split(i, "") {
			set[j] = true
		}
		count += len(set)
	}

	return count

}
