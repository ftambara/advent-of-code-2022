package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	monkeys, err := readMonkeys("day11/input.txt")
	if err != nil {
		panic(err)
	}

	// Find a the least common multiple of all test divisors
	// Seen at https://www.youtube.com/watch?v=Aot_ORkkvP4
	divisor := 1
	for _, monkey := range monkeys {
		divisor *= monkey.testDivisor
	}

	for turn := 0; turn < 10000*len(monkeys); turn++ {
		monkey := &(monkeys[turn%len(monkeys)])
		for _, item := range monkey.items {
			worry := monkey.operation(item)
			monkey.inspectedCount++
			var destMonkey *Monkey
			if worry%monkey.testDivisor == 0 {
				destMonkey = &monkeys[monkey.trueDest]
			} else {
				destMonkey = &monkeys[monkey.falseDest]
			}
			// Use modulo arithmetic to avoid overflow of worry
			worry = worry % divisor

			destMonkey.items = append(destMonkey.items, worry)
			monkey.items = monkey.items[1:]
		}
	}

	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", monkey.id, monkey.inspectedCount)
	}
}

type Monkey struct {
	id             int
	items          []int
	operation      func(int) int
	testDivisor    int
	trueDest       int
	falseDest      int
	inspectedCount int
}

func readMonkeys(filename string) ([]Monkey, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	monkeys := []Monkey{}
	for scanner.Scan() {
		ln := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(ln, "Monkey") {
			var monkeyID int
			fmt.Sscanf(ln, "Monkey %d", &monkeyID)
			scanner.Scan()

			ln = strings.TrimSpace(scanner.Text())
			ln = strings.TrimPrefix(ln, "Starting items: ")
			startingItems := []int{}
			for _, item := range strings.Split(ln, ", ") {
				num, err := strconv.Atoi(item)
				if err != nil {
					return nil, err
				}
				startingItems = append(startingItems, num)
			}

			var operation func(int) int
			scanner.Scan()
			ln = strings.TrimSpace(scanner.Text())
			operation = readMonkeyOperation(strings.TrimPrefix(ln, "Operation: "))

			var testDivisor int
			scanner.Scan()
			ln = strings.TrimSpace(scanner.Text())
			fmt.Sscanf(ln, "Test: divisible by %d", &testDivisor)

			var trueDest int
			scanner.Scan()
			ln = strings.TrimSpace(scanner.Text())
			fmt.Sscanf(ln, "If true: throw to monkey %d", &trueDest)

			var falseDest int
			scanner.Scan()
			ln = strings.TrimSpace(scanner.Text())
			fmt.Sscanf(ln, "If false: throw to monkey %d", &falseDest)

			monkeys = append(monkeys, Monkey{
				id:          monkeyID,
				items:       startingItems,
				operation:   operation,
				testDivisor: testDivisor,
				trueDest:    trueDest,
				falseDest:   falseDest,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return monkeys, err
	}
	return monkeys, err
}

func readMonkeyOperation(s string) func(int) int {
	var operator, operand1, operand2 string
	fmt.Sscanf(s, "new = %s %s %s", &operand1, &operator, &operand2)
	var mathOp func(int, int) int
	switch operator {
	case "+":
		mathOp = func(x, y int) int { return x + y }
	case "-":
		mathOp = func(x, y int) int { return x - y }
	case "*":
		mathOp = func(x, y int) int { return x * y }
	case "/":
		mathOp = func(x, y int) int { return x / y }
	default:
		panic(fmt.Sprintf("Unknown operator: %s", operator))
	}

	if operand1 != "old" {
		num, err := strconv.Atoi(operand1)
		if err != nil {
			panic(fmt.Sprintf("Unknown operand: %s", operand1))
		}
		return func(x int) int { return mathOp(num, x) }
	} else if operand2 != "old" {
		num, err := strconv.Atoi(operand2)
		if err != nil {
			panic(fmt.Sprintf("Unknown operand: %s", operand2))
		}
		return func(x int) int { return mathOp(x, num) }
	} else {
		return func(x int) int { return mathOp(x, x) }
	}
}
