package universe

import (
	"fmt"
	"math"
)

type position struct {
	x int
	y int
}

type positionPair struct {
	a position
	b position
}

func (p position) id() string {
	return fmt.Sprintf("%d|%d", p.x, p.y)
}

func Pairs(positions []position) []positionPair {
	var result []positionPair
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			result = append(result, positionPair{
				a: positions[i],
				b: positions[j],
			})
		}
	}
	return result
}

func (p positionPair) ShortestPath() int {
	deltaX := math.Abs(float64(p.a.x - p.b.x))
	deltaY := math.Abs(float64(p.a.y - p.b.y))
	return int(deltaX + deltaY)
}
