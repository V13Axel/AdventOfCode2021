package main

import (
	"strconv"
	"strings"
)

func Day1(sweepLines []string) (int, int) {
	var sweepInts []int

	for _, x := range sweepLines {
		integer, err := strconv.Atoi(strings.TrimSpace(x))

		check(err)

		sweepInts = append(sweepInts, integer)
	}

	var second, third, timesIncremented, previousWindowSum, slidingIncrementCount int

	for index, first := range sweepInts {
		if index >= len(sweepInts)-2 {
			if third > second {
				timesIncremented++
			}
			break
		}

		second = sweepInts[index+1]
		third = sweepInts[index+2]

		if second > first {
			timesIncremented++
		}

		if index > 0 {
			if previousWindowSum < (first + second + third) {
				slidingIncrementCount++
			}
		}

		previousWindowSum = first + second + third
	}

	return timesIncremented, slidingIncrementCount
}
