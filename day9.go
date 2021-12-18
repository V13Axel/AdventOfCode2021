package main

import (
	"strconv"
	"strings"
)

func Day9(input []string) (int, int) {
	riskSum := 0
	for rowIndex, row := range input {
		columns := strings.Split(row, "")
		for colIndex, col := range columns {
			height, err := strconv.Atoi(col)
			check(err)
			if colIndex > 0 {
				if value, err := strconv.Atoi(string(input[rowIndex][colIndex-1])); value <= height {
					check(err)
					continue
				}
			}

			if colIndex+1 < len(row) {
				if value, err := strconv.Atoi(string(input[rowIndex][colIndex+1])); value <= height {
					check(err)
					continue
				}
			}

			if rowIndex > 0 {
				if value, err := strconv.Atoi(string(input[rowIndex-1][colIndex])); value <= height {
					check(err)
					continue
				}
			}

			if rowIndex+1 < len(input) {
				if value, err := strconv.Atoi(string(input[rowIndex+1][colIndex])); value <= height {
					check(err)
					continue
				}
			}

			// dump("Found: ", col)
			riskSum += height + 1
		}
	}

	return riskSum, 0
}
