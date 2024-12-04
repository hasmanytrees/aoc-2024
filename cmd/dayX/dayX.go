package dayX

import (
	"bufio"
	"fmt"
	"os"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("dayX", star1, star2)
}

func star1(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// result := 0

	for scanner.Scan() {
		// line := scanner.Text()

	}

	// fmt.Printf("Result = %v\n", result)
}

func star2(inputFile string) {

}
