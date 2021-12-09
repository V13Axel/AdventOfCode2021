package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day4(input []string) (int, int) {
	calls := splitCalls(input[0])
	boards := splitBoards(input[1:])
	scorecards := makeScoreCards(len(boards))
	fmt.Println(calls, boards, scorecards)

	return 0, 0
}

func makeScoreCards(amount int) (cards [][]int) {
	for i := 0; i < amount; i++ {
		cards = append(cards, make([]int, 25))
	}
	return
}

func splitCalls(input string) (calls []string) {
	calls = strings.Split(input, ",")
	return
}

func splitBoards(input []string) (boards [][]int) {
	for i := 0; i < len(input); i += 6 {
		var ints []int
		for _, row := range input[i+1 : i+6] {
			fmt.Println(row)
			for _, numString := range strings.Split(strings.TrimSpace(row), " ") {
				if numString == "" {
					continue
				}
				fmt.Println(numString)
				val, err := strconv.Atoi(numString)
				check(err)
				ints = append(ints, int(val))
			}
		}

		if len(ints) != 25 {
			fmt.Println("Expected length 25, got " + strconv.Itoa(len(ints)) + "instead")
		}

		boards = append(boards, ints)
	}

	return
}
