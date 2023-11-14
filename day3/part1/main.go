package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Get list of rucksacks
	rucksacks, err := readRucksacks("day3/input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	// Get items on each compartment
	totalPriority := 0
	for _, r := range rucksacks {
		left, right := evenSplit(r)
		common := setInter(rucksackToSet(left), rucksackToSet(right))
		for item := range common {
			totalPriority += itemPriority(item)
		}
	}
	fmt.Printf("Total priority: %v\n", totalPriority)
}

type Set map[byte]bool

func rucksackToSet(rucksack []byte) Set {
	set := make(Set, len(rucksack))
	for _, item := range rucksack {
		set[item] = true
	}
	return set
}

func setInter(set1, set2 Set) Set {
	// Return set1 & set2
	res := Set{}
	for k := range set1 {
		if set2[k] {
			res[k] = true
		}
	}
	return res
}

func readRucksacks(filename string) ([][]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	rucksacks := [][]byte{}
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}
		rucksacks = append(rucksacks, []byte(ln))
	}
	if err := scanner.Err(); err != nil {
		return rucksacks, err
	}
	return rucksacks, nil
}

func itemPriority(c byte) int {
	switch {
	case c >= 'A' && c <= 'Z':
		return int(c - 'A' + 27)
	case c >= 'a' && c <= 'z':
		return int(c - 'a' + 1)
	default:
		panic("Illegal item type")
	}
}

func evenSplit(items []byte) ([]byte, []byte) {
	return items[:len(items)/2], items[len(items)/2:]
}
