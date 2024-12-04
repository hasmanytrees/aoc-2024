package dayX

import "hasmanytrees.com/aoc-2024/cmd"

var d *cmd.Day

func init() {
	d = cmd.NewDay("dayX", star1, star2)
	d.Init()
}
