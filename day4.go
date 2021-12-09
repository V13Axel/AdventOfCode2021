package main

import (
	"os"
	"strconv"
	"strings"
)

func Day4(input []string) (int, int) {
	calls := splitCalls(input[0])
	boards := splitBoards(input[1:])
	scorecards := makeScoreCards(len(boards))
	var firstWinner, lastWinner int
	winners := make([]int, len(boards))
	for i := 0; i < len(boards); i++ {
		winners[i] = -1
	}

	for _, call := range calls {
		for boardIndex, board := range boards {
			if winners[boardIndex] > 0 {
				continue
			}

			if contains, at := boardContains(board, call); contains {
				scorecards[boardIndex][at] = call
				boards[boardIndex][at] = -1
				if cardWins(scorecards[boardIndex]) {
					score := totalBoard(boards[boardIndex])
					winners[boardIndex] = score * call

					if firstWinner == 0 {
						firstWinner = winners[boardIndex]
					}

					lastWinner = boardIndex
				}
			}
		}
	}

	return firstWinner, winners[lastWinner]
}

func totalBoard(board []int) (total int) {
	for _, val := range board {
		if val > 0 {
			total += val
		}
	}

	return
}

func cardWins(card []int) bool {
	for row := 0; row < 5; row++ {
		winningRow := true
		for col := row * 5; col < (row*5)+5; col++ {
			if card[col] == -1 {
				winningRow = false
			}
		}
		if winningRow {
			return true
		}
	}

	for col := 0; col < 5; col++ {
		winningColumn := true
		row := 0
		for row = 0; row < 5; row++ {
			// fmt.Println(col, row, (row*5)+col)
			if card[(row*5)+col] == -1 {
				winningColumn = false
			}
		}
		if winningColumn {
			return true
		}
	}

	return false
}

func boardContains(board []int, value int) (contains bool, index int) {
	for index, space := range board {
		if space == value {
			return true, index
		}
	}

	return false, 0
}

func makeScoreCards(amount int) (cards [][]int) {
	for i := 0; i < amount; i++ {
		cards = append(cards, make([]int, 25))
		for j := 0; j <= 24; j++ {
			cards[i][j] = -1
		}
	}
	return
}

func splitCalls(input string) (calls []int) {
	callStrings := strings.Split(input, ",")
	for _, val := range callStrings {
		intVal, err := strconv.Atoi(val)
		check(err)
		calls = append(calls, intVal)
	}
	return calls
}

func splitBoards(input []string) (boards [][]int) {
	for i := 0; i < len(input); i += 6 {
		var ints []int
		for _, row := range input[i+1 : i+6] {
			for _, numString := range strings.Split(strings.TrimSpace(row), " ") {
				if numString == "" {
					continue
				}
				val, err := strconv.Atoi(numString)
				check(err)
				ints = append(ints, int(val))
			}
		}

		if len(ints) != 25 {
			os.Exit(1)
		}

		boards = append(boards, ints)
	}

	return
}
