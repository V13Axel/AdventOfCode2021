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
	testDay("4", 4512, 0, t)
}
