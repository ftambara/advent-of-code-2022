package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	win = iota
	draw
	lose
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
		myHand := counter(hisHand, whichOutcome(round[1]))

		hisTotal += score(hisHand, myHand)
		myTotal += score(myHand, hisHand)

		fmt.Printf("Round: %v\n", round)
		fmt.Printf("His hand: %v\n", hisHand)
		fmt.Printf("My hand: %v\n", myHand)
		fmt.Printf("His total: %v\n", hisTotal)
		fmt.Printf("My total: %v\n", myTotal)
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
func counter(hand, outcome int) int {
	if outcome == draw {
		return hand
	}

	strategy := map[int][2]int{
		// hand: to win, to lose
		rock:     {paper, scissors},
		scissors: {rock, paper},
		paper:    {scissors, rock},
	}
	pair, ok := strategy[hand]
	if !ok {
		panic("Illegal hand!")
	}
	toWin, toLose := pair[0], pair[1]
	if outcome == win {
		return toWin
	} else {
		return toLose
	}
}

func whichHand(repr string) int {
	switch repr {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	default:
		panic(fmt.Errorf("illegal play! %v", repr))
	}
}

func whichOutcome(repr string) int {
	switch repr {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	default:
		panic("Illegal outcome!")
	}
}
