package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"hasmanytrees.com/aoc-2024/cmd"
)

func init() {
	cmd.AddDay("day7", star1, star2)
}

func scanNums(s string) []int {
	nums := []int{}

	fields := strings.Fields(s)

	for _, field := range fields {
		v, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, v)
	}

	return nums
}

func permutateOperations(operators []string, l int) [][]string {
	prefixes := [][]string{}

	for _, o := range operators {
		prefixes = append(prefixes, []string{o})
	}

	results := prefixes

	for range l - 1 {
		results = [][]string{}

		for _, o := range operators {
			for _, p := range prefixes {
				results = append(results, append(p, o))
			}
		}

		prefixes = results
	}

	return results
}

func evaluate(a, b int, o string) int {
	switch o {
	case "+":
		return a + b
	case "*":
		return a * b
	case "||":
		return a*int(math.Pow10(len(strconv.Itoa(b)))) + b
	}

	return 0
}

func isEquationTrue(testValue int, nums []int, operators [][]string) bool {
	for _, ops := range operators {
		v := nums[0]

		// fmt.Printf("\t%v %v\n", nums, ops)

		for i, o := range ops {
			v = evaluate(v, nums[i+1], o)

			if v > testValue {
				break
			}
		}

		if v == testValue {
			// fmt.Printf("%v = %v %v\n", testValue, nums, ops)
			return true
		}
	}

	return false
}

type PermCache struct {
	operators    []string
	permutations map[int][][]string
}

func NewPermCache(operators []string) *PermCache {
	return &PermCache{
		operators:    operators,
		permutations: map[int][][]string{},
	}
}

func (p *PermCache) getPermutations(l int) [][]string {
	if _, ok := p.permutations[l]; !ok {
		p.permutations[l] = permutateOperations(p.operators, l)
	}

	return p.permutations[l]
}

func solveWithOperators(scanner *bufio.Scanner, operators []string) int {
	result := 0

	permCache := NewPermCache(operators)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")

		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		nums := scanNums(parts[1])

		perms := permCache.getPermutations(len(nums) - 1)

		if isEquationTrue(testValue, nums, perms) {
			result += testValue
		}
	}

	return result
}

func star1(scanner *bufio.Scanner) {
	result := solveWithOperators(scanner, strings.Fields("+ *"))

	fmt.Printf("Result = %v\n", result)
}

func star2(scanner *bufio.Scanner) {
	result := solveWithOperators(scanner, strings.Fields("+ * ||"))

	// 12676092111112 too low
	// 9223372036854775807 max
	fmt.Printf("Result = %v\n", result)
}
