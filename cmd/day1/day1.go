package day1

import (
	"hasmanytrees.com/aoc-2024/cmd"
)

var d *cmd.Day

func init() {
	d = cmd.NewDay("day1", star1, star2)
	d.Init()
}
