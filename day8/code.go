package main

import (
	"fmt"
	"github.com/MSkrzypietz/aoc-2023/day8/network"
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

	steps := network.Build(file).Navigate1()
	return strconv.Itoa(steps)
}

func solvePart2() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	steps := network.Build(file).Navigate2()
	return strconv.Itoa(steps)
}
