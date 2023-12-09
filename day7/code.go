package main

import (
	"bufio"
	"fmt"
	"github.com/MSkrzypietz/aoc-2023/day7/hand"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var day = os.Getenv("DAY")

func main() {
	fmt.Printf("Solutions to day %s\n", day)
	fmt.Println("Part 1:", solvePart(false))
	fmt.Println("Part 2:", solvePart(true))
}

func solvePart(usesJokers bool) string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []hand.Hand
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		bid, _ := strconv.Atoi(line[1])
		hands = append(hands, hand.NewHand(line[0], bid, usesJokers))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j])
	})

	result := 0
	for i, h := range hands {
		result += h.Bid * (i + 1)
	}

	return strconv.Itoa(result)
}
