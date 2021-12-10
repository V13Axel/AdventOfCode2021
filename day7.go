package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day7(input []string) (int, int) {
	numStrings := strings.Split(input[0], ",")
	// Convert to integer slice
	var numbers []int
	for _, numString := range numStrings {
		number, err := strconv.Atoi(numString)
		check(err)
		numbers = append(numbers, number)
	}

	// Sort slice
	sort.Ints(numbers)

	// Determine median value
	medianPos := int(math.Round(float64(len(numbers) / 2)))
	median := numbers[medianPos]

	var part1FuelUsage int
	for _, number := range numbers {
		// Calculate part 1 using median
		part1FuelUsage += int(math.Abs(float64(number - median)))
	}

	// Determine mean value
	total := 0.0
	for _, v := range numbers {
		total += float64(v)
	}

	mean := int(math.Round(float64(total / float64(len(numbers)))))

	// Calculate part 2 by guess-and-check
	var part2FuelUsage, previous int
	for i := mean - 3; i < (mean + 3); i++ {
		previous = part2FuelUsage
		part2FuelUsage = 0
		for _, number := range numbers {
			diff := number - i
			part2FuelUsage += triangularNumber(int(math.Abs(float64(diff))))
		}

		if previous > 0 && part2FuelUsage > previous {
			part2FuelUsage = previous
			break
		}
	}

	return part1FuelUsage, part2FuelUsage
}
