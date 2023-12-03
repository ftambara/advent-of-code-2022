package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	forest, err := readForest("day8/example.txt")
	if err != nil {
		panic(err)
	}
	for _, row := range forest {
		fmt.Printf("%v: visible (left to right): %v\n", row, visible(row))
	}
}

func visible(row []int) int {
	highest := row[0]
	count := 1 // First one counts
	for _, tree := range row {
		if tree > highest {
			highest = tree
			count++
		}
	}
	return count
}

func readForest(filename string) ([][]int, error) {
	// Return the forest as a slice (rows) of slice (columns) of int
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	forest := [][]int{}
	row := 0
	for scanner.Scan() {
		ln := scanner.Text()
		forest = append(forest, make([]int, len(ln)))
		for i, c := range ln {
			forest[row][i] = int(c - '0')
		}
		row++
	}
	if err := scanner.Err(); err != nil {
		return forest, nil
	}
	return forest, nil
}
