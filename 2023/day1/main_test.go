package main

import (
	"reflect"
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		"Example 1": {input: "two1nine", expected: 29},
		"Example 2": {input: "eightwothree", expected: 83},
		"Example 3": {input: "abcone2threexyz", expected: 13},
		"Example 4": {input: "xtwone3four", expected: 24},
		"Example 5": {input: "4nineeightseven2", expected: 42},
		"Example 6": {input: "zoneight234", expected: 14},
		"Example 7": {input: "7pqrstsixteen", expected: 76},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := convertV2(tc.input)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
