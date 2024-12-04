package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func star1(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

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
