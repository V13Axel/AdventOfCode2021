package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type SevenSegment struct {
	digits   []string
	signal   []string
	segments map[string]int
}

func (s *SevenSegment) sortInputs(input []string, sortMembers bool) (result []string) {
	for _, digit := range input {
		result = append(result, sortChars(digit))
	}

	if sortMembers {
		sort.Slice(result, func(i, j int) bool {
			return len(result[i]) < len(result[j])
		})
	}

	return
}

func (s *SevenSegment) sort() {
	s.digits = s.sortInputs(s.digits, true)
	s.signal = s.sortInputs(s.signal, false)
}

func (s *SevenSegment) disambiguate() {
	s.segments[s.digits[0]] = 1
	s.segments[s.digits[2]] = 4
	s.segments[s.digits[1]] = 7
	s.segments[s.digits[9]] = 8

	// The digits with 5 segments are:
	// - 3: Found because it shares two segments with 1
	// - 2: Found because it shares two segments with 4
	// - 5: Found because it is the only remaining option
	for i := 3; i <= 5; i++ {
		if matchingChars(s.digits[i], s.digits[0]) == 2 {
			s.segments[s.digits[i]] = 3
			continue
		}

		if matchingChars(s.digits[i], s.digits[2]) == 2 {
			s.segments[s.digits[i]] = 2
			continue
		}

		s.segments[s.digits[i]] = 5
	}

	// The digits with 6 segments are:
	// - 6: Found because it shares one segment with 1
	// - 9: Found because it contains all 5 segments of 5
	// - 0: Found because it is the only remaining option
	for i := 6; i <= 8; i++ {
		if matchingChars(s.digits[i], getKey(s.segments, 1)) == 1 {
			s.segments[s.digits[i]] = 6
			continue
		}

		if matchingChars(s.digits[i], getKey(s.segments, 5)) == 5 {
			s.segments[s.digits[i]] = 9
			continue
		}

		s.segments[s.digits[i]] = 0
	}

	return
}

func (s *SevenSegment) convertSignal() (signal int) {
	var convertedSignal string
	for _, digit := range s.signal {
		convertedSignal += fmt.Sprint(s.segments[digit])
	}
	signal, err := strconv.Atoi(convertedSignal)
	check(err)
	return
}

func Day8(input []string) (int, int) {
	// Part 1 is simple. Let's hope I don't regret this for part 2.
	part1 := 0
	for _, pattern := range input {
		outputVals := strings.Split(strings.Split(pattern, " | ")[1], " ")
		for _, outputVal := range outputVals {
			length := len(outputVal)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				part1++
			}
		}
	}

	// Update: I only regretted it a little
	total := 0
	for _, line := range input {
		parts := strings.Split(line, " | ")
		display := SevenSegment{
			digits:   strings.Split(parts[0], " "),
			signal:   strings.Split(parts[1], " "),
			segments: make(map[string]int, 10),
		}

		display.sort()
		display.disambiguate()

		total += display.convertSignal()
	}

	return part1, total
}
