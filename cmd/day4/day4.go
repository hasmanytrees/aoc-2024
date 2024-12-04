package day4

import "hasmanytrees.com/aoc-2024/cmd"

var d *cmd.Day

func init() {
	d = cmd.NewDay("day4", star1, star2)
	d.Init()
}

type board struct {
	lines []string
}

func (b *board) inBounds(x, y int) bool {
	return x >= 0 && x < len(b.lines[0]) && y >= 0 && y < len(b.lines)
}

func (b *board) value(x, y int) byte {
	if !b.inBounds(x, y) {
		return ' '
	}

	return b.lines[y][x]
}

func (b *board) searchN(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && y >= len(term)-1 {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			y--
		}
	}

	return result
}

func (b *board) searchS(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && y <= len(b.lines)-len(term) {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			y++
		}
	}

	return result
}

func (b *board) searchE(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && x <= len(b.lines[0])-len(term) {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			x++
		}
	}

	return result
}

func (b *board) searchW(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && x >= len(term)-1 {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			x--
		}
	}

	return result
}

func (b *board) searchNE(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && x <= len(b.lines[0])-len(term) && y >= len(term)-1 {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			x++
			y--
		}
	}

	return result
}

func (b *board) searchSE(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && x <= len(b.lines[0])-len(term) && y <= len(b.lines)-len(term) {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			x++
			y++
		}
	}

	return result
}

func (b *board) searchSW(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && x >= len(term)-1 && y <= len(b.lines)-len(term) {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			x--
			y++
		}
	}

	return result
}

func (b *board) searchNW(term string, x, y int) bool {
	result := false

	if b.inBounds(x, y) && x >= len(term)-1 && y >= len(term)-1 {
		result = true

		for i := 0; i < len(term); i++ {
			if term[i] != b.value(x, y) {
				result = false
				break
			}

			x--
			y--
		}
	}

	return result
}

func (b *board) find(term string) int {
	result := 0

	for y := range b.lines {
		for x := range b.lines[y] {
			if b.searchN(term, x, y) {
				result++
			}
			if b.searchNE(term, x, y) {
				result++
			}
			if b.searchE(term, x, y) {
				result++
			}
			if b.searchSE(term, x, y) {
				result++
			}
			if b.searchS(term, x, y) {
				result++
			}
			if b.searchSW(term, x, y) {
				result++
			}
			if b.searchW(term, x, y) {
				result++
			}
			if b.searchNW(term, x, y) {
				result++
			}
		}
	}

	return result
}