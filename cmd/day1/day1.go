package day1

import (
	"cmp"
	"slices"

	"hasmanytrees.com/aoc-2024/cmd"
)

var d *cmd.Day

func init() {
	d = cmd.NewDay("day1", star1, star2)
	d.Init()
}

func Insert[T cmp.Ordered](ts []T, t T) []T {
	// find the slot to insert the new value
	i, _ := slices.BinarySearch(ts, t)

	return slices.Insert(ts, i, t)
}
