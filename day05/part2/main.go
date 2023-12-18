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
	for _, step := range steps {
		stacks = moveAtOnce(stacks, step)
	}
	printStacks(stacks)

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

type stacks [][]byte

// Move all crates specified by the step at once
func moveAtOnce(stacks stacks, step step) stacks {
	// Step positions count from 1, rectify
	from := step.from - 1
	to := step.to - 1
	fromLast := len(stacks[from])
	fromFirst := fromLast - step.n
	// Check if move is valid
	if fromFirst < 0 {
		panic("trying to move too many crates")
	}
	// Grow the 'to' slice
	stacks[to] = append(stacks[to], stacks[from][fromFirst:fromLast]...)
	// Shrink the 'from' slice
	stacks[from] = stacks[from][:fromFirst]
	return stacks
}

func printStacks(stacks stacks) {
	for _, stack := range stacks {
		for _, crate := range stack {
			// Print crates left to right
			fmt.Printf("[%c] ", crate)
		}
		fmt.Println()
	}
}

func readInstructions(filename string) (stacks, []step) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stacks := readStacks(f)
	f.Seek(0, io.SeekStart)
	steps := readSteps(f)
	return stacks, steps
}

func readStacks(f *os.File) stacks {
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
	stacks := make(stacks, stacksN)
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
