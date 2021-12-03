package main

import (
	"strconv"
	"strings"
)

func Day2(input []string) (int, int) {
	var position, aim, depth int
	for _, instruction := range input {
		parts := strings.Split(instruction, " ")
		action := parts[0]
		units, err := strconv.Atoi(parts[1])
		check(err)

		switch action {
		case "forward":
			position += units
			depth += (aim * units)
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}

	return (position * aim), (position * depth)
}
