package main

import (
	"fmt"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 6

func nextEmptyLineIdx(lines []string, startIdx int) int {
	numLines := len(lines)
	for i := startIdx; i < numLines; i++ {
		if lines[i] == "" {
			return i
		}
	}

	return -1
}

func getYesAnswerCntAny(answers []string) int {
	if len(answers) == 1 {
		return len(answers[0])
	}
	concat := strings.Join(answers[:], "")
	seen := make(map[rune]bool)

	cnt := 0
	for _, cur := range concat {
		if _, ok := seen[cur]; ok {
			continue
		}
		seen[cur] = true
		cnt++
	}

	return cnt
}

func getYesAnswerCntAll(answers []string) int {
	if len(answers) == 1 {
		return len(answers[0])
	}

	seen := make(map[rune]bool)

	cnt := 0
	for _, curAnswer := range answers {
		for _, curRune := range curAnswer {
			if _, ok := seen[curRune]; ok {
				continue
			}
			seen[curRune] = true

			// check if in rune is in all entries - could be optimized
			containedInAll := true
			for _, el := range answers {
				if !strings.ContainsRune(el, curRune) {
					containedInAll = false
					break
				}
			}
			if containedInAll {
				cnt++
			}

		}
	}

	return cnt
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadString("input6.txt", "\n")

	startIdx := 0
	for {
		emptyIdx := nextEmptyLineIdx(inputItems, startIdx)
		if emptyIdx == -1 {
			break
		}
		curGroup := inputItems[startIdx:emptyIdx]
		answer1 += getYesAnswerCntAny(curGroup)
		startIdx = emptyIdx + 1

	}

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	startIdx = 0
	for {
		emptyIdx := nextEmptyLineIdx(inputItems, startIdx)
		if emptyIdx == -1 {
			break
		}
		curGroup := inputItems[startIdx:emptyIdx]
		answer2 += getYesAnswerCntAll(curGroup)
		startIdx = emptyIdx + 1

	}

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
