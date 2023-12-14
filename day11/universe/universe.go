package universe

import (
	"bufio"
	"io"
)

type Universe struct {
	Galaxies []position
}

func Build(r io.Reader, isOldUniverse bool) *Universe {
	var grid [][]int32
	filledCols := make(map[int]bool)
	filledRows := make(map[int]bool)
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		row := make([]int32, 0)
		for j, ch := range scanner.Text() {
			row = append(row, ch)
			if ch == '#' {
				filledRows[i] = true
				filledCols[j] = true
			}
		}
		grid = append(grid, row)
	}

	offset := 1
	if isOldUniverse {
		offset = 1_000_000 - 1
	}

	var positions []position
	var rowOffset int
	for i := 0; i < len(grid); i++ {
		if !filledRows[i] {
			rowOffset += offset
		}
		colOffset := 0
		for j := 0; j < len(grid[i]); j++ {
			if !filledCols[j] {
				colOffset += offset
			}
			if grid[i][j] == '#' {
				positions = append(positions, position{x: j + colOffset, y: i + rowOffset})
			}
		}
	}

	return &Universe{Galaxies: positions}
}
