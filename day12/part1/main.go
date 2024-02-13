package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	heightMap, start, end := readHeightMap("day12/input.txt")

	path := findPath(heightMap, start, end)
	fmt.Printf("Shortest path steps: %d\n", len(path)-1)
}

func decodeHeight(height byte) (h int, start bool, end bool) {
	switch height {
	case 'S':
		return 0, true, false
	case 'E':
		return 'z' - 'a', false, true
	default:
		return int(height - 'a'), false, false
	}
}

type Position struct {
	x, y int
}

func readHeightMap(filename string) (heightMap [][]int, start, end Position) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var row []int
		for _, height := range scanner.Text() {
			value, isStart, isEnd := decodeHeight(byte(height))
			row = append(row, value)
			if isStart {
				start = Position{len(row) - 1, len(heightMap)}
			} else if isEnd {
				end = Position{len(row) - 1, len(heightMap)}
			}
		}
		heightMap = append(heightMap, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return heightMap, start, end
}

func findPath(heightMap [][]int, start, end Position) []Position {
	shortestPaths := map[Position][]Position{
		start: {start},
	}
	positionsQueue := []Position{start}

	for i := 0; i < len(positionsQueue); i++ {
		position := positionsQueue[i]
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				// Skip self and diagonals
				if (dx == 0 && dy == 0) || (dx != 0 && dy != 0) {
					continue
				}
				x, y := position.x+dx, position.y+dy
				// Skip out of bounds
				if x < 0 || x >= len(heightMap[0]) || y < 0 || y >= len(heightMap) {
					continue
				}
				// Skip if the height difference is more than 1
				if heightMap[y][x]-heightMap[position.y][position.x] > 1 {
					continue
				}
				// Skip if a shorter path already exists
				if _, ok := shortestPaths[Position{x, y}]; ok {
					continue
				}
				shortestPaths[Position{x, y}] = append(shortestPaths[position], Position{x, y})

				// If we reached the destination, we can stop
				if x == end.x && y == end.y {
					return shortestPaths[end]
				}

				positionsQueue = append(positionsQueue, Position{x, y})
			}
		}
	}

	// If we reached this point, there is no path
	return nil
}
