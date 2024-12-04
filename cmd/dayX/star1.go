package dayX

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

	for scanner.Scan() {
		// line := scanner.Text()

	}

	// fmt.Printf("Result = %v\n", result)
}
