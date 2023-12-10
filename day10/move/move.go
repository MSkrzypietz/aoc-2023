package move

import "fmt"

type Move int

const (
	Up Move = iota
	Right
	Down
	Left
)

func AllMoves() []Move {
	return []Move{Up, Right, Down, Left}
}

func Delta(m Move) (int, int) {
	switch m {
	case Up:
		return 0, -1
	case Right:
		return 1, 0
	case Down:
		return 0, 1
	case Left:
		return -1, 0
	default:
		panic(fmt.Sprintf("Unknown move: %v", m))
	}
}
