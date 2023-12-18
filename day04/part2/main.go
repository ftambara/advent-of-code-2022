package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pairs, err := readAssignmentPairs("day4/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	overlapCount := 0
	for _, pair := range pairs {
		if overlap(pair[0], pair[1]) {
			overlapCount++
		}
	}
	fmt.Printf("A total of %v pairs overlap.\n", overlapCount)
}

func readAssignmentPairs(filename string) ([][2]assignment, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	assignments := [][2]assignment{}
	for scanner.Scan() {
		ln := scanner.Text()
		assignments = append(assignments, decodeAssignmentPairs(ln))
	}
	if err := scanner.Err(); err != nil {
		return assignments, err
	}
	return assignments, nil
}

// line format: 1-3,5-7
// output: [assignment{1,3}, assignment{5,7}]
func decodeAssignmentPairs(line string) [2]assignment {
	parts := strings.Split(line, ",")
	assignments := [2]assignment{}
	for i, part := range parts {
		pair := strings.Split(part, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		assignments[i] = assignment{start, end}
	}
	return assignments
}

type assignment [2]int

func overlap(a1, a2 assignment) bool {
	return !(a1[0] > a2[1] || a1[1] < a2[0])
}
