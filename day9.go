package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HeightMap struct {
	grid         [][]int
	basins       map[int][]string
	basinTotals  map[int]int
	known        map[string]int // Grid indexes
	walls        map[string]bool
	currentBasin int
}

func (h *HeightMap) discoverBasins() {
	for rowIndex, row := range h.grid {
		for colIndex, height := range row {
			token := fmt.Sprint(rowIndex) + "," + fmt.Sprint(colIndex)
			if _, ok := h.known[token]; ok {
				dump(token+":", "Known item "+token+" present in basin", h.known[token])
				continue
			}

			dump(token+":", token+" not found yet")

			if height < 9 {
				// We are in an unknown basin... somewhere.
				dump(token+":", "Found ", height, " in basin ", h.currentBasin)
				h.known[token] = h.currentBasin
				h.searchWithinBasin(rowIndex, colIndex)
				h.currentBasin++

				continue
			}

			dump(token+":", token, " is a 9")
		}
	}
}

func (h *HeightMap) searchWithinBasin(x int, y int) bool {
	token := fmt.Sprint(x) + "," + fmt.Sprint(y)

	// Left, same, right
	for xdir := -1; xdir < 2; xdir++ {
		// Up, same, down
		for ydir := -1; ydir < 2; ydir++ {
			checkx := x + xdir
			checky := y + ydir
			// dump(checkx, checky)
			token = fmt.Sprint(checkx) + "," + fmt.Sprint(checky)
			h.debugGrid(checkx, checky)

			if (xdir != 0 && ydir != 0) || xdir+ydir == 0 {
				// dump(token+":", "Skipping stationary or diagonal")
				continue
			}

			if checkx < 0 || checkx >= len(h.grid) {
				dump(token+":", token+" would be off-map.")
				continue
			}

			if checky < 0 || checky >= len(h.grid[checkx]) {
				dump(token+":", token+" would be off-map.")
				continue
			}

			if _, ok := h.known[token]; ok {
				dump(token+":", "Known item "+token+" present in basin", h.known[token])
				continue
			}

			dump(token+":", "Proper check starting at "+token)

			for h.grid[checkx][checky] < 9 {
				token = fmt.Sprint(checkx) + "," + fmt.Sprint(checky)

				dump(token+":", "Found ", h.grid[checkx][checky], " in basin ", h.currentBasin)
				h.known[token] = h.currentBasin
				h.basinTotals[h.currentBasin]++

				dump(token+":", "Moving ", xdir, " X and ", ydir, " Y")
				checkx += xdir
				checky += ydir

				if checkx+xdir < 0 || checkx+xdir >= len(h.grid) {
					break
				}

				if checky+ydir < 0 || checky+ydir >= len(h.grid[checkx]) {
					break
				}
				h.debugGrid(checkx, checky)
			}
		}
	}

	return true
}

func (h *HeightMap) debugGrid(x int, y int) {
	fmt.Print("\n\n+")
	for i := 0; i < len(h.grid[0]); i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")
	for xIndex, row := range h.grid {
		fmt.Print("|")
		for yIndex, value := range row {
			if yIndex == y && xIndex == x {
				fmt.Print("X")
			} else {
				fmt.Print(value)
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("+")
	for i := 0; i < len(h.grid[0]); i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n\n")
}

func (h *HeightMap) makeIntGridFrom(input []string) (riskSum int) {
	for rowIndex, row := range input {
		columns := strings.Split(row, "")
		var rowInts []int
		for colIndex, col := range columns {
			height, err := strconv.Atoi(col)
			check(err)

			rowInts = append(rowInts, height)

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

		h.grid = append(h.grid, rowInts)
	}

	return
}

func Day9(input []string) (int, int) {
	heightmap := HeightMap{
		basins:      make(map[int][]string),
		basinTotals: make(map[int]int),
		known:       make(map[string]int),
		walls:       make(map[string]bool),
	}
	riskSum := heightmap.makeIntGridFrom(input)

	heightmap.discoverBasins()

	return riskSum, 0
}
