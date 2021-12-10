package main

import (
	"strconv"
	"strings"
)

type School struct {
	fish []int
}

func (s *School) Populate(add int) {
	s.fish[add]++
}

func (s *School) getPopulation() (total int) {
	for _, count := range s.fish {
		total += count
	}
	return
}

func (s *School) tick(tick int) {
	zero := s.fish[0]

	for i := 0; i < 8; i++ {
		s.fish[i] = s.fish[i+1]
	}

	s.fish[6] += zero
	s.fish[8] = zero

	dump("After ", tick, "days: ", s.fish)
}

func Day6(input []string) (int, int) {
	school := parseFish(input[0])
	part1Days := 80
	part2Days := 256
	var part1 int
	for i := 1; i <= part2Days; i++ {
		school.tick(i)
		if i == part1Days {
			part1 = school.getPopulation()
		}
	}
	return part1, school.getPopulation()
}

func parseFish(input string) School {
	school := School{}
	for i := 0; i <= 8; i++ {
		school.fish = append(school.fish, 0)
	}

	fishNumbers := strings.Split(input, ",")
	for _, numString := range fishNumbers {
		days, err := strconv.Atoi(numString)
		check(err)
		school.Populate(days)
	}

	return school
}
