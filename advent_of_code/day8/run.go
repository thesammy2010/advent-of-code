package main

import (
	"advent-of-code/advent_of_code/day8/part1"
	"advent-of-code/advent_of_code/day8/part2"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1")
	c := part1.Main("advent_of_code/day8/input.txt")
	fmt.Println(c)

	fmt.Println("Part 2")
	d := part2.Main("advent_of_code/day8/input.txt")
	fmt.Println(d)

	os.Exit(0)

}
