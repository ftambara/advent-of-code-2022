package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbersList, err := readCaloriesFile("day1/input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt:", err)
	}
	var totals []int
	for _, numbers := range numbersList {
		totals = append(totals, sum(numbers))
	}
	fmt.Printf("Max total: %v\n", largest(totals))
}

func readCaloriesFile(filename string) ([][]int, error) {
	// TODO: Solve by manually finding lines
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	var (
		bundles [][]int
		numbers []int
	)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			bundles = append(bundles, numbers)
			numbers = make([]int, 0)
			continue
		}
		num, err := strconv.Atoi(ln)
		if err != nil {
			return bundles, err
		}
		numbers = append(numbers, num)
	}
	if err := scanner.Err(); err != nil {
		return bundles, err
	}

	return bundles, nil
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func largest(numbers []int) int {
	if len(numbers) == 0 {
		panic("Do not call me with an emtpy slice!")
	}
	index := 0
	for i := range numbers[1:] {
		if numbers[i] > numbers[index] {
			index = i
		}
	}
	return numbers[index]
}
