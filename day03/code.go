package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var day = os.Getenv("DAY")

type NumberIndexRange struct {
	value int
	from  int
	to    int
}

type Line struct {
	value string
	index int
}

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

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	result := 0
	for i, curr := range lines {
		adjLines := []string{curr}
		if i > 0 {
			adjLines = append(adjLines, lines[i-1])
		}
		if i < len(lines)-1 {
			adjLines = append(adjLines, lines[i+1])
		}

		for _, numberIndexRange := range getNumberIndexRanges(curr) {
			for _, adjLine := range adjLines {
				if ids := getEnginePartIds(adjLine, i, numberIndexRange.from, numberIndexRange.to, isEnginePart); len(ids) > 0 {
					result += numberIndexRange.value
					break
				}
			}
		}
	}

	return strconv.Itoa(result)
}

func getEnginePartIds(line string, lineIndex, from, to int, predicate func(int32) bool) []string {
	if from > 0 {
		from--
	}
	if to < len(line) {
		to++
	}

	var result []string
	for i := from; i < to; i++ {
		if predicate(int32(line[i])) {
			result = append(result, fmt.Sprintf("%d:%d", lineIndex, i))
		}
	}
	return result
}

func isEnginePart(ch int32) bool {
	return !isDigit(ch) && ch != '.'
}

func getNumberIndexRanges(line string) []*NumberIndexRange {
	var result []*NumberIndexRange

	digitStartIndex := -1
	for i, ch := range line {
		if isDigit(ch) {
			if digitStartIndex == -1 {
				digitStartIndex = i
			}
		} else if digitStartIndex != -1 {
			value, _ := strconv.Atoi(line[digitStartIndex:i])
			result = append(result, &NumberIndexRange{
				value: value,
				from:  digitStartIndex,
				to:    i,
			})
			digitStartIndex = -1
		}
	}

	if digitStartIndex != -1 {
		value, _ := strconv.Atoi(line[digitStartIndex:])
		result = append(result, &NumberIndexRange{
			value: value,
			from:  digitStartIndex,
			to:    len(line),
		})
	}

	return result
}

func isDigit(ch int32) bool {
	return '0' <= ch && ch <= '9'
}

func solvePart2() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	gearNumbers := make(map[string][]int)
	for i, curr := range lines {
		adjLines := []Line{{value: curr, index: i}}
		if i > 0 {
			adjLines = append(adjLines, Line{value: lines[i-1], index: i - 1})
		}
		if i < len(lines)-1 {
			adjLines = append(adjLines, Line{value: lines[i+1], index: i + 1})
		}

		for _, numberIndexRange := range getNumberIndexRanges(curr) {
			for _, adjLine := range adjLines {
				for _, gearId := range getEnginePartIds(adjLine.value, adjLine.index, numberIndexRange.from, numberIndexRange.to, isEngineGear) {
					gearNumbers[gearId] = append(gearNumbers[gearId], numberIndexRange.value)
				}
			}
		}
	}

	result := 0
	for _, numbers := range gearNumbers {
		if len(numbers) == 2 {
			result += numbers[0] * numbers[1]
		}
	}

	return strconv.Itoa(result)
}

func isEngineGear(ch int32) bool {
	return ch == '*'
}
