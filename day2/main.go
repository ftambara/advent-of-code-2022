package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ROCK     int = 0
	PAPER    int = 1
	SCISSORS int = 2
)

func main() {
	rounds, err := readStrategy("day2/input.txt")
	if err != nil {
		fmt.Printf("Error while reading day2/input.txt: %v", err)
	}
	myTotal, hisTotal := 0, 0
	for i, round := range rounds {
		mine, his := score(round)
		myTotal += mine
		hisTotal += his
		fmt.Printf("L%v: %v -> mine: %v, his: %v\n", i, round, mine, his)
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

func score(round [2]string) (mine, his int) {
	his = whichHand(round[0])
	mine = whichHand(round[1])

	myScore := scoreHand(mine)
	fmt.Println(scoreHand(mine))
	fmt.Println(scoreOutcome(mine, his))
	myScore += scoreOutcome(mine, his)

	hisScore := scoreHand(his)
	hisScore += scoreOutcome(his, mine)

	return myScore, hisScore
}

func whichHand(play string) int {
	switch play {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	default:
		panic("Illegal play!")
	}
}

func scoreHand(hand int) int {
	switch hand {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	default:
		panic("Illegal hand!")
	}
}

func scoreOutcome(mine, his int) int {
	switch {
	case mine == ROCK && his == SCISSORS ||
		mine == PAPER && his == ROCK ||
		mine == SCISSORS && his == PAPER:
		return 6 // Won
	case mine == his:
		return 3 // Draw
	default:
		return 0 // Lost
	}
}
