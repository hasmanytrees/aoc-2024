package day4

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

	term := "MAS"

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	b := &board{
		lines,
	}

	result := b.findCrosses(term)

	fmt.Printf("Result = %v\n", result)
}
