package main

import (
	"testing"
)

func TestDayOnePartOneExample(t *testing.T) {
	input := dayInput("day1", "/test_inputs/")
	wants1 := 7
	wants2 := 5
	result1, result2 := Day1(input)

	if wants1 != result1 {
		testFails(wants1, result1, "Day 1 part 1", t)
	}

	if wants2 != result2 {
		testFails(wants2, result2, "Day 1 part 2", t)
	}
}
