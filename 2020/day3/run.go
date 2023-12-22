package main

import (
	"advent-of-code/advent_of_code/day3/part1"
	"advent-of-code/advent_of_code/day3/part2"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1")
	c := part1.Main("advent_of_code/day3/input.txt")
	fmt.Println(c)

	fmt.Println("Part 2")
	c = part2.Main("advent_of_code/day3/input.txt")
	fmt.Println(c)

	os.Exit(0)

}
