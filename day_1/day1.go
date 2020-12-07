package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 1

const target int = 2020

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var answer1 int
	var answer2 int
	eReportItems := aoc2020_shared.LoadInt("input1.txt", "\n")

	// 2 product
	fmt.Printf("Day %d\n", day)

	for _, item := range eReportItems {
		rest := target - item
		if contains(eReportItems, rest) {
			fmt.Printf("%d & %d\n", item, rest)
			answer1 = rest * item
			break

		}
	}

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	// 3 product
	var answerFound bool = false
	for idx, item := range eReportItems {
		rest := target - item

		for otherIdx, otherItem := range eReportItems {
			if otherIdx == idx {
				continue
			}
			otherRest := rest - otherItem
			if contains(eReportItems, otherRest) {
				fmt.Printf("%d & %d & %d\n", item, otherItem, otherRest)
				answer2 = item * otherItem * otherRest
				answerFound = true
				break
			}

		}
		if answerFound {
			break
		}

	}

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
