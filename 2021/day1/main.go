package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	// read file
	bytes, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(bytes), "\n")

	count := 0
	prev := 999999
	for idx, number := range data {
		if number == "" {
			break
		}
		if idx == 0 {
			prev, err = strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			continue
		}
		n, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println(idx, number)
			panic(err)
		}
		if n > prev {
			count++
		}

		prev = n
	}

	fmt.Printf("Part 1: %d\n", count)

}

func part2() {
	// read file
	bytes, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(bytes), "\n")

	count := 0
	prev1 := 9999999
	prev2 := 9999999
	prev3 := 9999999
	for idx, n := range data {
		if n == "" {
			break
		}
		number, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		switch idx {
		case 0:
			prev3 = number
			continue
		case 1:
			prev2 = number
			continue
		case 2:
			prev1 = number
			continue
		}

		if (number + prev1 + prev2) > (prev1 + prev2 + prev3) {
			count++
		}
		prev1, prev2, prev3 = number, prev1, prev2
		fmt.Println(number, prev1, prev2, prev3)

	}

	fmt.Printf("Part 2: %d\n", count)

}

func main() {
	part1()
	part2()
}
