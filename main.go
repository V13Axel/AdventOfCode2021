package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var StubStorage = map[string]interface{}{
	"Day1": Day1,
	"Day2": Day2,
	"Day3": Day3,
	"Day4": Day4,
	"Day5": Day5,
	"Day6": Day6,
	"Day7": Day7,
	"Day8": Day8,
	"Day9": Day9,
	// "Day10": Day10,
	// "Day11": Day11,
	// "Day13": Day13,
	// "Day14": Day14,
	// "Day15": Day15,
	// "Day16": Day16,
	// "Day17": Day17,
	// "Day18": Day18,
	// "Day19": Day19,
	// "Day20": Day20,
	// "Day21": Day21,
	// "Day22": Day22,
	// "Day23": Day23,
	// "Day24": Day24,
	// "Day25": Day25,
}

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		panic("You must supply a day name.")
	}

	if len(args[1:]) > 0 {
		Call(args[0], args[1:])
		return
	}

	res, err := Call(args[0], dayInput(strings.ToLower(args[0]), ""))
	check(err)
	if res == nil {
		fmt.Println("method call didn't return anything")
	}

	out("%v part 1 result: %d\n%v part 2 result: %d\n", args[0], res[0], args[0], res[1])
}

func dayInput(day string, subdir string) []string {
	ex, err := os.Getwd()
	check(err)

	if subdir == "" {
		subdir = "/inputs/"
	}

	daysPath := ex + subdir

	dat, err := os.ReadFile(daysPath + day + ".txt")
	check(err)

	return strings.Split(string(dat), "\n")
}

func Call(funcName string, params ...interface{}) (result []int, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is out of index")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	res := f.Call(in)
	result1 := int(res[0].Int())
	result2 := int(res[1].Int())

	result = append(result, result1)
	result = append(result, result2)
	return
}

func out(toPrint string, params ...interface{}) {
	fmt.Printf(toPrint+"\n", params...)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func testFails(wants int, result int, day string, t *testing.T) {
	t.Fatalf(`%v with example input wants %v, received %v instead.`, day, wants, result)
}

func testDay(day string, wants1 int, wants2 int, t *testing.T) {
	input := dayInput("day"+day, "/test_inputs/")
	res, err := Call("Day"+day, input)

	check(err)

	result1 := res[0]
	result2 := res[1]

	if wants1 != result1 {
		testFails(wants1, result1, "Day "+day+" part 1", t)
	}

	if wants2 > 0 && wants2 != result2 {
		testFails(wants2, result2, "Day "+day+" part 2", t)
	}
}

func BtoI(target bool) (result int) {
	result = 0

	if target {
		result = 1
	}

	return
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func triangularNumber(number int) (result int) {
	if number >= 0 {
		result = number + triangularNumber(number-1)
		return
	}
	return 0
}

func dd(anything ...interface{}) {
	fmt.Println(anything...)
	os.Exit(1)
}

func dump(anything ...interface{}) {
	fmt.Println(anything...)
}

func sortChars(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func matchingChars(needle string, haystack string) (found int) {
	for _, search := range needle {
		if strings.Contains(haystack, string(search)) {
			found++
		}
	}
	return
}

func getKey(m map[string]int, value int) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			return
		}
	}
	return
}
