package day6

import "fmt"

type guard struct {
	direction       direction
	startX, startY  int
	x, y            int
	history         map[string]direction
	newObstructions map[string]bool
}

func newGuard(x, y int, d direction) *guard {
	return &guard{
		direction:       d,
		startX:          x,
		startY:          y,
		x:               x,
		y:               y,
		history:         map[string]direction{},
		newObstructions: map[string]bool{},
	}
}

func (g *guard) clone() *guard {
	historyClone := map[string]direction{}

	for k, v := range g.history {
		historyClone[k] = v
	}

	newObstructionClone := map[string]bool{}

	for k, v := range g.newObstructions {
		newObstructionClone[k] = v
	}

	return &guard{
		direction:       g.direction,
		startX:          g.x,
		startY:          g.y,
		x:               g.x,
		y:               g.y,
		history:         historyClone,
		newObstructions: newObstructionClone,
	}
}

func (g *guard) trackHistory() {
	g.history[fmt.Sprintf("%v,%v", g.x, g.y)] |= g.direction
}

func (g *guard) getHistoryAt(x, y int) (direction, bool) {
	d, ok := g.history[fmt.Sprintf("%v,%v", x, y)]
	return d, ok
}

func (g *guard) hasNewObstructionAt(x, y int) bool {
	return g.newObstructions[fmt.Sprintf("%v,%v", x, y)]
}

func (g *guard) putNewObstructionAt(x, y int) {
	g.newObstructions[fmt.Sprintf("%v,%v", x, y)] = true

}

func (g *guard) next() (x, y int) {
	switch g.direction {
	case Up:
		return g.x, g.y - 1
	case Right:
		return g.x + 1, g.y
	case Down:
		return g.x, g.y + 1
	default:
		return g.x - 1, g.y
	}
}

func (g *guard) move() {
	g.x, g.y = g.next()
}

func (g *guard) nextDirection() direction {
	switch g.direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	default:
		return Up
	}
}

func (g *guard) turn() {
	g.direction = g.nextDirection()
}
