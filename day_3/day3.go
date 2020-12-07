package main

import (
	"fmt"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 3

type coord struct {
	row int
	col int
}

func isTree(forest []string, c coord) bool {
	return forest[c.row][c.col] == []byte("#")[0]
}

func treeCount(forest []string, rowInc int, colInc int) int {
	curPos := coord{0, 0}

	numSteps := len(forest) / rowInc
	numCols := len(forest[0])

	stepCnt := 0
	trees := 0

	for stepCnt < numSteps-1 {
		curPos.row += rowInc
		curPos.col += colInc

		if curPos.col >= numCols {
			curPos.col -= numCols
		}
		fmt.Printf("Step %d out of %d: %v\n", stepCnt+1, numSteps, curPos)
		if isTree(forest, curPos) {
			trees++
		}
		stepCnt++
	}
	return trees
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int

	forest := aoc2020_shared.LoadString("input3.txt", "\n")

	answer1 = treeCount(forest, 1, 3)
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	var rowIncs = []int{1, 1, 1, 1, 2}
	var colIncs = []int{1, 3, 5, 7, 1}

	treeProd := 0
	for i := 0; i < len(rowIncs); i++ {
		curCnt := treeCount(forest, rowIncs[i], colIncs[i])
		if i == 0 {
			treeProd = curCnt
		} else {
			treeProd *= curCnt
		}
	}

	answer2 = treeProd
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
