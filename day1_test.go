package main

import (
	"strings"
	"testing"
)

func TestDayOnePartOneExample(t *testing.T) {
	input := `199
	200
	208
	210
	200
	207
	240
	269
	260
	263`
	wants1 := 7
	wants2 := 5
	result1, result2 := Day1(strings.Split(string(input), "\n"))

	if wants1 != result1 {
		t.Fatalf(`Day1() result 1 with example input wants %v, received %v instead.`, wants1, result1)
	}

	if wants2 != result2 {
		t.Fatalf(`Day1() result 2 with example input wants %v, received %v instead.`, wants2, result2)
	}
}
