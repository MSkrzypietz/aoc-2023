package main

import (
	"fmt"
	"github.com/MSkrzypietz/aoc-2023/day06/race"
	"log"
	"math"
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

	result := 1
	races := race.BuildRaces(file)
	for _, r := range races {
		x1, x2 := solveQuadraticFormula(-1.0, float64(r.Time), -1.0*float64(r.Distance))
		result *= int(math.Ceil(x2) - math.Floor(x1) - 1)
	}

	return strconv.Itoa(result)
}

func solvePart2() string {
	file, err := os.Open(fmt.Sprintf("day%s/input.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := race.BuildGiantRaces(file)
	x1, x2 := solveQuadraticFormula(-1.0, float64(r.Time), -1.0*float64(r.Distance))
	result := int(math.Ceil(x2) - math.Floor(x1) - 1)

	return strconv.Itoa(result)
}

func solveQuadraticFormula(a, b, c float64) (float64, float64) {
	discriminant := math.Pow(b, 2) - 4.0*a*c
	x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return x1, x2
}
