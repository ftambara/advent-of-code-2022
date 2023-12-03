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
	total := 0
	for i := range forest {
		row := forest[i]
		// Count left to right
		lr := genSequence(Coord{0, i}, moveRight, len(row))
		total += visible(forest, lr)
		// Count right to left
		rl := genSequence(Coord{len(row) - 1, i}, moveLeft, len(row))
		total += visible(forest, rl)
	}
	// Row indices are inverted due to reading order
	for j := range forest[0] {
		// Count top to bottom
		tb := genSequence(Coord{j, 0}, increaseRow, len(forest))
		total += visible(forest, tb)
		// Count bottom to top
		bt := genSequence(Coord{j, len(forest) - 1}, decreaseRow, len(forest))
		total += visible(forest, bt)
	}
	fmt.Printf("Total visible (wrong!): %d\n", total)
}

type Coord struct {
	col int
	row int
}

func moveRight(c Coord) Coord {
	return Coord{c.col + 1, c.row}
}

func moveLeft(c Coord) Coord {
	return Coord{c.col - 1, c.row}
}

func decreaseRow(c Coord) Coord {
	return Coord{c.col, c.row - 1}
}

func increaseRow(c Coord) Coord {
	return Coord{c.col, c.row + 1}
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
