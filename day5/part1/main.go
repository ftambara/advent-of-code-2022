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
	fmt.Print(stacks, steps)
}

type step struct {
	n    int
	from int
	to   int
}

type site [][]byte

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
	// Seek the first line that contains a crate
	for scanner.Scan() {
		ln := scanner.Text()
		if strings.Contains(ln, "[") {
			break
		}
	}
	// Read until the next line that doesn't
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

	// Seek the first line that contains a move
	for scanner.Scan() {
		ln := scanner.Text()
		if strings.Contains(ln, "move") {
			break
		}
	}
	steps := []step{}
	for scanner.Scan() {
		ln := scanner.Text()
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
