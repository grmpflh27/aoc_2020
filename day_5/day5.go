package main

import (
	"fmt"
	"sort"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 5

const rows int = 128
const cols int = 8

func _getId(part string, min int, max int) int {
	targetId := 0
	for _, curChar := range part {
		//fmt.Println(i, min, max)
		if curChar == []rune("F")[0] || curChar == []rune("L")[0] {
			max = max - (max-min+1)/2
		}
		if curChar == []rune("B")[0] || curChar == []rune("R")[0] {
			min = min + (max-min+1)/2
		}
	}
	if part[len(part)-1] == []byte("F")[0] {
		targetId = min
	} else {
		targetId = max
	}
	return targetId
}

func getSeatId(seatCode string) int {
	rowPart := seatCode[:7]
	colPart := seatCode[7:]

	row := _getId(rowPart, 0, 127)
	col := _getId(colPart, 0, 7)
	fmt.Println(row, col)
	return row*8 + col
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadString("input5.txt", "\n")

	collectedIds := make([]int, len(inputItems))
	for _, it := range inputItems {
		cur := getSeatId(it)
		if cur > answer1 {
			answer1 = cur
		}
		collectedIds = append(collectedIds, cur)
		fmt.Println(it, " = ", cur)
	}

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	sort.Ints(collectedIds)

	for i, c := range collectedIds {
		if c == 0 {
			continue
		}
		if collectedIds[i-1] == 0 {
			continue
		}
		if c-collectedIds[i-1] > 1 {
			answer2 = c - 1
			break
		}
	}

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
