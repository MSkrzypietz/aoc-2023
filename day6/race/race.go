package race

import (
	"bufio"
	"github.com/MSkrzypietz/aoc-2023/utils"
	"io"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func BuildRaces(r io.Reader) []Race {
	scanner := bufio.NewScanner(r)

	times := utils.IntFields(readInputValues(scanner))
	distances := utils.IntFields(readInputValues(scanner))

	var races []Race
	for i, time := range times {
		races = append(races, Race{
			Time:     time,
			Distance: distances[i],
		})
	}
	return races
}

func BuildGiantRaces(r io.Reader) Race {
	scanner := bufio.NewScanner(r)

	concatTime := strings.Join(strings.Fields(readInputValues(scanner)), "")
	time, _ := strconv.Atoi(concatTime)

	concatDistance := strings.Join(strings.Fields(readInputValues(scanner)), "")
	distance, _ := strconv.Atoi(concatDistance)

	return Race{
		Time:     time,
		Distance: distance,
	}
}

func readInputValues(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.Split(scanner.Text(), ":")[1]
}
