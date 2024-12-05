package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day5", star1, star2)
}

func scanNums(s string) []int {
	nums := []int{}

	fields := strings.Split(s, ",")

	for _, field := range fields {
		v, _ := strconv.Atoi(field)
		nums = append(nums, v)
	}

	return nums
}

func isUpdateInRightOrder(update []int, rules map[string]bool) bool {
	result := true

	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if !rules[fmt.Sprintf("%v|%v", update[i], update[j])] {
				result = false
				break
			}
		}

		if !result {
			break
		}
	}

	return result
}

func star1(scanner *bufio.Scanner) {
	result := 0

	rules := map[string]bool{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 {
			rules[line] = true
		} else {
			break
		}

	}

	for scanner.Scan() {
		line := scanner.Text()

		update := scanNums(line)

		if isUpdateInRightOrder(update, rules) {
			result += update[len(update)/2]
		}
	}

	fmt.Printf("Result = %v\n", result)
}

func star2(scanner *bufio.Scanner) {

}
