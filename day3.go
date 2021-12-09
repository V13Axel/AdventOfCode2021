package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day3(input []string) (int, int) {
	return PowerConsumption(input), OxygenGeneratorRating(input)
}

func OxygenGeneratorRating(input []string) int {
	filter := int64(math.Pow(2, float64(len(input[0])-1)))
	oxygenRating, err := strconv.ParseInt(strconv.Itoa(OxygenRecursion(input, true, filter)), 2, 16)
	check(err)

	CO2ScrubberRating, err := strconv.ParseInt(strconv.Itoa(OxygenRecursion(input, false, filter)), 2, 16)
	check(err)
	fmt.Println("Oxygen rating:", oxygenRating, "CO2 Scrubber:", CO2ScrubberRating)

	return int(oxygenRating) * int(CO2ScrubberRating)
}

func OxygenRecursion(input []string, findMostCommon bool, filter int64) int {
	fmt.Println(strconv.FormatInt(filter, 2))

	var bitMatches, bitMisses []string
	for _, row := range input {
		rowInt, err := strconv.ParseInt(row, 2, 16)
		check(err)

		if (rowInt & filter) == filter {
			bitMatches = append(bitMatches, row)
			continue
		}

		bitMisses = append(bitMisses, row)
	}

	fmt.Println("Matches", bitMatches, "Misses", bitMisses)

	var mostCommon, leastCommon []string
	if len(bitMatches) >= len(bitMisses) {
		mostCommon = bitMatches
		leastCommon = bitMisses
	} else {
		mostCommon = bitMisses
		leastCommon = bitMatches
	}

	if filter > 1 && ((len(mostCommon) > 1 && findMostCommon) || (len(leastCommon) > 1 && !findMostCommon)) {
		newFilter := filter / 2

		if findMostCommon {
			if len(mostCommon) == len(leastCommon) {
				return OxygenRecursion(mostCommon, findMostCommon, newFilter)
			}

			if len(mostCommon) > len(leastCommon) {
				return OxygenRecursion(mostCommon, findMostCommon, newFilter)
			}

			return OxygenRecursion(leastCommon, findMostCommon, newFilter)
		}

		if len(mostCommon) >= len(leastCommon) {
			return OxygenRecursion(leastCommon, findMostCommon, newFilter)
		}

		return OxygenRecursion(mostCommon, findMostCommon, newFilter)
	}

	if findMostCommon {
		result, err := strconv.Atoi(mostCommon[0])
		check(err)

		return result
	}

	result, err := strconv.Atoi(leastCommon[0])
	check(err)

	return result
}

func PowerConsumption(input []string) int {
	// var gammaRate int
	totalValues := len(input)
	rotatedInput := make([][]int, len(input[0]))
	for _, row := range input {
		for j, col := range strings.Split(row, "") {
			intval, err := strconv.Atoi(col)
			rotatedInput[j] = append(rotatedInput[j], intval)
			check(err)
		}
		// fmt.Println(rotatedInput)
	}

	var gammaBits, epsilonBits []int
	for _, col := range rotatedInput {
		oneCount := 0
		for _, row := range col {
			oneCount += row
		}

		if oneCount > (totalValues / 2) {
			gammaBits = append(gammaBits, 1)
			epsilonBits = append(epsilonBits, 0)
		} else {
			gammaBits = append(gammaBits, 0)
			epsilonBits = append(epsilonBits, 1)
		}
	}

	var gammaString, epsilonString string
	for index, digit := range gammaBits {
		gammaString += strconv.Itoa(digit)
		epsilonString += strconv.Itoa(epsilonBits[index])
	}

	gamma, err := strconv.ParseInt(gammaString, 2, 16)
	check(err)
	epsilon, err := strconv.ParseInt(epsilonString, 2, 16)
	check(err)

	return int(gamma * epsilon)
}
