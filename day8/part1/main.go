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
	for i, row := range forest {
		// Count left to right
		lr := genSequence(Coord{0, i}, moveRight, len(row))
		count := visible(forest, lr)
		fmt.Printf("%v: visible (left to right): %v\n", row, count)
	}
}

type Coord struct {
	col int
	row int
}

func moveRight(c Coord) Coord {
	return Coord{c.col + 1, c.row}
}

func genSequence(start Coord, next func(Coord) Coord, n int) []Coord {
	seq := []Coord{start}
	pos := start
	// First one is already in, do one less than n
	for i := 0; i < n-1; i++ {
		pos = next(pos)
		seq = append(seq, pos)
	}
	return seq
}

func visible(forest [][]int, sequence []Coord) int {
	firstC := sequence[0]
	highest := forest[firstC.row][firstC.col]
	count := 1 // First one counts
	for _, c := range sequence {
		tree := forest[c.row][c.col]
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
