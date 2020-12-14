package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 8

type instruction struct {
	op  string
	arg int
}

func parse(rule string) instruction {
	parts := strings.Split(rule, " ")

	arg, _ := strconv.Atoi(parts[1])
	return instruction{op: parts[0], arg: arg}
}

func runTrough(instrs []instruction) (int, map[int]int) {
	accum := 0

	curIdx := 0
	visitCnt := make(map[int]int)
	for {
		if curIdx == len(instrs) {
			return accum, visitCnt
		}
		instr := instrs[curIdx]
		visitCnt[curIdx]++

		if visitCnt[curIdx] > 1 {
			return accum, visitCnt
		}

		switch instr.op {
		case "acc":
			accum += instr.arg
			curIdx++
		case "jmp":
			curIdx += instr.arg
		case "nop":
			curIdx++
		}
	}
}

func anyCntUnequal1(visitCnt map[int]int) bool {
	for _, v := range visitCnt {
		if v != 1 {
			return true
		}
	}
	return false
}

func iterFinishes(idxs []int, from string, to string, instrs []instruction) (bool, int) {
	finished := false
	answer := -1
	for _, idx := range idxs {
		// swap
		instrs[idx].op = to

		accum, visitCnt := runTrough(instrs)
		if !anyCntUnequal1(visitCnt) {
			finished = true
			answer = accum
			break
		}

		// swap back
		instrs[idx].op = from
	}
	return finished, answer
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer2 int
	inputItems := aoc2020_shared.LoadString("input8.txt", "\n")

	instrs := make([]instruction, len(inputItems))

	for i, cur := range inputItems {
		instrs[i] = parse(cur)
	}

	answer1, _ := runTrough(instrs)

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	// either a jmp is supposed to be a nop, or a nop is supposed to be a jmp
	nop_idxs := make([]int, 0)
	jmp_idxs := make([]int, 0)

	for i, cur := range instrs {
		if cur.op == "nop" {
			nop_idxs = append(nop_idxs, i)
		} else if cur.op == "jmp" {
			jmp_idxs = append(jmp_idxs, i)
		}
	}

	finished, ans := iterFinishes(nop_idxs, "nop", "jmp", instrs)

	if !finished {
		finished, ans = iterFinishes(jmp_idxs, "jmp", "nop", instrs)
	}

	answer2 = ans
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")
}
