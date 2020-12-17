package main

import (
	"fmt"
	"strconv"
	"strings"
)

const day int = 15

func wasFirstTimeSpoken(numSpoken []int, num int) bool {
	cnt := 0
	for _, n := range numSpoken {
		if n == num {
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}
	return true
}

func getLastIdxof(numSpoken []int, num int) int {
	for i := len(numSpoken) - 1; i >= 0; i-- {
		if numSpoken[i] == num {
			return i + 1
		}
	}
	return -1
}

func ParseIntMap(s, sep string) map[int]int {
	m := make(map[int]int)
	for i, line := range strings.Split(s, sep) {
		val, _ := strconv.Atoi(line)
		m[i] = val
	}
	return m
}

func main() {
	fmt.Printf("Day %d\n", day)

	// idiomatic
	var answer1 int
	var answer2 int

	numSpoken := []int{0, 3, 1, 6, 7, 5}
	curNum := numSpoken[2]

	numOfInterest := 2020
	curRound := len(numSpoken)

	for curRound < numOfInterest {
		if wasFirstTimeSpoken(numSpoken, curNum) {
			curNum = 0
		} else {
			lastIdx := getLastIdxof(numSpoken[:len(numSpoken)-1], curNum)
			curNum = curRound - lastIdx
		}

		numSpoken = append(numSpoken, curNum)
		curRound++
	}

	answer1 = curNum

	// better approach
	list := ParseIntMap(`0,3,1,6,7,5`, ",")
	usageIndices := make(map[int][]int)
	recently := 0
	for i := 0; i < 30000000; i++ {
		v, exist := list[i]
		if !exist {
			lastIndex, exist := usageIndices[recently]
			if exist && len(lastIndex) > 1 {
				v = lastIndex[len(lastIndex)-1] - lastIndex[len(lastIndex)-2]
			} else {
				v = 0
			}
		}

		recently = v
		usageIndices[v] = append(usageIndices[v], i)
		if i == 2019 {
			answer1 = v
		}
	}
	answer2 = recently

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
