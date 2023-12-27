package main

import (
	"reflect"
	"testing"
)

func TestExtractSet(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected Set
	}{
		"Example 1": {input: "3 blue, 4 red", expected: Set{blue: 3, red: 4, green: 0}},
		"Example 2": {input: "1 red, 2 green", expected: Set{red: 1, green: 2, blue: 0}},
		"Example 3": {input: "6 blue; 2 green", expected: Set{blue: 6, green: 2, red: 0}},
		"Example 4": {input: "10 red", expected: Set{red: 10}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := extractSet(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestExtractGames(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected Game
	}{
		"Example 1": {
			input: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: Game{
				id: 1,
				sets: []Set{
					{blue: 3, red: 4, green: 0},
					{red: 1, green: 2, blue: 6},
					{blue: 0, green: 2, red: 0},
				},
			},
		},
		"Example 2": {
			input: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: Game{
				id: 2,
				sets: []Set{
					{blue: 1, green: 2},
					{green: 3, blue: 4, red: 1},
					{green: 1, blue: 1},
				},
			},
		},
		"Example 3": {
			input: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: Game{
				id: 3,
				sets: []Set{
					{green: 8, blue: 6, red: 20},
					{blue: 5, red: 4, green: 13},
					{green: 5, red: 1},
				},
			},
		},
		"Example 4": {
			input: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: Game{
				id: 4,
				sets: []Set{
					{green: 1, red: 3, blue: 6},
					{green: 3, red: 6},
					{green: 3, blue: 15, red: 14},
				},
			},
		},
		"Example 5": {
			input: "Game 50: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: Game{
				id: 50,
				sets: []Set{
					{red: 6, blue: 1, green: 3},
					{blue: 2, red: 1, green: 2},
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := extractGame(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestValidateGame(t *testing.T) {
	tests := map[string]struct {
		bag      Set
		game     Game
		expected bool
	}{
		"Example 1": {
			bag: Set{red: 12, green: 13, blue: 14},
			game: Game{
				id: 1,
				sets: []Set{
					{blue: 3, red: 4, green: 0},
					{red: 1, green: 2, blue: 6},
					{blue: 0, green: 2, red: 0},
				},
			},
			expected: true,
		},
		"Example 2": {
			bag: Set{red: 12, green: 13, blue: 14},
			game: Game{
				id: 2,
				sets: []Set{
					{blue: 1, green: 2},
					{green: 3, blue: 4, red: 1},
					{green: 1, blue: 1},
				},
			},
			expected: true,
		},
		"Example 3": {
			bag: Set{red: 12, green: 13, blue: 14},
			game: Game{
				id: 3,
				sets: []Set{
					{green: 8, blue: 6, red: 20},
					{blue: 5, red: 4, green: 13},
					{green: 5, red: 1},
				},
			},
			expected: false,
		},
		"Example 4": {
			bag: Set{red: 12, green: 13, blue: 14},
			game: Game{
				id: 4,
				sets: []Set{
					{green: 1, red: 3, blue: 6},
					{green: 3, red: 6},
					{green: 3, blue: 15, red: 14},
				},
			},
			expected: false,
		},
		"Example 5": {
			bag: Set{red: 12, green: 13, blue: 14},
			game: Game{
				id: 50,
				sets: []Set{
					{red: 6, blue: 1, green: 3},
					{blue: 2, red: 1, green: 2},
				},
			},
			expected: true,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := validateGame(tc.bag, tc.game)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
