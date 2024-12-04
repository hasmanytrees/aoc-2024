package day1

import (
	"cmp"
	"slices"
	"strconv"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day1", star1, star2)
}

func sortedInsert[T cmp.Ordered](ts []T, t T) []T {
	// find the slot to insert the new value
	i, _ := slices.BinarySearch(ts, t)

	return slices.Insert(ts, i, t)
}

func scanNums(s string) []int {
	i := 0
	j := 0
	values := []int{0}

	for i = 0; i < len(s); i++ {
		if s[i] == ' ' {
			values = append(values, 0)
			j++

			for ; i < len(s); i++ {
				if s[i] != ' ' {
					break
				}
			}
		}

		v, _ := strconv.Atoi(string(s[i]))
		values[j] = values[j]*10 + v
	}

	return values
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
