package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota
	paper
	scissors
)

func main() {
	rounds, err := readStrategy("day2/input.txt")
	if err != nil {
		fmt.Printf("Error while reading day2/input.txt: %v", err)
	}
	myTotal, hisTotal := 0, 0
	for _, round := range rounds {
		hisHand := whichHand(round[0])
		myHand := whichHand(round[1])

		myTotal += score(myHand, hisHand)
		hisTotal += score(hisHand, myHand)
	}
	fmt.Printf("My total score: %v\n", myTotal)
	fmt.Printf("His total score: %v\n", hisTotal)
}

func readStrategy(filename string) ([][2]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	rounds := [][2]string{}
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			continue
		}
		round := strings.Split(ln, " ")
		if len(round) != 2 {
			err = fmt.Errorf("unrecognized line format: %v", ln)
			return rounds, err
		}
		rounds = append(rounds, [2]string(round))
	}
	if err := scanner.Err(); err != nil {
		return rounds, err
	}
	return rounds, nil
}

func score(mine, his int) int {
	s := scoreHand(mine)
	s += scoreOutcome(mine, his)
	return s
}

func whichHand(play string) int {
	switch play {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		panic("Illegal play!")
	}
}

func scoreHand(hand int) int {
	switch hand {
	case rock:
		return 1
	case paper:
		return 2
	case scissors:
		return 3
	default:
		panic("Illegal hand!")
	}
}

func scoreOutcome(mine, his int) int {
	switch {
	case mine == rock && his == scissors ||
		mine == paper && his == rock ||
		mine == scissors && his == paper:
		return 6 // Won
	case mine == his:
		return 3 // Draw
	default:
		return 0 // Lost
	}
}
