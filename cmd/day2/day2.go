package day2

import (
	"strconv"

	"hasmanytrees.com/aoc-2024/cmd"
)

var d *cmd.Day

func init() {
	d = cmd.NewDay("day2", star1, star2)
	d.Init()
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
