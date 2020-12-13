package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 9

func getSmaller(preamble []int, than int) []int {
	smallerThan := make([]int, 0)
	for _, cur := range preamble {
		if cur < than {
			smallerThan = append(smallerThan, cur)
		}
	}
	return smallerThan
}

func subtractNext(preamble []int, sub int) []int {
	tmp := make([]int, len(preamble))
	for i := range preamble {
		tmp[i] = sub - preamble[i]
	}
	return tmp
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func hasIntersect(a []int, b []int) bool {
	for _, cur := range a {
		if contains(b, cur) {
			return true
		}
	}
	return false
}

func isValidNext(preamble []int, next int) bool {
	smallerInPreamble := getSmaller(preamble, next)
	subtracted := subtractNext(smallerInPreamble, next)
	return hasIntersect(smallerInPreamble, subtracted)

}

func contiguousSumMatches(searchSpace []int, val int) bool {
	sum := 0
	idx := 0
	for sum < val {
		sum += searchSpace[idx]
		idx++
	}
	return sum == val
}

func intMinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadInt("input9.txt", "\n")

	curStart := 0
	curEnd := 25

	for {
		section := inputItems[curStart:curEnd]
		if !isValidNext(section, inputItems[curEnd]) {
			break
		}
		curStart++
		curEnd++
	}
	answer1 = inputItems[curEnd]

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	curStart = 0
	searchSpace := inputItems[curStart:curEnd]

	for {
		if contiguousSumMatches(searchSpace, answer1) {
			break
		}
		searchSpace = searchSpace[1:]
	}

	chunk := []int{searchSpace[0]}

	sum := searchSpace[0]
	idx := 1
	for sum < answer1 {
		sum += searchSpace[idx]
		chunk = append(chunk, searchSpace[idx])
		idx++
	}
	min, max := intMinMax(chunk)
	answer2 = min + max

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
