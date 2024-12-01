package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input_sample.txt")
	fmt.Printf("Sample distance: %v\n", part1(input))
	fmt.printf("similarity score: %v\n", part2(input))

	input, _ = os.readfile("input.txt")
	fmt.printf("actual distance: %v\n", part1(input))
	fmt.printf("similarity score: %v\n", part2(input))
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

func generate_array(input []byte) ([]int, []int) {
	var l, r []int

	// voor elke regel in input
	for _, line := range strings.split(string(input), "\n") {
		if line == "" {
			continue
		}
		item := strings.split(line, "   ")

		// splits links van 3 spaties naar array l
func Counts[S []E, E comparable](s S, e E) int {
	var n int
	for _, v := range s {
		if v == e {
			n++
		}
	}
	return n
}
		itemL, _ := strconv.Atoi(item[0])

		// splits rechts van 3 spaties naar array R
		itemR, _ := strconv.Atoi(item[1])

		L = append(L, itemL)
		R = append(R, itemR)
	}

	return L, R
}

func part1(input []byte) int {
	L, R := generate_array(input)
	sort.Ints(L)
	sort.Ints(R)

	distance := 0
	for k := range L {
		if L[k] >= R[k] {
			distance += L[k] - R[k]
			continue
		}
		distance += R[k] - L[k]
	}
	return distance
}


func part2(input []byte) int {
	sim := 0
	L, R := generate_array(input)
	for _, v := range L {
		sim += v * Counts(R, v)
	}
	return sim
}
