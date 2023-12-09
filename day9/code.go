package main

import (
	"bufio"
	"fmt"
	"github.com/MSkrzypietz/aoc-2023/utils"
	"log"
	"os"
	"strconv"
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
		rows := buildRows(scanner)

		rows[len(rows)-1] = append(rows[len(rows)-1], 0)
		for i := len(rows) - 2; i >= 0; i-- {
			rows[i] = append(rows[i], rows[i+1][len(rows[i+1])-1]+rows[i][len(rows[i])-1])
		}

		result += rows[0][len(rows[0])-1]
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
		rows := buildRows(scanner)

		rows[len(rows)-1] = append([]int{0}, rows[len(rows)-1]...)
		for i := len(rows) - 2; i >= 0; i-- {
			rows[i] = append([]int{rows[i][0] - rows[i+1][0]}, rows[i]...)
		}

		result += rows[0][0]
	}

	return strconv.Itoa(result)
}

func buildRows(scanner *bufio.Scanner) [][]int {
	rows := append([][]int{}, utils.IntFields(scanner.Text()))

	lastRow := rows[len(rows)-1]
	for !containsOnlyZeroes(lastRow) {
		row := make([]int, len(lastRow)-1)
		for i := 0; i < len(lastRow)-1; i++ {
			row[i] = lastRow[i+1] - lastRow[i]
		}
		rows = append(rows, row)
		lastRow = row
	}

	return rows
}

func containsOnlyZeroes(row []int) bool {
	for _, item := range row {
		if item != 0 {
			return false
		}
	}
	return true
}
