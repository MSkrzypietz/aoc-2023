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
var spelledNumbers = [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

		calibrationValue := ""
		for _, ch := range line {
			if isDigit(ch) {
				calibrationValue += string(ch)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			ch := int32(line[i])
			if isDigit(ch) {
				calibrationValue += string(ch)
				break
			}
		}

		value, _ := strconv.ParseInt(calibrationValue, 10, 64)
		result += int(value)
	}

	return strconv.Itoa(result)
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

		calibrationValue := ""
		seen := ""
		for _, ch := range line {
			seen = seen + string(ch)
			if v, ok := getCalibrationValue(seen, ch); ok {
				calibrationValue += v
				break
			}
		}

		seen = ""
		for i := len(line) - 1; i >= 0; i-- {
			ch := int32(line[i])
			seen = string(ch) + seen
			if v, ok := getCalibrationValue(seen, ch); ok {
				calibrationValue += v
				break
			}
		}

		value, _ := strconv.ParseInt(calibrationValue, 10, 64)
		result += int(value)
	}

	return strconv.Itoa(result)
}

func getCalibrationValue(seen string, ch int32) (string, bool) {
	if isDigit(ch) {
		return string(ch), true
	}

	if number, ok := getSpelledNumber(seen); ok {
		return strconv.Itoa(number), true
	}

	return "", false
}

func isDigit(ch int32) bool {
	return '0' <= ch && ch <= '9'
}

func getSpelledNumber(value string) (int, bool) {
	for i, number := range spelledNumbers {
		if strings.Index(value, number) != -1 {
			return i + 1, true
		}
	}
	return -1, false
}
