package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var lookup = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}
var lookUpV2 = map[int]*regexp.Regexp{
	1: regexp.MustCompile("one"),
	2: regexp.MustCompile("two"),
	3: regexp.MustCompile("three"),
	4: regexp.MustCompile("four"),
	5: regexp.MustCompile("five"),
	6: regexp.MustCompile("six"),
	7: regexp.MustCompile("seven"),
	8: regexp.MustCompile("eight"),
	9: regexp.MustCompile("nine"),
}

type Character struct {
	position int
	value    int
}

func extractNumbers(line string) int {
	var numbers []int
	for _, char := range line {
		if number, err := strconv.Atoi(string(char)); err == nil {
			numbers = append(numbers, number)
		}
	}
	if cap(numbers) == 0 {
		return 0
	}
	numberPairs := fmt.Sprintf("%d%d", numbers[0], numbers[len(numbers)-1])
	number, err := strconv.Atoi(numberPairs)
	if err != nil {
		log.Fatalf("Couldn't process %s into : %v", numberPairs, err)
	}
	return number
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		number := extractNumbers(line)
		sum += number
	}
	return sum
}

func convertV2(input string) int {

	var positions []Character
	for number, pointer := range lookUpV2 {
		// check for numbers
		idx1 := strings.Index(input, fmt.Sprintf("%d", number))
		if idx1 > -1 {
			//positions[idx] = number
			positions = append(positions, Character{position: idx1, value: number})
		}

		idx2 := strings.LastIndex(input, fmt.Sprintf("%d", number))
		if idx2 > -1 {
			//positions[idx] = number
			positions = append(positions, Character{position: idx2, value: number})
		}

		// check for words
		results := pointer.FindAllStringSubmatchIndex(input, -1)
		for _, match := range results {
			//positions[match[0]] = number
			positions = append(positions, Character{position: match[0], value: number})

		}
	}

	// sort by key
	sort.Slice(positions, func(i, j int) bool {
		return positions[i].position < positions[j].position
	})

	// combine to numbers
	if cap(positions) == 0 {
		return 0
	}
	numberPairs := fmt.Sprintf("%d%d", positions[0].value, positions[len(positions)-1].value)
	number, err := strconv.Atoi(numberPairs)
	if err != nil {
		log.Fatalf("Couldn't process %s into : %v", numberPairs, err)
	}
	return number

}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		//convertedLine := convert(line)
		//number := extractNumbers(convertedLine)
		//fmt.Printf("Orginal line: `%s`, Converted line: `%s`, Number `%d`\n", line, convertedLine, number)

		number := convertV2(line)
		fmt.Printf("Orginal line: `%s`, Number `%d`\n", line, number)
		sum += number
	}
	return sum
}

func main() {
	fmt.Println("Day 1")
	filename := "./2023/day1/input.txt"
	if len(os.Args[1:]) > 1 {
		filename = os.Args[1]
	}
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	lines := strings.Split(string(body), "\n")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))

}
