package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	stacks, steps := readInstructions("day5/input.txt")
	for _, s := range steps {
		stacks = move(stacks, s)
	}
	printSite(stacks)

	// Print topmost crates
	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}
	fmt.Println()
}

type step struct {
	n    int
	from int
	to   int
}

type site [][]byte

func move(stacks site, s step) site {
	// Step positions count from 1, rectify
	from := s.from - 1
	to := s.to - 1
	for i := 0; i < s.n; i++ {
		fromHeight := len(stacks[from]) - 1
		// Check if move is valid
		if fromHeight < 0 {
			panic("trying to take a crate from an empty stack")
		}
		// Grow the 'to' slice
		stacks[to] = append(stacks[to], stacks[from][fromHeight])
		// Shrink the 'from' slice
		stacks[from] = stacks[from][:fromHeight]
	}
	return stacks
}

func printSite(s site) {
	for _, stack := range s {
		for _, crate := range stack {
			// Print crates left to right
			fmt.Printf("[%c] ", crate)
		}
		fmt.Println()
	}
}

func readInstructions(filename string) (site, []step) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stacks := readStacks(f)
	f.Seek(0, io.SeekStart)
	steps := readSteps(f)
	return stacks, steps
}

func readStacks(f *os.File) site {
	scanner := bufio.NewScanner(f)

	lines := []string{}
	// Read until the line doesn't contain a crate
	for scanner.Scan() {
		ln := scanner.Text()
		if !strings.Contains(ln, "[") {
			break
		}
		lines = append(lines, ln)
	}

	// Build stacks from the base
	baseLine := lines[len(lines)-1]
	stacksN := strings.Count(baseLine, "[")
	stacks := make(site, stacksN)
	start := strings.Index(baseLine, "[") + 1
	if start <= 0 {
		panic("malformed crate string")
	}
	curStack := 0
	// Traverse by column
	for i := start; i < len(baseLine); i += len("[_] ") {
		// Go up until you reach the limit or there is no crate
		for j := len(lines) - 1; j >= 0; j-- {
			crateLetter := lines[j][i]
			if crateLetter == ' ' {
				// End of the stack, continue with the next
				break
			}
			// Add one more crate
			stacks[curStack] = append(stacks[curStack], crateLetter)
		}
		curStack++
	}

	return stacks
}

func readSteps(f *os.File) []step {
	scanner := bufio.NewScanner(f)

	steps := []step{}
	for scanner.Scan() {
		ln := scanner.Text()
		// Skip lines that don't contain moves
		if !strings.Contains(ln, "move") {
			continue
		}
		var (
			n    int
			from int
			to   int
		)
		fmt.Sscanf(ln, "move %d from %d to %d", &n, &from, &to)
		steps = append(steps, step{
			n:    n,
			from: from,
			to:   to,
		})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return steps
}
