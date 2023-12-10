package grid

import (
	"bufio"
	"github.com/MSkrzypietz/aoc-2023/day10/move"
	"github.com/MSkrzypietz/aoc-2023/day10/tile"
	"io"
)

type loop struct {
	start        position
	path         []position
	isInnerRight bool
}

type position struct {
	x        int
	y        int
	prevPos  *position
	prevMove move.Move
}

type Grid struct {
	grid  [][]tile.Tile
	start position
}

func Build(r io.Reader) *Grid {
	grid := &Grid{}

	i := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		var row []tile.Tile
		for j, ch := range line {
			t := tile.NewTile(ch)
			if t == tile.Start {
				grid.start = position{x: j, y: i}
			}
			row = append(row, t)
		}
		grid.grid = append(grid.grid, row)
		i++
	}

	return grid
}

func (g *Grid) isValidPosition(pos position) bool {
	if pos.x < 0 || pos.x >= len(g.grid[0]) {
		return false
	} else if pos.y < 0 || pos.y >= len(g.grid) {
		return false
	}
	return true
}

func (g *Grid) MaxSteps() int {
	seen := make([][]bool, len(g.grid))
	for i := range g.grid {
		seen[i] = make([]bool, len(g.grid[0]))
	}

	maxSteps := 0
	currSteps := 0
	positions := []position{g.start}
	for len(positions) > 0 {
		currPositions := append([]position{}, positions...)
		positions = make([]position, 0)

		for len(currPositions) > 0 {
			curr := currPositions[0]
			currPositions = currPositions[1:]

			if !g.isValidPosition(curr) || seen[curr.y][curr.x] || !g.isValidPrevMove(curr) {
				continue
			}

			seen[curr.y][curr.x] = true
			if currSteps > maxSteps {
				maxSteps = currSteps
			}

			positions = append(positions, g.allNextPositions(curr)...)
		}
		currSteps++
	}

	return maxSteps
}

func (g *Grid) EnclosedTiles() int {
	checked := make([][]bool, len(g.grid))
	for i := range g.grid {
		checked[i] = make([]bool, len(g.grid[0]))
	}

	var loops []loop
	for i := 0; i < len(g.grid); i++ {
		for j := 0; j < len(g.grid[i]); j++ {
			if checked[i][j] {
				continue
			}

			l, path, isLoop := g.findLoop(position{x: j, y: i})
			if isLoop {
				loops = append(loops, l)
			}

			for _, p := range path {
				checked[p.y][p.x] = true
			}
		}
	}

	return g.countEnclosedTiles(loops)
}

func (g *Grid) countEnclosedTiles(loops []loop) int {
	checked := make([][]bool, len(g.grid))
	for i := range g.grid {
		checked[i] = make([]bool, len(g.grid[0]))
	}

	for _, l := range loops {
		for _, pos := range l.path {
			checked[pos.y][pos.x] = true
		}
	}

	count := 0
	for _, l := range loops {
		for _, pos := range l.path {
			var innerFields []position
			t := g.grid[pos.y][pos.x]

			if l.isInnerRight {
				switch t {
				case tile.VerticalPipe:
					if pos.prevMove == move.Down {
						innerFields = append(innerFields, position{x: pos.x - 1, y: pos.y})
					} else if pos.prevMove == move.Up {
						innerFields = append(innerFields, position{x: pos.x + 1, y: pos.y})
					}
				case tile.HorizontalPipe:
					if pos.prevMove == move.Right {
						innerFields = append(innerFields, position{x: pos.x, y: pos.y + 1})
					} else if pos.prevMove == move.Left {
						innerFields = append(innerFields, position{x: pos.x, y: pos.y - 1})
					}
				case tile.BottomRightBend:
					if pos.prevMove == move.Right {
						innerFields = append(innerFields, position{x: pos.x + 1, y: pos.y}, position{x: pos.x, y: pos.y + 1})
					}
				case tile.BottomLeftBend:
					if pos.prevMove == move.Down {
						innerFields = append(innerFields, position{x: pos.x - 1, y: pos.y}, position{x: pos.x, y: pos.y + 1})
					}
				case tile.TopRightBend:
					if pos.prevMove == move.Up {
						innerFields = append(innerFields, position{x: pos.x + 1, y: pos.y}, position{x: pos.x, y: pos.y - 1})
					}
				case tile.TopLeftBend:
					if pos.prevMove == move.Left {
						innerFields = append(innerFields, position{x: pos.x - 1, y: pos.y}, position{x: pos.x, y: pos.y - 1})
					}
				}
			} else {
				switch t {
				case tile.VerticalPipe:
					if pos.prevMove == move.Down {
						innerFields = append(innerFields, position{x: pos.x + 1, y: pos.y})
					} else if pos.prevMove == move.Up {
						innerFields = append(innerFields, position{x: pos.x - 1, y: pos.y})
					}
				case tile.HorizontalPipe:
					if pos.prevMove == move.Right {
						innerFields = append(innerFields, position{x: pos.x, y: pos.y - 1})
					} else if pos.prevMove == move.Left {
						innerFields = append(innerFields, position{x: pos.x, y: pos.y + 1})
					}
				case tile.TopLeftBend:
					if pos.prevMove == move.Up {
						innerFields = append(innerFields, position{x: pos.x - 1, y: pos.y}, position{x: pos.x, y: pos.y - 1})
					}
				case tile.TopRightBend:
					if pos.prevMove == move.Right {
						innerFields = append(innerFields, position{x: pos.x + 1, y: pos.y}, position{x: pos.x, y: pos.y - 1})
					}
				case tile.BottomRightBend:
					if pos.prevMove == move.Down {
						innerFields = append(innerFields, position{x: pos.x + 1, y: pos.y}, position{x: pos.x, y: pos.y + 1})
					}
				case tile.BottomLeftBend:
					if pos.prevMove == move.Left {
						innerFields = append(innerFields, position{x: pos.x - 1, y: pos.y}, position{x: pos.x, y: pos.y + 1})
					}
				}
			}

			for _, innerField := range innerFields {
				count += g.innerLoopBfs(innerField, &checked)
			}
		}

	}
	return count
}

func (g *Grid) innerLoopBfs(start position, checked *[][]bool) int {
	count := 0
	positions := []position{start}
	for len(positions) > 0 {
		currPositions := append([]position{}, positions...)
		positions = make([]position, 0)

		for len(currPositions) > 0 {
			curr := currPositions[0]
			currPositions = currPositions[1:]

			if !g.isValidPosition(curr) || (*checked)[curr.y][curr.x] {
				continue
			}

			(*checked)[curr.y][curr.x] = true
			positions = append(positions, g.allNextPositions(curr)...)
			count++
		}
	}

	return count
}

func (g *Grid) findLoop(start position) (loop, []position, bool) {
	seen := make([][]bool, len(g.grid))
	for i := range g.grid {
		seen[i] = make([]bool, len(g.grid[0]))
	}

	var path []position
	isLoop := g.dfs(start, start, &path, seen)
	if !isLoop {
		return loop{}, path, false
	}

	return loop{
		start:        start,
		path:         path,
		isInnerRight: isInnerRightLoop(path),
	}, path, true
}

func isInnerRightLoop(path []position) bool {
	var leftMoves, rightMoves int
	for i := 0; i < len(path); i++ {
		j := (i + 1) % (len(path) - 1)
		if path[i].prevMove == move.Up && path[j].prevMove == move.Right {
			rightMoves++
		} else if path[i].prevMove == move.Up && path[j].prevMove == move.Left {
			leftMoves++
		} else if path[i].prevMove == move.Right && path[j].prevMove == move.Down {
			rightMoves++
		} else if path[i].prevMove == move.Right && path[j].prevMove == move.Up {
			leftMoves++
		} else if path[i].prevMove == move.Down && path[j].prevMove == move.Left {
			rightMoves++
		} else if path[i].prevMove == move.Down && path[j].prevMove == move.Right {
			leftMoves++
		} else if path[i].prevMove == move.Left && path[j].prevMove == move.Up {
			rightMoves++
		} else if path[i].prevMove == move.Left && path[j].prevMove == move.Down {
			leftMoves++
		}
	}
	return rightMoves > leftMoves
}

func (g *Grid) dfs(curr, start position, path *[]position, seen [][]bool) bool {
	if len(*path) > 2 && curr.x == start.x && curr.y == start.y && g.isValidPrevMove(curr) {
		return true
	} else if !g.isValidPosition(curr) || seen[curr.y][curr.x] || !g.isValidPrevMove(curr) {
		return false
	}

	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	for _, pos := range g.nextValidPositions(curr) {
		if ok := g.dfs(pos, start, path, seen); ok {
			return ok
		}
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func (g *Grid) nextValidPositions(currPos position) []position {
	var moves []move.Move

	switch g.grid[currPos.y][currPos.x] {
	case tile.Start:
		moves = move.AllMoves()
	case tile.VerticalPipe:
		moves = []move.Move{move.Up, move.Down}
	case tile.HorizontalPipe:
		moves = []move.Move{move.Left, move.Right}
	case tile.TopLeftBend:
		moves = []move.Move{move.Right, move.Down}
	case tile.TopRightBend:
		moves = []move.Move{move.Left, move.Down}
	case tile.BottomRightBend:
		moves = []move.Move{move.Left, move.Up}
	case tile.BottomLeftBend:
		moves = []move.Move{move.Right, move.Up}
	}

	return calcNextPositions(currPos, moves)
}

func (g *Grid) allNextPositions(currPos position) []position {
	return calcNextPositions(currPos, move.AllMoves())
}

func calcNextPositions(currPos position, moves []move.Move) []position {
	var result []position
	for _, m := range moves {
		deltaX, deltaY := move.Delta(m)
		result = append(result, position{
			x:        currPos.x + deltaX,
			y:        currPos.y + deltaY,
			prevPos:  &currPos,
			prevMove: m,
		})
	}
	return result
}

func (g *Grid) isValidPrevMove(currPos position) bool {
	currTile := g.grid[currPos.y][currPos.x]
	if currPos.prevPos == nil || currTile == tile.Start {
		return true
	}

	switch currPos.prevMove {
	case move.Up:
		return currTile == tile.VerticalPipe || currTile == tile.TopRightBend || currTile == tile.TopLeftBend
	case move.Right:
		return currTile == tile.HorizontalPipe || currTile == tile.BottomRightBend || currTile == tile.TopRightBend
	case move.Down:
		return currTile == tile.VerticalPipe || currTile == tile.BottomRightBend || currTile == tile.BottomLeftBend
	case move.Left:
		return currTile == tile.HorizontalPipe || currTile == tile.BottomLeftBend || currTile == tile.TopLeftBend
	}
	return false
}
