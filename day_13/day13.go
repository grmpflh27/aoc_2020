package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 13

type constraint struct {
	busId  int
	offset int
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadString("input13.txt", "\n")

	busStrs := strings.Split(inputItems[0], ",")

	busIds := make([]int, 0)
	for _, cur := range busStrs {
		if val, err := strconv.Atoi(cur); err == nil {
			busIds = append(busIds, val)
		}
	}

	earliestTimestamp := 1001796

	minWait := 1001796
	minBusId := -1
	for _, busId := range busIds {
		mult := earliestTimestamp/busId + 1
		diff := mult*busId - earliestTimestamp
		//fmt.Printf("%d => %d diff\n", busId, diff)
		if diff < minWait {
			minWait = diff
			minBusId = busId
		}
	}
	answer1 = minWait * minBusId

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	busStrs = strings.Split(inputItems[0], ",")

	vals := make([]int, 0)
	for i, cur := range busStrs {
		if val, err := strconv.Atoi(cur); err == nil {
			vals = append(vals, val+i)
		}
	}

	fmt.Println(LCM(vals[0], vals[1], vals[2:]...))

	// idiomatic - won't scale
	// for {
	// 	allGood := true
	// 	for _, cur := range constraints {
	// 		if (t+cur.offset)%cur.busId != 0 {
	// 			allGood = false
	// 			break
	// 		}
	// 	}

	// 	if allGood {
	// 		break
	// 	}
	// 	t++
	// }

	// answer2 = t

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
