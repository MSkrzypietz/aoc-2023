package almanac

import (
	"bufio"
	"github.com/MSkrzypietz/aoc-2023/day5/interval"
	"github.com/MSkrzypietz/aoc-2023/utils"
	"io"
	"math"
	"sort"
	"strings"
)

type MapEntry struct {
	Src int
	Dst int
	Len int
}

type MapTable struct {
	entries []MapEntry
}

type Almanac struct {
	mapTables []MapTable
	seeds     []int
}

func Build(r io.Reader) *Almanac {
	almanac := &Almanac{}

	scanner := bufio.NewScanner(r)

	if scanner.Scan() {
		almanac.seeds = utils.IntFields(strings.Split(scanner.Text(), ":")[1])
	}

	scanner.Scan()
	scanner.Scan()

	var mapTables []MapTable
	var mapEntries []MapEntry
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			mapTables = append(mapTables, MapTable{mapEntries})
			mapEntries = make([]MapEntry, 0)
			scanner.Scan()
			continue
		}

		fields := utils.IntFields(line)
		mapEntries = append(mapEntries, MapEntry{
			Src: fields[1],
			Dst: fields[0],
			Len: fields[2],
		})
	}
	mapTables = append(mapTables, MapTable{mapEntries})

	var outputMapTables []MapTable
	for _, mapTable := range mapTables {
		entries := mapTable.entries
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Src < entries[j].Src
		})

		var orderedMapEntries []MapEntry
		if entries[0].Src > 0 {
			orderedMapEntries = append([]MapEntry{NewMapEntry(0, entries[0].Src-1)})
		} else {
			orderedMapEntries = append([]MapEntry{entries[0]})
			entries = entries[1:]
		}

		for len(entries) > 0 {
			next := entries[0]
			entries = entries[1:]

			upper := orderedMapEntries[len(orderedMapEntries)-1].SrcMax()
			if upper+1 != next.Src {
				orderedMapEntries = append(orderedMapEntries, NewMapEntry(upper+1, next.Src-1))
			}
			orderedMapEntries = append(orderedMapEntries, next)
		}

		upper := orderedMapEntries[len(orderedMapEntries)-1].SrcMax()
		orderedMapEntries = append(orderedMapEntries, NewMapEntry(upper, math.MaxInt))

		outputMapTables = append(outputMapTables, MapTable{orderedMapEntries})
	}

	almanac.mapTables = outputMapTables
	return almanac
}

func (a *Almanac) Seeds() []int {
	return a.seeds
}

func (a *Almanac) SeedIntervals() []interval.Interval {
	var seeds []interval.Interval
	for i := 0; i < len(a.seeds); i += 2 {
		seeds = append(seeds, interval.Interval{
			Min: a.seeds[i],
			Max: a.seeds[i] + a.seeds[i+1] - 1,
		})
	}
	return seeds
}

func (a *Almanac) OrderedMapTables() []MapTable {
	return a.mapTables
}

func (m *MapTable) OrderedMapEntries() []MapEntry {
	return m.entries
}

func (m MapEntry) Apply(i interval.Interval) interval.Interval {
	return interval.Interval{
		Min: m.Dst + i.Min - m.Src,
		Max: m.Dst + i.Max - m.Src,
	}
}

func (m MapEntry) SrcInterval() interval.Interval {
	return interval.Interval{
		Min: m.Src,
		Max: m.Src + m.Len - 1,
	}
}

func (m MapEntry) SrcMax() int {
	return m.Src + m.Len - 1
}

func NewMapEntry(min, max int) MapEntry {
	return MapEntry{
		Src: min,
		Dst: min,
		Len: max - min + 1,
	}
}
