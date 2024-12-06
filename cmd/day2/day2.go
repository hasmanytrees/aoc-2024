package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day2", star1, star2)
}

func scanNums(s string) []int {
	nums := []int{}

	fields := strings.Fields(s)

	for _, field := range fields {
		v, _ := strconv.Atoi(field)
		nums = append(nums, v)
	}

	return nums
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

type direction int

const (
	Increasing direction = iota
	Decreasing
)

func getDirection(x, y int) direction {
	if x < y {
		return Increasing
	}

	return Decreasing
}

func areLevelsSafe(levels []int) bool {
	prevLevel := levels[0]
	prevDirection := getDirection(levels[0], levels[1])

	var currLevel int
	var currDirection direction

	for i := 1; i < len(levels); i++ {
		currLevel = levels[i]
		currDirection = getDirection(prevLevel, currLevel)

		diff := abs(currLevel - prevLevel)

		if currDirection != prevDirection || diff < 1 || diff > 3 {
			return false
		}

		prevLevel = currLevel
	}

	return true
}

func star1(scanner *bufio.Scanner) {
	numSafeReports := 0

	for scanner.Scan() {
		line := scanner.Text()

		levels := scanNums(line)

		if areLevelsSafe(levels) {
			numSafeReports++
		}
	}

	fmt.Printf("# Safe Reports = %v\n", numSafeReports)
}

func star2(scanner *bufio.Scanner) {
	numSafeReports := 0

	for scanner.Scan() {
		line := scanner.Text()

		levels := scanNums(line)

		if areLevelsSafe(levels) {
			numSafeReports++
		} else {
			for i := 0; i < len(levels); i++ {
				newLevels := append(append([]int{}, levels[:i]...), levels[i+1:]...)

				if areLevelsSafe(newLevels) {
					numSafeReports++
					break
				}
			}
		}
	}

	fmt.Printf("# Safe Reports = %v\n", numSafeReports)
}
