package day1

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

	similarity := 0

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		nums := scanNums(line)

		left = append(left, nums[0])
		right = append(right, nums[1])
	}

	rightMap := map[int]int{}

	for _, r := range right {
		rightMap[r]++
	}

	for i := 0; i < len(left); i++ {
		similarity += left[i] * rightMap[left[i]]
	}

	fmt.Printf("Similarity = %v\n", similarity)
}
