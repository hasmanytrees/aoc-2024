package day6

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Guard struct {
	lab     *Lab
	start   Step
	current Step
	steps   mapset.Set[Step]
}

func NewGuard(lab *Lab, x, y int) *Guard {
	p := Point{x, y}
	s := Step{p, Up}

	steps := mapset.NewSet[Step]()
	steps.Add(s)

	return &Guard{
		lab:     lab,
		start:   s,
		current: s,
		steps:   steps,
	}
}

func (g *Guard) clone() *Guard {
	return &Guard{
		lab:     g.lab,
		start:   g.start,
		current: g.current,
		steps:   g.steps.Clone(),
	}
}

func (g *Guard) inside() bool {
	x := g.current.p.x
	y := g.current.p.y

	return x >= 0 && x < g.lab.width && y >= 0 && y < g.lab.height
}

func (g *Guard) move() bool {
	next := g.current.p.next(g.current.d)

	if g.lab.obstructions.Contains(next) {
		g.current.d = g.current.d.turnRight()
	} else {
		g.current.p = next
	}

	inside := g.inside()

	if inside {
		g.steps.Add(g.current)
	}

	return inside
}
