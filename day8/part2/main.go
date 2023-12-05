package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	forest, err := readForest("day8/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		width  = len(forest[0])
		left   = 0
		height = len(forest)
		top    = 0
	)

	// Initialise scores matrix
	scores := make([][]int, height)
	underlying := make([]int, width*height)
	for i := range scores {
		scores[i], underlying = underlying[:width], underlying[width:]
	}

	for y := range forest {
		for x := range forest[0] {
			treeCoord := Coord{col: x, row: y}
			toTop := genSequence(treeCoord, decreaseRow, y-top+1)
			toBottom := genSequence(treeCoord, increaseRow, height-y)
			toLeft := genSequence(treeCoord, moveLeft, x-left+1)
			toRight := genSequence(treeCoord, moveRight, width-x)

			score := sees(forest, toTop) *
				sees(forest, toBottom) *
				sees(forest, toLeft) *
				sees(forest, toRight)
			scores[y][x] = score
		}
	}
	best := Coord{0, 0}
	for row, scoreRow := range scores {
		for col, score := range scoreRow {
			if score > scores[best.row][best.col] {
				best = Coord{row: row, col: col}
			}
		}
	}
	fmt.Printf("Highest score: %v @ %+v\n", scores[best.row][best.col], best)
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

// sees counts how many trees of forest a tree sees from the start
// of seq in the direction of the remaining coords
func sees(forest [][]int, seq []Coord) int {
	if len(seq) == 0 {
		return 0
	}
	start := seq[0]
	tree := forest[start.row][start.col]
	count := 0
	for _, c := range seq[1:] {
		count++
		if forest[c.row][c.col] >= tree {
			break
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
