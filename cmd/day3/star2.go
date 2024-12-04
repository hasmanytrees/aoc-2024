package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func star2(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

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
