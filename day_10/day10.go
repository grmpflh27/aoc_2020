package main

import (
	"fmt"
	"sort"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 10

func intMax(array []int) int {
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func getBuiltInJoltageAdapter(adapters []int) int {
	return intMax(adapters) + 3
}

func getWithinLimit(curJoltage int, outputJoltage []int) []int {
	found := make([]int, 0)
	for _, cur := range outputJoltage {
		if cur > curJoltage && cur <= curJoltage+3 {
			found = append(found, cur)
		}
	}
	return found
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	outputJoltage := aoc2020_shared.LoadInt("input10.txt", "\n")
	sort.Ints(outputJoltage)

	builtin := getBuiltInJoltageAdapter(outputJoltage)
	outputJoltage = append(outputJoltage, builtin)
	fmt.Println(outputJoltage)

	curJolt := 0
	differences := make(map[int]int)
	for curJolt < builtin {
		nexts := getWithinLimit(curJolt, outputJoltage)
		diff0 := nexts[0] - curJolt
		differences[diff0]++
		curJolt = nexts[0]
	}

	answer1 = differences[1] * differences[3]

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	curJolt = 0
	options := 0
	for curJolt < builtin {
		nexts := getWithinLimit(curJolt, outputJoltage)

		if len(nexts) > 1 {
			fmt.Println(nexts)
			options += len(nexts) + 1
		}
		curJolt = nexts[0]
	}

	answer2 = options
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
