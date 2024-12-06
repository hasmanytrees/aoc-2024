package day6

import (
	"bufio"
	"fmt"
)

type grid struct {
	width, height int
	obstructions  map[string]bool
}

func newGrid(scanner *bufio.Scanner) (*grid, *guard) {
	g := &grid{
		obstructions: map[string]bool{},
	}

	var guard *guard

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		g.width = len(line)

		for x, r := range line {
			if r == '#' {
				g.putObstructionAt(x, y)
			}
			if r == '^' {
				guard = newGuard(x, y, Up)
			}
		}

		y++
	}

	g.height = y

	return g, guard
}

func (g *grid) isInbounds(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

func (g *grid) hasObstructionAt(x, y int) bool {
	return g.obstructions[fmt.Sprintf("%v,%v", x, y)]
}

func (g *grid) putObstructionAt(x, y int) {
	g.obstructions[fmt.Sprintf("%v,%v", x, y)] = true
}

func (g *grid) move(guard *guard) bool {
	guard.trackHistory()

	if g.hasObstructionAt(guard.next()) {
		guard.turn()
	} else {
		guard.move()
	}

	return g.isInbounds(guard.x, guard.y)
}

func (g *grid) look(guard *guard) bool {
	guard.turn()

	for g.isInbounds(guard.x, guard.y) {
		g.move(guard)

		hd, ok := guard.getHistoryAt(guard.x, guard.y)

		if ok && hd&guard.direction != 0 {
			return true
		}
	}

	return false
}

func (g *grid) print(guard *guard) {
	for y := range g.height {
		for x := range g.width {
			c := "."

			if g.hasObstructionAt(x, y) {
				c = "#"
			}

			if guard.startX == x && guard.startY == y {
				c = "^"
			} else if h, ok := guard.getHistoryAt(x, y); ok {
				if h&(Up|Down) != 0 && h&(Left|Right) != 0 {
					c = "+"
				} else if h&(Up|Down) != 0 {
					c = "|"
				} else if h&(Left|Right) != 0 {
					c = "-"
				}
			}

			if guard.hasNewObstructionAt(x, y) {
				c = "O"
			}

			fmt.Printf("%v", c)
		}

		fmt.Println()
	}
}
