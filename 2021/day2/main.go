package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// read file
	bytes, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")

	x := 0
	y := 0

	for _, line := range lines {
		if line == "" {
			break
		}
		vals := strings.Split(line, " ")
		direction := vals[0]
		step, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			x += step
		case "down":
			y += step
		case "up":
			y -= step
		}

	}

	fmt.Printf("Part 1: %d\n", x * y)

}

func part2() {
	// read file
	bytes, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bytes), "\n")

	x := 0
	y := 0
	z := 0

	for _, line := range lines {
		if line == "" {
			break
		}
		vals := strings.Split(line, " ")
		direction := vals[0]
		step, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			x += step
			z += step *y
		case "down":
			y += step
		case "up":
			y -= step
		}

	}

	fmt.Printf("Part 2: %d\n", x * z)
}

func main() {
	part1()
	part2()
}

