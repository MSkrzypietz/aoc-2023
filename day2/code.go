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

LineLoop:
	for scanner.Scan() {
		line := scanner.Text()

		game, draws := splitLink(line, ":")
		gameId, _ := strconv.Atoi(strings.Split(game, " ")[1])
		for _, draw := range strings.Split(draws, ";") {
			for _, item := range strings.Split(draw, ",") {
				item = strings.TrimSpace(item)
				countString, color := splitLink(item, " ")
				count, _ := strconv.Atoi(countString)
				if !isPossibleBag(color, count) {
					continue LineLoop
				}
			}
		}
		result += gameId
	}

	return strconv.Itoa(result)
}

func splitLink(s string, sep string) (string, string) {
	chunks := strings.Split(s, sep)
	return chunks[0], chunks[1]
}

func isPossibleBag(color string, count int) bool {
	switch color {
	case "red":
		return count <= 12
	case "green":
		return count <= 13
	case "blue":
		return count <= 14
	default:
		return false
	}
}

func solvePart2() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var reds, greens, blues []int
		_, draws := splitLink(line, ":")
		for _, draw := range strings.Split(draws, ";") {
			for _, item := range strings.Split(draw, ",") {
				item = strings.TrimSpace(item)
				countString, color := splitLink(item, " ")
				count, _ := strconv.Atoi(countString)

				switch color {
				case "red":
					reds = append(reds, count)
				case "green":
					greens = append(greens, count)
				case "blue":
					blues = append(blues, count)
				}
			}
		}

		maxRed := slices.Max(reds)
		maxGreen := slices.Max(greens)
		maxBlue := slices.Max(blues)
		result += maxRed * maxGreen * maxBlue
	}

	return strconv.Itoa(result)
}
