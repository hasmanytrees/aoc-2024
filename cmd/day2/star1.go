package day2

import (
	"bufio"
	"fmt"
	"os"
)

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

func areLevelsSafe(levels []int) bool {
	prev := levels[0]
	direction := getLevelsDirection(levels)

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

type levelDirection int

const (
	Increasing levelDirection = iota
	Decreasing
)

func getLevelsDirection(levels []int) levelDirection {
	var result levelDirection
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
