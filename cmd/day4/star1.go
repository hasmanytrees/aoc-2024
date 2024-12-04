package day4

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

	term := "XMAS"

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	b := &board{
		lines,
	}

	result := b.find(term)

	fmt.Printf("Result = %v\n", result)
}
