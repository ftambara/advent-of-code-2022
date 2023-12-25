package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := []int{1}

	f, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		ln := scanner.Text()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		instr := strings.Split(ln, " ")[0]
		switch instr {
		case "noop":
			values = noop(values)
		case "addx":
			change, err := strconv.Atoi(strings.Split(ln, " ")[1])
			if err != nil {
				panic(err)
			}
			values = addx(values, change)
		default:
			panic(fmt.Errorf("wrong instruction: %v", instr))
		}
	}

	// Analyze values
	width := 40
	var rows int
	if len(values)%width == 0 {
		rows = len(values) / width
	} else {
		rows = len(values)/width + 1
	}
	crt := make([][]byte, rows)
	for i := range values {
		column := i%width + 1
		row := i / width
		// Add one to distance due to the starting position of the sprite
		distance := values[i] - column + 1
		var pixel byte
		if distance <= 1 && distance >= -1 {
			// Sprite is visible
			pixel = '#'
		} else {
			pixel = '.'
		}
		crt[row] = append(crt[row], pixel)
	}

	// Print CRT
	for _, row := range crt {
		for _, pixel := range row {
			fmt.Printf("%c", pixel)
		}
		fmt.Println()
	}
}

func noop(values []int) []int {
	return append(values, values[len(values)-1])
}

func addx(values []int, change int) []int {
	last := values[len(values)-1]
	return append(values, last, last+change)
}
