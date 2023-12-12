package main

import (
	"fmt"
	"github.com/MSkrzypietz/aoc-2023/day05/almanac"
	"github.com/MSkrzypietz/aoc-2023/day05/interval"
	"log"
	"math"
	"os"
	"slices"
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

	almanac := almanac.Build(file)

	var locations []int
	for _, seed := range almanac.Seeds() {
		position := seed
		for _, mapTable := range almanac.OrderedMapTables() {
			for _, mapEntry := range mapTable.OrderedMapEntries() {
				srcInterval := mapEntry.SrcInterval()
				if srcInterval.Min <= position && position <= srcInterval.Max {
					position = position + mapEntry.Dst - mapEntry.Src
					break
				}
			}
		}
		locations = append(locations, position)
	}

	minLocation := slices.Min(locations)
	return strconv.Itoa(minLocation)
}

func solvePart2() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	almanac := almanac.Build(file)

	var inputs, outputs []interval.Interval
	inputs = append(inputs, almanac.SeedIntervals()...)

	for _, mapTable := range almanac.OrderedMapTables() {
		for len(inputs) > 0 {
			input := inputs[0]
			inputs = inputs[1:]

			for _, mapEntry := range mapTable.OrderedMapEntries() {
				srcRange := mapEntry.SrcInterval()
				if srcRange.Contains(input) {
					outputs = append(outputs, mapEntry.Apply(input))
					break
				}
				if overlap, found := srcRange.Overlaps(input); found {
					inputs = append(inputs, interval.Interval{
						Min: mapEntry.SrcMax() + 1,
						Max: input.Max,
					})
					outputs = append(outputs, mapEntry.Apply(overlap))
					break
				}
			}
		}
		inputs = append(inputs, outputs...)
		outputs = make([]interval.Interval, 0)
	}

	minLocation := math.MaxInt
	for _, input := range inputs {
		if input.Min < minLocation {
			minLocation = input.Min
		}
	}
	return strconv.Itoa(minLocation)
}
