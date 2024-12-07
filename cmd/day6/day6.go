package day6

import (
	"bufio"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day6", star1, star2)
}

type Point struct {
	x, y int
}

func (p Point) next(d Direction) Point {
	switch d {
	case Up:
		return Point{p.x, p.y - 1}
	case Down:
		return Point{p.x, p.y + 1}
	case Left:
		return Point{p.x - 1, p.y}
	default:
		return Point{p.x + 1, p.y}
	}
}

type Direction uint8

const (
	Up Direction = 1 << iota
	Down
	Left
	Right
)

func (d Direction) turnRight() Direction {
	switch d {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	default:
		return Down
	}
}

type Step struct {
	p Point
	d Direction
}

func parseInput(scanner *bufio.Scanner) *Guard {
	lab := &Lab{
		obstructions: mapset.NewSet[Point](),
	}

	var guard *Guard

	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		lab.width = len(line)

		for x, r := range line {
			if r == '#' {
				lab.obstructions.Add(Point{x, y})
			}
			if r == '^' {
				guard = NewGuard(lab, x, y)
			}
		}

		y++
	}

	lab.height = y

	return guard
}

func printMap(g *Guard, newObstructions mapset.Set[Point]) {
	for y := range g.lab.height {
		for x := range g.lab.width {
			c := "."
			p := Point{x, y}

			if g.lab.obstructions.Contains(p) {
				c = "#"
			}

			if g.steps.Contains(Step{p, Up}) || g.steps.Contains(Step{p, Down}) {
				c = "|"
			}

			if g.steps.Contains(Step{p, Left}) || g.steps.Contains(Step{p, Right}) {
				if c == "|" {
					c = "+"
				} else {
					c = "-"
				}
			}

			if newObstructions.Contains(p) {
				c = "O"
			}

			if g.start.p == p {
				c = "^"
			}

			fmt.Print(c)
		}

		fmt.Println()
	}
}

func star1(scanner *bufio.Scanner) {
	guard := parseInput(scanner)

	points := mapset.NewSet[Point]()

	for ok := true; ok; ok = guard.move() {
		points.Add(guard.current.p)
	}

	printMap(guard, mapset.NewSet[Point]())

	fmt.Printf("Result = %v\n", points.Cardinality())
}

func canMakeLoop(g *Guard, next Point) bool {
	// add a temporary obstruction to test if it makes a loop
	g.lab.obstructions.Add(next)
	defer g.lab.obstructions.Remove(next)

	for g.move() {
		next := g.current.p.next(g.current.d)

		if g.steps.Contains(Step{next, g.current.d}) {
			return true
		}
	}

	return false
}

func star2(scanner *bufio.Scanner) {
	guard := parseInput(scanner)

	points := mapset.NewSet[Point]()
	newObstructions := mapset.NewSet[Point]()

	for ok := true; ok; ok = guard.move() {
		points.Add(guard.current.p)
		next := guard.current.p.next(guard.current.d)

		if !guard.lab.obstructions.Contains(next) && !points.Contains(next) && canMakeLoop(guard.clone(), next) {
			newObstructions.Add(next)
		}
	}

	printMap(guard, newObstructions)

	fmt.Printf("Result = %v\n", newObstructions.Cardinality())
}
