package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	sets []Set
}

var patterns = map[string]*regexp.Regexp{
	"red":   regexp.MustCompile(`(\d{1,2}) red`),
	"green": regexp.MustCompile(`(\d{1,2}) green`),
	"blue":  regexp.MustCompile(`(\d{1,2}) blue`),
	"game":  regexp.MustCompile(`Game (\d{1,4}):`),
}

func extractSet(line string) Set {
	game := Set{}
	redString := patterns["red"].FindStringSubmatch(line)
	if cap(redString) > 1 {
		game.red, _ = strconv.Atoi(strings.Join(redString[1:], ""))
	}
	blueString := patterns["blue"].FindStringSubmatch(line)
	if cap(blueString) > 1 {
		game.blue, _ = strconv.Atoi(strings.Join(blueString[1:], ""))
	}
	greenString := patterns["green"].FindStringSubmatch(line)
	if cap(greenString) > 1 {
		game.green, _ = strconv.Atoi(strings.Join(greenString[1:], ""))
	}
	return game
}

func extractGame(line string) Game {
	games := Game{}
	games.id, _ = strconv.Atoi(patterns["game"].FindStringSubmatch(line)[1])
	for _, game := range strings.Split(patterns["game"].ReplaceAllString(line, ""), ";") {
		games.sets = append(games.sets, extractSet(game))
	}
	return games
}

func validateGame(bag Set, game Game) bool {
	validSets := 0
	totalSets := len(game.sets)
	for _, set := range game.sets {
		if set.red <= bag.red && set.green <= bag.green && set.blue <= bag.blue {
			validSets++
		}
	}
	return validSets == totalSets
}

func minimumBagRequired(game Game) Set {
	bag := Set{}
	for _, set := range game.sets {
		if set.red > bag.red {
			bag.red = set.red
		}
		if set.blue > bag.blue {
			bag.blue = set.blue
		}
		if set.green > bag.green {
			bag.green = set.green
		}
	}
	return bag
}

func part1(lines []string) int {
	sum := 0
	bag := Set{red: 12, green: 13, blue: 14}
	for _, line := range lines {
		game := extractGame(line)
		if validateGame(bag, game) {
			sum += game.id
		}
	}
	return sum
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		game := extractGame(line)
		minimumBag := minimumBagRequired(game)
		sum += minimumBag.red * minimumBag.green * minimumBag.blue
	}
	return sum
}

func main() {
	fmt.Println("Day 2")
	filename := "./2023/day2/input.txt"
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
