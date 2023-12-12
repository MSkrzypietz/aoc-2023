package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = os.Getenv("DAY")

func main() {
	fmt.Printf("Solutions to day %s\n", day)
	fmt.Println("Part 1:", solvePart1())
	fmt.Println("Part 2:", solvePart2())
}

func solvePart1() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cardValue := 0
		_, numbers := splitLink(line, ":")
		winners, picks := splitLink(numbers, "|")
		for _, pick := range strings.Fields(picks) {
			for _, winner := range strings.Fields(winners) {
				if pick == winner {
					if cardValue == 0 {
						cardValue++
					} else {
						cardValue *= 2
					}
				}
			}
		}
		result += cardValue
	}

	return strconv.Itoa(result)
}

func splitLink(s string, sep string) (string, string) {
	chunks := strings.Split(s, sep)
	return chunks[0], chunks[1]
}

func solvePart2() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	var cards []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		copies := 0
		if len(cards) > 0 {
			copies = cards[0]
			cards = cards[1:]
		}

		matches := 0
		_, numbers := splitLink(line, ":")
		winners, picks := splitLink(numbers, "|")
		for _, pick := range strings.Fields(picks) {
			for _, winner := range strings.Fields(winners) {
				if pick == winner {
					matches++
				}
			}
		}

		cardInstances := copies + 1
		result += cardInstances
		if matches > 0 {
			for i := 0; i < matches; i++ {
				if i < len(cards) {
					cards[i] += cardInstances
				} else {
					cards = append(cards, cardInstances)
				}
			}
		}
	}

	return strconv.Itoa(result)
}
