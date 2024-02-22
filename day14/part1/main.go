package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paths := ReadPaths("day14/input.txt")

	sandSource := Pos{500, 0}
	positions := []Pos{sandSource}
	for _, path := range paths {
		positions = append(positions, path...)
	}
	topLeft, bottomRight := GetCorners(positions)
	fmt.Printf("Top left: %v, Bottom right: %v\n", topLeft, bottomRight)

	cave := NewCave(topLeft, bottomRight)
	fmt.Println("Initial cave")
	cave.PrintCave()

	// Block the cells that the paths go through
	for _, path := range paths {
		cave.BlockFromPath(path)
	}
	fmt.Println("Blocked cells")
	cave.PrintCave()

	// Let the sand fall
	var sandPositions []Pos
	for fallsThrough := false; !fallsThrough; {
		// Start at the source of the sand
		sandPos := sandSource
		for {
			// Get the next position the sand will fall to
			var nextPos Pos
			nextPos, fallsThrough = cave.nextSandPos(sandPos)
			if fallsThrough {
				// If the sand falls through, stop
				break
			} else if sandPos == nextPos {
				// If the sand is at rest, add it to the list of sand positions
				// and continue with the next piece of sand
				sandPositions = append(sandPositions, sandPos)
				cave.BlockAt(sandPos)
				break
			} else {
				// If the sand is still falling, update the sand position
				sandPos = nextPos
			}
		}
	}
	fmt.Println("Final cave")
	cave.PrintCave(sandPositions...)
	fmt.Println("Total sand at rest: ", len(sandPositions))
}

func GetCorners(positions []Pos) (topLeft, bottomRight Pos) {
	minX, minY := positions[0].x, positions[0].y
	maxX, maxY := positions[0].x, positions[0].y
	for _, pos := range positions {
		if pos.x < minX {
			minX = pos.x
		}
		if pos.y < minY {
			minY = pos.y
		}
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
	}
	return Pos{minX, minY}, Pos{maxX, maxY}
}

func ReadPaths(filename string) [][]Pos {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var paths [][]Pos
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var path []Pos
		ln := scanner.Text()
		for _, pairStr := range strings.Split(ln, " -> ") {
			pair := strings.Split(pairStr, ",")
			x, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}
			path = append(path, Pos{x, y})
		}
		paths = append(paths, path)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return paths
}

type Cave struct {
	blocked              [][]bool
	topLeft, bottomRight Pos
}

func NewCave(topLeft, bottomRight Pos) *Cave {
	c := new(Cave)
	c.topLeft = topLeft
	c.bottomRight = bottomRight
	c.blocked = make([][]bool, bottomRight.x-topLeft.x+1)
	for i := range c.blocked {
		c.blocked[i] = make([]bool, bottomRight.y-topLeft.y+1)
	}
	return c
}

// BlockAt blocks the cell at the given position,
// subtracting the top left corner from the position
// to get the correct index in the blocked array
func (c *Cave) BlockAt(pos Pos) {
	c.blocked[pos.x-c.topLeft.x][pos.y-c.topLeft.y] = true
}

// IsBlocked returns true if the cell at the given position is blocked
// subtracting the top left corner from the position to get the correct index
func (c *Cave) IsBlocked(pos Pos) bool {
	return c.blocked[pos.x-c.topLeft.x][pos.y-c.topLeft.y]
}

func (c *Cave) BlockFromPath(path []Pos) {
	for i := 0; i < len(path)-1; i++ {
		start, end := path[i], path[i+1]
		// Get the direction of the path
		xDiff := end.x - start.x
		yDiff := end.y - start.y
		// Note: the y axis is inverted
		if yDiff < 0 {
			// If the path is going up
			for y := start.y; y >= end.y; y-- {
				c.BlockAt(Pos{start.x, y})
			}
		} else if yDiff > 0 {
			// If the path is going down
			for y := start.y; y <= end.y; y++ {
				c.BlockAt(Pos{start.x, y})
			}
		} else if xDiff < 0 {
			// If the path is going left
			for x := start.x; x >= end.x; x-- {
				c.BlockAt(Pos{x, start.y})
			}
		} else if xDiff > 0 {
			// If the path is going right
			for x := start.x; x <= end.x; x++ {
				c.BlockAt(Pos{x, start.y})
			}
		}
	}
}

func (c *Cave) nextSandPos(src Pos) (next Pos, fallsThrough bool) {
	if src.y == c.bottomRight.y {
		return Pos{src.x, src.y + 1}, true
	} else if !c.IsBlocked(Pos{src.x, src.y + 1}) {
		return Pos{src.x, src.y + 1}, false
	} else if src.x == c.topLeft.x {
		return Pos{src.x - 1, src.y + 1}, true
	} else if !c.IsBlocked(Pos{src.x - 1, src.y + 1}) {
		return Pos{src.x - 1, src.y + 1}, false
	} else if src.x == c.bottomRight.x {
		return Pos{src.x + 1, src.y + 1}, true
	} else if !c.IsBlocked(Pos{src.x + 1, src.y + 1}) {
		return Pos{src.x + 1, src.y + 1}, false
	} else {
		return src, false
	}
}

func (c *Cave) PrintCave(sandPositions ...Pos) {
	sandMap := make(map[Pos]bool)
	for _, pos := range sandPositions {
		sandMap[pos] = true
	}

	for y := c.topLeft.y; y <= c.bottomRight.y; y++ {
		for x := c.topLeft.x; x <= c.bottomRight.x; x++ {
			if c.blocked[x-c.topLeft.x][y-c.topLeft.y] {
				fmt.Print("#")
			} else if sandMap[Pos{x, y}] {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Pos struct {
	x, y int
}
