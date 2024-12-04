package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day2", star1, star2)
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

type direction int

const (
	Increasing direction = iota
	Decreasing
)

func getDirection(levels []int) direction {
	var result direction
	prev := levels[0]

	for i := 1; i < len(levels); i++ {
		curr := levels[i]

		if prev == curr {
			continue
		}

		if curr-prev > 0 {
			result = Increasing
			break
		} else {
			result = Decreasing
			break
		}
	}

	return result
}

func areLevelsSafe(levels []int) bool {
	prev := levels[0]
	direction := getDirection(levels)

	for i := 1; i < len(levels); i++ {
		curr := levels[i]
		diff := curr - prev

		if diff == 0 {
			return false
		} else {
			if diff > 0 && direction != Increasing {
				return false
			}
			if diff < 0 && direction != Decreasing {
				return false
			}
		}

		if abs(diff) > 3 {
			return false
		}

		prev = curr
	}

	return true
}

func star1(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

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

func star2(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

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
