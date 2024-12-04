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
