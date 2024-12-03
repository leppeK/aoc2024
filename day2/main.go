package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input_sample.txt")
	fmt.Printf("sample_p1: %v\n", part1(input))
	fmt.Printf("sample_p3: %v\n", part2(input))

	input, _ = os.ReadFile("input.txt")
	fmt.Printf("p1: %v\n", part1(input))
	fmt.Printf("p2: %v\n", part2(input))
}

func Counts[S []E, E comparable](s S, e E) int {
	var n int
	for _, v := range s {
		if v == e {
			n++
		}
	}
	return n
}

func generateSlice(input []byte) [][]int {
	var result [][]int
	// voor elke regel in input
	for _, line := range strings.Split(string(input), "\n") {
		var lineSlice []int
		if line == "" {
			continue
		}
		items := strings.Split(line, " ")
		for _, item := range items {
			val, _ := strconv.Atoi(item)
			lineSlice = append(lineSlice, val)
		}
		result = append(result, lineSlice)
	}

	return result
}

func isSafe(report []int, damping bool) bool {
	var ascending, control, descending bool
	descending = (report[0] > report[1] && report[0]-report[1] <= 3)
	ascending = (report[1] > report[0] && report[1]-report[0] <= 3)
	control = ascending || descending
	if !control {
		return false
	}
	for i := 0; i < len(report)-1; i++ {

		if (report[i] < report[i+1] && report[i+1]-report[i] <= 3) && ascending {
			continue
		}
		if (report[i] > report[i+1] && report[i]-report[i+1] <= 3) && descending {
			continue
		}
		if damping {
			for ii := 0; ii < len(report); ii++ {
				xa := append(report[:ii], report[ii+1:]...)

				if isSafe(xa, false) {
					return true
				}
			}
		}
		return false
	}
	return true
}

func part1(input []byte) int {
	reports := generateSlice(input)
	var safe int
	for _, report := range reports {
		if isSafe(report, false) {
			safe++
		}
	}
	return safe
}

func part2(input []byte) int {
	reports := generateSlice(input)
	var safe int
	for _, report := range reports {
		if isSafe(report, true) {
			safe++
		}
	}
	return safe
}
