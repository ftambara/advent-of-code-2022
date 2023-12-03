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
	totalVisible := 0
	for _, row := range forest {
		// Count left to right
		count := visible(row, genSequence(0, len(row)))
		fmt.Printf("%v: visible (left to right): %v\n", row, count)
		totalVisible += count
		fmt.Println()
		// Count right to left
		count = visible(row, genSequence(len(row)-1, -1))
		fmt.Printf("%v: visible (right to left): %v\n", row, count)
		totalVisible += count
		fmt.Println()
	}
}

func genSequence(start, end int) []int {
	var step int
	switch {
	case start < end:
		step = 1
	case start > end:
		step = -1
	default:
		return nil
	}
	seq := make([]int, (end-start)/step)
	for i := 0; i < len(seq); i++ {
		seq[i] = start + i*step
	}
	return seq
}

func visible(row, sequence []int) int {
	highest := row[sequence[0]]
	count := 1 // First one counts
	for _, i := range sequence {
		if row[i] > highest {
			highest = row[i]
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
