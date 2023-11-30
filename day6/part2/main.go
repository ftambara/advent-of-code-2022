package main

import (
	"bufio"
        "fmt"
	"os"
)

func main() {
	f, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	last14 := newTuner(14)
	for scanner.Scan() {
		// Assume string is a valid ASCII character
		last14 = *last14.push(scanner.Text()[0])
                marker, done := last14.marker()
                if !done {
                    continue
                }
		fmt.Printf("The marker is at position %v\n", marker)
                return
	}

}

type Tuner struct {
	contents []byte
	max      int
	count    int
}

func newTuner(max int) Tuner {
	return Tuner{
		contents: make([]byte, max),
		count:    0,
		max:      max,
	}
}

func (s *Tuner) push(b byte) *Tuner {
	if s.count < s.max {
		s.contents[s.count] = b
	} else {
		newContents := make([]byte, s.max)
		// Drop the first element to make room for the new one
		copy(newContents, s.contents[1:s.max])
		newContents[s.max-1] = b
		s.contents = newContents
	}
	s.count++
	return s
}

func (s *Tuner) marker() (pos int, done bool) {
	if s.count < s.max {
		return s.count, false
	}
	// Check if all characters are different
	unique := make(map[string]bool, s.count)
	for _, b := range s.contents {
		if unique[string(b)] {
			return s.count, false
		}
                unique[string(b)] = true
	}
	return s.count, true
}
