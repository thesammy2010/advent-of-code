package part2

import (
	"advent-of-code/advent_of_code/helpers"
	"strings"
)

func compileGroupsV2(data string) [][]string {
	groups := make([][]string, 0)
	newData := strings.Split(data, "\n\n")
	for _, i := range newData {
		group := make([]string, 0)
		temp := strings.Split(i, "\n")
		for _, j := range temp {
			group = append(group, j)
		}
		groups = append(groups, group)
	}
	return groups
}

func groupData(group []string) map[string]int {
	set := make(map[string]int)
	set["count"] = len(group)
	for _, p := range group {
		letters := strings.Split(p, "")
		for _, c := range letters {
			set[c] += 1
		}
	}
	return set
}

func getCountFromGroup(set map[string]int) int {
	count := 0
	for _, v := range set {
		if v == set["count"] {
			count++
		}
	}
	return count - 1
}

func Main(filepath string) int {

	file, err := helpers.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	groups := compileGroupsV2(file)
	count := 0

	for _, g := range groups {
		set := groupData(g)
		count += getCountFromGroup(set)
	}

	return count

}
