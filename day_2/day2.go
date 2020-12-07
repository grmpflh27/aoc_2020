package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 2

func isValid1(line string) bool {
	parts := strings.Split(line, ":")
	password := strings.TrimSpace(parts[1])

	ruleParts := strings.Split(parts[0], " ")
	minMax := strings.Split(ruleParts[0], "-")

	minOcc, _ := strconv.Atoi(minMax[0])
	maxOcc, _ := strconv.Atoi(minMax[1])
	curChar := strings.TrimSpace(ruleParts[1])

	fmt.Printf("%v [%d ; %d] times in %v\n", curChar, minOcc, maxOcc, password)

	occurrences := strings.Count(password, curChar)

	return occurrences >= minOcc && occurrences <= maxOcc

}

func isValid2(line string) bool {
	parts := strings.Split(line, ":")
	password := strings.TrimSpace(parts[1])

	ruleParts := strings.Split(parts[0], " ")
	indices := strings.Split(ruleParts[0], "-")

	index1, _ := strconv.Atoi(indices[0])
	index2, _ := strconv.Atoi(indices[1])
	curChar := strings.TrimSpace(ruleParts[1])

	fmt.Printf("%v at index %d and %d in %v\n", curChar, index1, index2, password)

	pass1 := string(password[index1-1]) == curChar
	pass2 := string(password[index2-1]) == curChar

	return pass1 != pass2

}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadString("input2.txt", "\n")

	var cnt int = 0
	for _, line := range inputItems {
		if isValid1(line) {
			cnt++
		}
	}

	answer1 = cnt
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	cnt = 0
	for _, line := range inputItems {
		if isValid2(line) {
			cnt++
		}
	}

	answer2 = cnt
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
