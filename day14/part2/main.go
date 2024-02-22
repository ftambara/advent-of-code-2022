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

	// Create the cave
	cave := NewCave(paths, Pos{500, 0}, 2)
	fmt.Println("Initial cave")
	cave.PrintCave()

	// Simulate sand fall
	sourceBlocked := false
	for !sourceBlocked {
		sourceBlocked = cave.SimulateSandFall()
	}

	// Count the sand pieces
	fmt.Printf("Sand pieces: %d\n", len(cave.sand))
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

type Cave struct {
	rock                 map[Pos]bool
	sand                 map[Pos]bool
	floorY               int
	topLeft, bottomRight Pos
	sandSource           Pos
	movingSand           Pos
}

func NewCave(paths [][]Pos, sandSource Pos, floorYIncr int) *Cave {
	positions := []Pos{sandSource}
	for _, path := range paths {
		positions = append(positions, path...)
	}
	topLeft, bottomRight := GetCorners(positions)

	floorY := bottomRight.y + floorYIncr
	if sandSource.y > floorY {
		panic("Sand source is below the floor")
	}
	// Update corners with floor
	bottomRight.y = floorY

	cave := &Cave{
		rock:        make(map[Pos]bool),
		sand:        make(map[Pos]bool),
		floorY:      floorY,
		topLeft:     topLeft,
		bottomRight: bottomRight,
		sandSource:  sandSource,
	}

	for _, path := range paths {
		cave.RocksFromPath(path)
	}

	return cave
}

func (c *Cave) AddRock(pos Pos) {
	if pos.y > c.floorY {
		panic("Position is below the floor")
	}
	// Update the corners of the cave
	if pos.x < c.topLeft.x {
		c.topLeft.x = pos.x
	} else if pos.x > c.bottomRight.x {
		c.bottomRight.x = pos.x
	}
	if pos.y < c.topLeft.y {
		c.topLeft.y = pos.y
	} else if pos.y > c.bottomRight.y {
		c.bottomRight.y = pos.y
	}

	// Add the rock
	c.rock[pos] = true
}

func (c *Cave) AddSand(pos Pos) {
	if pos.y > c.floorY {
		panic("Position is below the floor")
	}
	// Update the corners of the cave
	if pos.x < c.topLeft.x {
		c.topLeft.x = pos.x
	} else if pos.x > c.bottomRight.x {
		c.bottomRight.x = pos.x
	}
	if pos.y < c.topLeft.y {
		c.topLeft.y = pos.y
	} else if pos.y > c.bottomRight.y {
		c.bottomRight.y = pos.y
	}

	// Add the sand piece
	c.sand[pos] = true
}

func (c *Cave) IsBlocked(pos Pos) bool {
	if pos.y > c.floorY {
		panic("Position is below the floor")
	}
	return pos.y == c.floorY || c.rock[pos] || c.sand[pos]
}

func (c *Cave) RocksFromPath(path []Pos) {
	for i := 0; i < len(path)-1; i++ {
		start, end := path[i], path[i+1]
		// Get the direction of the path
		xDiff := end.x - start.x
		yDiff := end.y - start.y
		// Note: the y axis is inverted
		if yDiff < 0 {
			// If the path is going up
			for y := start.y; y >= end.y; y-- {
				c.AddRock(Pos{start.x, y})
			}
		} else if yDiff > 0 {
			// If the path is going down
			for y := start.y; y <= end.y; y++ {
				c.AddRock(Pos{start.x, y})
			}
		} else if xDiff < 0 {
			// If the path is going left
			for x := start.x; x >= end.x; x-- {
				c.AddRock(Pos{x, start.y})
			}
		} else if xDiff > 0 {
			// If the path is going right
			for x := start.x; x <= end.x; x++ {
				c.AddRock(Pos{x, start.y})
			}
		}
	}
}

func (c *Cave) SimulateSandFall() (sourceBlocked bool) {
	if c.movingSand == (Pos{}) {
		// If no sand is moving and source is blocked, return
		if c.IsBlocked(c.sandSource) {
			return true
		} else {
			c.movingSand = c.sandSource
		}
	}

	nextPosOpts := []Pos{
		// Down
		{c.movingSand.x, c.movingSand.y + 1},
		// Diagonally to the left
		{c.movingSand.x - 1, c.movingSand.y + 1},
		// Diagonally to the right
		{c.movingSand.x + 1, c.movingSand.y + 1},
	}
	for _, nextPos := range nextPosOpts {
		if !c.IsBlocked(nextPos) {
			// The sand can move, update the moving sand position
			c.movingSand = nextPos
			return false
		}
	}
	// The sand can't move, add it to the sand pile
	c.AddSand(c.movingSand)
	c.movingSand = Pos{}
	return false
}

func (c *Cave) PrintCave() {
	for y := c.topLeft.y; y <= c.bottomRight.y; y++ {
		for x := c.topLeft.x; x <= c.bottomRight.x; x++ {
			pos := Pos{x, y}
			if pos == c.sandSource && !c.sand[pos] {
				fmt.Printf("v")
			} else if c.rock[pos] || c.floorY == y {
				fmt.Printf("░")
			} else if c.sand[pos] {
				fmt.Printf("▓")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
}

type Pos struct {
	x, y int
}
