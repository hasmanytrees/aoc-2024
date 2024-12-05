package day1

import (
	"bufio"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

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

func star1(scanner *bufio.Scanner) {
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

func star2(scanner *bufio.Scanner) {
	similarity := 0

	left := []int{}
	rightMap := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()

		nums := scanNums(line)

		left = append(left, nums[0])

		rightMap[nums[1]]++
	}

	for i := 0; i < len(left); i++ {
		similarity += left[i] * rightMap[left[i]]
	}

	fmt.Printf("Similarity = %v\n", similarity)
}
