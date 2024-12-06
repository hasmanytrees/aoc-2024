package day6

import (
	"bufio"
	"fmt"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day6", star1, star2)
}

type direction uint8

const (
	Up direction = 1 << iota
	Down
	Left
	Right
)

func star1(scanner *bufio.Scanner) {
	g, guard := newGrid(scanner)

	cont := true

	for cont {
		cont = g.move(guard)
	}

	g.print(guard)

	fmt.Printf("Result = %v\n", len(guard.history))
}

func star2(scanner *bufio.Scanner) {
	g, guard := newGrid(scanner)

	cont := true

	for cont {
		if g.look(guard.clone()) {
			guard.putNewObstructionAt(guard.next())
		}

		cont = g.move(guard)
	}

	g.print(guard)

	// 2143 incorrect
	fmt.Printf("Result = %v\n", len(guard.newObstructions))
}
