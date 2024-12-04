package day2

import (
	"bufio"
	"fmt"
	"os"
)

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
