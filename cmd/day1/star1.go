package day1

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

	distance := 0

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		nums := scanNums(line)

		left = sortedInsert(left, nums[0])
		right = sortedInsert(right, nums[1])
	}

	for i := 0; i < len(left); i++ {
		distance += abs(left[i] - right[i])
	}

	fmt.Printf("Distance = %v\n", distance)
}
