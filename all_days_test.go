package main

import (
	"testing"
)

func TestDayOne(t *testing.T) {
	testDay("1", 7, 5, t)
}

func TestDayTwo(t *testing.T) {
	testDay("2", 150, 900, t)
}

func TestDayThree(t *testing.T) {
	testDay("3", 198, 230, t)
}

func TestDayFour(t *testing.T) {
	testDay("4", 4512, 1924, t)
}

func TestDayFive(t *testing.T) {
	testDay("5", 5, 12, t)
}

func TestDaySix(t *testing.T) {
	testDay("6", 5934, 26984457539, t)
}

func TestDaySeven(t *testing.T) {
	testDay("7", 37, 168, t)
}

func TestDayEight(t *testing.T) {
	testDay("8", 26, 61229, t)
}
