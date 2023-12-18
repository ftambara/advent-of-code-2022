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
	totalStrength := 0
	for cycle := 20; cycle < len(values); cycle += 40 {
		signalStrength := values[cycle-1] * cycle
		fmt.Printf("%vth cycle: adding %v\n", cycle, signalStrength)
		totalStrength += signalStrength
	}
	fmt.Printf("Total signal strength: %v\n", totalStrength)
}

func noop(values []int) []int {
	return append(values, values[len(values)-1])
}

func addx(values []int, change int) []int {
	last := values[len(values)-1]
	return append(values, last, last+change)
}
