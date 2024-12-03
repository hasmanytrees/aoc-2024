package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func star1(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	distance := 0

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		a := scanner.Text()

		nums := scanNums(a)

		left = Insert(left, nums[0])
		right = Insert(right, nums[1])
	}

	for i := 0; i < len(left); i++ {
		distance += abs(left[i] - right[i])
	}

	fmt.Printf("Distance = %v\n", distance)
}

func scanNums(s string) []int {
	i := 0
	j := 0
	values := []int{0}

	for i = 0; i < len(s); i++ {
		if s[i] == ' ' {
			values = append(values, 0)
			j++

			for ; i < len(s); i++ {
				if s[i] != ' ' {
					break
				}
			}
		}

		v, _ := strconv.Atoi(string(s[i]))
		values[j] = values[j]*10 + v
	}

	return values
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
