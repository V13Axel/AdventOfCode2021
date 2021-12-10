package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntGrid struct {
	rows [][]int
}

func (i *IntGrid) initialize(width int, height int) {
	for y := 0; y <= height; y++ {
		var columns []int
		for x := 0; x <= width; x++ {
			columns = append(columns, 0)
		}
		i.rows = append(i.rows, columns)
	}
}

type Vent struct {
	x1    int
	y1    int
	x2    int
	y2    int
	drawn bool
}

func (v *Vent) slope() (xs int, ys int) {
	rise := v.y2 - v.y1
	run := v.x2 - v.x1
	ydir := 1
	xdir := 1
	// fmt.Println(rise, run)
	if rise == 0 {
		if run > 0 {
			return 0, 1
		}
		return 0, -1
	}
	if run == 0 {
		if rise > 0 {
			return 1, 0
		}
		return -1, 0
	}

	if run < 0 {
		xdir = -1
		run *= -1
	}
	if rise < 0 {
		ydir = -1
		run *= -1
	}

	gcd := GCD(run, rise)

	return (rise / gcd) * ydir, (run / gcd) * xdir
}

type VentGraph struct {
	vents   []Vent
	intgrid IntGrid
}

func (v *VentGraph) findLimits() (xl int, yl int) {
	for _, vent := range v.vents {
		if vent.x1 > xl {
			xl = vent.x1 + 1
		}
		if vent.x2 > xl {
			xl = vent.x2 + 1
		}
		if vent.y1 > yl {
			yl = vent.y2 + 1
		}
		if vent.y2 > yl {
			yl = vent.y2 + 1
		}
	}
	return
}

func (v *VentGraph) createCanvas() {
	v.intgrid.initialize(v.findLimits())
}

func (v *VentGraph) display() {
	for _, line := range v.intgrid.rows {
		fmt.Println(line)
	}
}

func (v *VentGraph) countIntersections() (intersections int) {
	for _, row := range v.intgrid.rows {
		for _, cell := range row {
			if cell > 1 {
				intersections++
			}
		}
	}

	return
}

func (v *VentGraph) draw(x int, y int) {
	add := 1
	// dump(x, y, add)
	if v.intgrid.rows[y][x] == -1 {
		add++
	}

	v.intgrid.rows[y][x] += add
	// v.display()
}

func (v *VentGraph) drawNext(ignoreDiagonals bool) {
	for ventIndex, vent := range v.vents {
		if vent.drawn {
			continue
		}

		ydir, xdir := vent.slope()

		if ignoreDiagonals && xdir != 0 && ydir != 0 {
			// dump("Skipping diagonal", xdir, ydir)
			continue
		}

		// dump("Drawing", ydir, xdir)

		x := vent.x1
		xtarget := vent.x2
		y := vent.y1
		ytarget := vent.y2

		dump(x, xtarget, y, ytarget)

		targetHit := false

		for !targetHit {
			if x == xtarget && y == ytarget {
				targetHit = true
			}
			v.draw(x, y)
			x += xdir
			y += ydir
			// dump("-----------")
		}

		v.vents[ventIndex].drawn = true
	}
}

func Day5(input []string) (int, int) {
	graph := VentGraph{
		vents: parseVents(input),
	}
	graph.createCanvas()
	// graph.display()
	graph.drawNext(true)
	part1 := graph.countIntersections()
	// graph.display()
	graph.drawNext(false)
	part2 := graph.countIntersections()
	return part1, part2
}

func parseVents(input []string) (vents []Vent) {
	for _, line := range input {
		coords := strings.Split(line, " -> ")
		pos1 := strings.Split(coords[0], ",")
		pos2 := strings.Split(coords[1], ",")
		x1 := intify(pos1[0])
		y1 := intify(pos1[1])
		x2 := intify(pos2[0])
		y2 := intify(pos2[1])

		vents = append(vents, Vent{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})
	}
	return
}

func intify(str string) int {
	res, err := strconv.Atoi(str)
	check(err)
	return int(res)
}
