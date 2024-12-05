package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day3", star1, star2)
}

func star1(scanner *bufio.Scanner) {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindAllSubmatch([]byte(line), -1)

		for _, match := range matches {
			x, _ := strconv.Atoi(string(match[1]))
			y, _ := strconv.Atoi(string(match[2]))

			result += x * y
		}
	}

	fmt.Printf("Result = %v\n", result)
}

func star2(scanner *bufio.Scanner) {
	r, _ := regexp.Compile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)

	result := 0

	enabled := true

	for scanner.Scan() {
		line := scanner.Text()

		matches := r.FindAllSubmatch([]byte(line), -1)

		for _, match := range matches {
			command := string(match[0])

			if command == "do()" {
				enabled = true
				continue
			}

			if command == "don't()" {
				enabled = false
				continue
			}

			if enabled {
				x, _ := strconv.Atoi(string(match[2]))
				y, _ := strconv.Atoi(string(match[3]))

				result += x * y
			}
		}
	}

	fmt.Printf("Result = %v\n", result)
}
