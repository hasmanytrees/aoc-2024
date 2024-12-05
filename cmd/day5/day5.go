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

func scanNums(s string, sep string) []int {
	nums := []int{}

	fields := strings.Split(s, sep)

	for _, field := range fields {
		v, _ := strconv.Atoi(field)
		nums = append(nums, v)
	}

	return nums
}

func isUpdateInRightOrder(update []int, rules map[int]map[int]bool) bool {
	result := true

	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if !rules[update[i]][update[j]] {
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

func reorderUpdate(update []int, rules map[int]map[int]bool) []int {
	result := []int{}

	// make a map showing all the valid page orders for pages included in the update
	updateMap := map[int]map[int]bool{}
	for i := 0; i < len(update); i++ {
		for j := 0; j < len(update); j++ {
			if i == j {
				continue
			} else {
				if rules[update[i]][update[j]] {
					if _, ok := updateMap[update[i]]; !ok {
						updateMap[update[i]] = map[int]bool{}
					}
					updateMap[update[i]][update[j]] = true
				}
			}
		}
	}

	// now look for a key in this map that doesn't exist in any other value map until there is one update left
	for len(updateMap) > 1 {
		for k1 := range updateMap {
			found := false

			for k2, v := range updateMap {
				if k1 != k2 {
					if _, ok := v[k1]; ok {
						found = true
						continue
					}
				}
			}

			// if we didn't find it then all other remaining pages must be after it so append it to the result
			if !found {
				result = append(result, k1)
			}
		}

		// delete the kay from the update map that we just stored into the result
		delete(updateMap, result[len(result)-1])
	}

	// this last update needs to be appended to the result
	for k1 := range updateMap {
		result = append(result, k1)

		for k2 := range updateMap[k1] {
			result = append(result, k2)
		}
	}

	return result
}

func star1(scanner *bufio.Scanner) {
	result := 0

	rulesMap := map[int]map[int]bool{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 {
			nums := scanNums(line, "|")

			if _, ok := rulesMap[nums[0]]; !ok {
				rulesMap[nums[0]] = map[int]bool{}
			}

			rulesMap[nums[0]][nums[1]] = true
		} else {
			break
		}

	}

	for scanner.Scan() {
		line := scanner.Text()

		update := scanNums(line, ",")

		if isUpdateInRightOrder(update, rulesMap) {
			result += update[len(update)/2]
		}
	}

	fmt.Printf("Result = %v\n", result)
}

func star2(scanner *bufio.Scanner) {
	result := 0

	rules := map[int]map[int]bool{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 0 {
			nums := scanNums(line, "|")

			if _, ok := rules[nums[0]]; !ok {
				rules[nums[0]] = map[int]bool{}
			}

			rules[nums[0]][nums[1]] = true
		} else {
			break
		}

	}

	for scanner.Scan() {
		line := scanner.Text()

		update := scanNums(line, ",")

		if !isUpdateInRightOrder(update, rules) {
			update = reorderUpdate(update, rules)
			result += update[len(update)/2]
		}
	}

	fmt.Printf("Result = %v\n", result)
}
