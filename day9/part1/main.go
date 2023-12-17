package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var head, tail [2]int
	moves := ReadMoves("day9/input.txt")
	tailPositions := make(map[[2]int]bool, len(moves))
	// Include initial (0, 0) position
	tailPositions[tail] = true

	for _, move := range moves {
		for i := 0; i < move.times; i++ {
			head = MoveHead(head, move.direction)
			tail = MoveTail(head, tail)
			tailPositions[tail] = true
		}
	}
	fmt.Printf("Unique tail positions: %v\n", len(tailPositions))
}

type Move struct {
	direction byte
	times     int
}

const (
	Up    = 'U'
	Down  = 'D'
	Left  = 'L'
	Right = 'R'
)

func ReadMoves(filename string) []Move {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	moves := []Move{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var times int
		var direction byte
		ln := scanner.Text()
		fmt.Sscanf(ln, "%c %d\n", &direction, &times)
		moves = append(moves, Move{direction, times})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return moves
}

func MoveHead(head [2]int, dir byte) [2]int {
	x := head[0]
	y := head[1]
	switch dir {
	case Up:
		return [2]int{x, y + 1}
	case Down:
		return [2]int{x, y - 1}
	case Left:
		return [2]int{x - 1, y}
	case Right:
		return [2]int{x + 1, y}
	default:
		panic(fmt.Errorf("invalid direction: %c", dir))
	}
}

func MoveTail(head, tail [2]int) [2]int {
	dx := head[0] - tail[0]
	dy := head[1] - tail[1]
	absDx := int(math.Abs(float64(dx)))
	absDy := int(math.Abs(float64(dy)))
	if absDx <= 1 && absDy <= 1 {
		// Tail is adjacent to head
		return tail
	}
	return [2]int{tail[0] + Sign(dx), tail[1] + Sign(dy)}
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}
