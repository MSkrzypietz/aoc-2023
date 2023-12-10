package tile

import "fmt"

type Tile int

const (
	Start Tile = iota
	Ground
	VerticalPipe
	HorizontalPipe
	TopLeftBend
	TopRightBend
	BottomRightBend
	BottomLeftBend
)

func NewTile(ch int32) Tile {
	switch ch {
	case 'S':
		return Start
	case '.':
		return Ground
	case '|':
		return VerticalPipe
	case '-':
		return HorizontalPipe
	case 'F':
		return TopLeftBend
	case '7':
		return TopRightBend
	case 'J':
		return BottomRightBend
	case 'L':
		return BottomLeftBend
	default:
		panic(fmt.Sprintf("Unknown tile: %v", ch))
	}
}
