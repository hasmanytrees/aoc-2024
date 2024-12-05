package dayX

import (
	"bufio"
	"fmt"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("dayX", star1, star2)
}

func star1(scanner *bufio.Scanner) {
	result := 0

	for scanner.Scan() {
		// line := scanner.Text()

	}

	fmt.Printf("Result = %v\n", result)
}

func star2(scanner *bufio.Scanner) {

}
