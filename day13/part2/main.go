package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	pairs := readPacketPairs("day13/input.txt")

	// Flatten pairs
	packets := []ListItem{}
	for _, pair := range pairs {
		left, right := pair[0], pair[1]
		packets = append(packets, left)
		packets = append(packets, right)
	}

	// Append divider packets
	div1 := ListItem{ListItem{IntItem(2)}}
	div2 := ListItem{ListItem{IntItem(6)}}
	packets = append(packets, div1)
	packets = append(packets, div2)

	// Sort packets
	slices.SortStableFunc(packets, compareListItems)

	// Find divider packets positions
	div1Pos := -1
	div2Pos := -1
	for i, packet := range packets {
		if packet.CompareTo(div1) == 0 {
			if div1Pos != -1 {
				panic("Duplicate div1")
			}
			div1Pos = i + 1
		}
		if packet.CompareTo(div2) == 0 {
			if div2Pos != -1 {
				panic("Duplicate div2")
			}
			div2Pos = i + 1
		}
	}

	fmt.Printf("Decoder key: %d\n", div1Pos*div2Pos)
}

func readPacketPairs(filename string) [][2]ListItem {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var pairs [][2]ListItem
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		left := parsePacket(line)

		scanner.Scan()
		line = scanner.Text()
		right := parsePacket(line)

		pairs = append(pairs, [2]ListItem{left, right})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs
}

func parsePacket(packetStr string) ListItem {
	if !strings.HasPrefix(packetStr, "[") || !strings.HasSuffix(packetStr, "]") {
		panic("Invalid packet string: missing enclosing brackets")
	}
	var itemListItem ListItem
	for i := 1; i < len(packetStr)-1; {
		if packetStr[i] == '[' {
			// Parse list
			// Find the end of the list
			end := i + 1
			bracketCount := 1
			for bracketCount > 0 {
				if end == len(packetStr)-1 {
					panic("Invalid packet string: missing closing bracket")
				}
				switch packetStr[end] {
				case '[':
					bracketCount++
				case ']':
					bracketCount--
				}
				end++
			}

			// Parse the list
			itemListItem = append(itemListItem, parsePacket(packetStr[i:end]))

			// Skip to the end of the list
			i = end
		} else if isNumber(packetStr[i]) {
			// Parse number
			// Find the end of the number
			end := i + 1
			for end < len(packetStr) && isNumber(packetStr[end]) {
				end++
			}

			// Parse the number
			num, err := strconv.Atoi(packetStr[i:end])
			if err != nil {
				panic(err)
			}
			itemListItem = append(itemListItem, IntItem(num))

			// Skip to the end of the number
			i = end
		}
		// Verify that the next character is either a comma or the end of the packet
		if packetStr[i] != ',' && packetStr[i] != ']' {
			panic("Invalid packet string: unexpected character")
		}

		// Skip to the next character after the comma
		i++
	}
	return itemListItem
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

type Comparable interface {
	CompareTo(Comparable) int
}

type IntItem int

func (i IntItem) CompareTo(c Comparable) int {
	switch val := c.(type) {
	case IntItem:
		if i < val {
			return -1
		} else if i > val {
			return 1
		}
		return 0
	case ListItem:
		return IntToListItem(i).CompareTo(c)
	default:
		panic("Invalid type")
	}
}

type ListItem []Comparable

func IntToListItem(i IntItem) ListItem {
	return ListItem{i}
}

func (l ListItem) CompareTo(c Comparable) int {
	switch val := c.(type) {
	case IntItem:
		return l.CompareTo(IntToListItem(c.(IntItem)))
	case ListItem:
		l1 := l
		l2 := val
		for i := 0; i < len(l1) && i < len(l2); i++ {
			if cmp := l1[i].CompareTo(l2[i]); cmp != 0 {
				return cmp
			}
		}
		if len(l1) < len(l2) {
			return -1
		}
		if len(l1) > len(l2) {
			return 1
		}
		return 0
	default:
		panic("Invalid type")
	}
}

func areOrdered(left, right ListItem) bool {
	return left.CompareTo(right) == -1
}

func compareListItems(a, b ListItem) int {
	return a.CompareTo(b)
}
