package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rucksacks, err := readRucksacks("day3/input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	groups := groupsOf3(rucksacks)
	prioritySum := 0
	for _, g := range groups {
		s1 := rucksackToSet(g[0])
		s2 := rucksackToSet(g[1])
		s3 := rucksackToSet(g[2])
		badges := setToSlice(setInter(s1, setInter(s2, s3)))
		if len(badges) == 0 {
			panic("No badge found!")
		} else if len(badges) > 1 {
			panic("There should only be one badge!")
		}
		badge := badges[0]
		prioritySum += itemPriority(badge)
	}
	fmt.Printf("Priority sum: %v\n", prioritySum)
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

func groupsOf3(rucksacks [][]byte) [][3][]byte {
	n := len(rucksacks) / 3
	groups := make([][3][]byte, n)
	for i := 0; i < n; i++ {
		groups[i] = [3][]byte(rucksacks[3*i : 3*i+3])
	}
	return groups
}

type Set map[byte]bool

func rucksackToSet(rucksack []byte) Set {
	set := make(Set, len(rucksack))
	for _, item := range rucksack {
		set[item] = true
	}
	return set
}

func setToSlice(s Set) []byte {
	slice := make([]byte, len(s))
	i := 0
	for k := range s {
		slice[i] = k
		i++
	}
	return slice
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
