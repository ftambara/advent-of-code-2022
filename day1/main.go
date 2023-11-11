package main

import (
	"fmt"
	"slices"
)

func main() {
	numbersList, err := readCaloriesFile("day1/input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}
	var totals []int
	for _, numbers := range numbersList {
		totals = append(totals, sum(numbers))
	}

	// Part 1
	fmt.Printf("Max total: %v\n", largest(totals))

	// Part 2
	slices.SortFunc(totals, func(a, b int) int {
		switch {
		case a < b:
			return 1
		case a == b:
			return 0
		default:
			return -1
		}
	})
	n := 3
	topN := totals[:n]
	fmt.Printf("Top %v totals: %v\n", n, topN)

	fmt.Printf("These elves carry a total of %v calories\n", sum(topN))
}
