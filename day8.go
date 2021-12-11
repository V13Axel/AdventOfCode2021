package main

import "strings"

func Day8(input []string) (int, int) {
	// Part 1 is simple. Let's hope I don't regret this for part 2.
	part1 := 0
	for _, pattern := range input {
		outputVals := strings.Split(strings.Split(pattern, " | ")[1], " ")
		for _, outputVal := range outputVals {
			length := len(outputVal)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				part1++
			}
		}
	}
	return part1, 0
}
