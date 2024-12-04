package day1

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
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
