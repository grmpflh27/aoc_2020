package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 18

type solveFunc func(string) int

func solvePart1(formula string) int {
	curIdx := 0
	erg := 0
	instrs := strings.Split(formula, " ")
	for curIdx < len(instrs) {
		cur := instrs[curIdx]
		if cur == "" {
			curIdx++
			continue
		}
		curNum, err := strconv.Atoi(cur)
		if err != nil {
			// +*()
			switch cur {
			case "+":
				summand, _ := strconv.Atoi(instrs[curIdx+1])
				erg += summand
				curIdx += 2
			case "*":
				prodant, _ := strconv.Atoi(instrs[curIdx+1])
				erg *= prodant
				curIdx += 2
			}
		} else {
			erg = curNum
			curIdx++
		}
	}
	return erg
}

func solvePart2(formula string) int {
	// + precedence over *
	// resolve all +

	if !strings.Contains(formula, "*") {
		return solvePart1(formula)
	}

	// beginning
	if strings.Index(formula, "+") < strings.Index(formula, "*") {
		multIdx := strings.Index(formula, "*")
		subFormula := formula[:multIdx-1]
		subAnswer := solvePart1(subFormula)
		formula = strconv.Itoa(subAnswer) + formula[multIdx-1:]
	}

	// ending
	if strings.LastIndex(formula, "+") > strings.LastIndex(formula, "*") {
		multIdx := strings.LastIndex(formula, "*")
		subFormula := formula[multIdx+1:]
		subAnswer := solvePart1(subFormula)
		formula = formula[:multIdx+2] + strconv.Itoa(subAnswer)
	}

	// this works well for middle pieces
	for strings.Contains(formula, "+") {
		plusIdx := strings.Index(formula, "+")
		start := strings.LastIndex(formula[:plusIdx], "*")
		end := plusIdx + strings.Index(formula[plusIdx+1:], "*")
		subFormula := formula[start+1 : end]
		subAnswer := solvePart1(subFormula)
		formula = formula[:start] + "* " + strconv.Itoa(subAnswer) + formula[end:]
		start = end
	}

	return solvePart1(formula)
}

func solveFormula(formula string, solve solveFunc) int {
	erg := 0
	for len(strings.Split(formula, " ")) > 1 {
		if strings.Contains(formula, "(") {
			start := strings.LastIndex(formula, "(") + 1
			end := start + strings.Index(formula[start:], ")")
			subFormula := formula[start:end]
			subAnswer := solve(subFormula)
			formula = formula[:start-1] + strconv.Itoa(subAnswer) + formula[end+1:]
		} else {
			erg = solve(formula)
			break
		}
	}
	return erg
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	formulas := aoc2020_shared.LoadString("input18.txt", "\n")

	sum := 0
	for _, curFormula := range formulas {
		sum += solveFormula(curFormula, solvePart1)
	}

	answer1 = sum

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	// advanced math
	sum = 0
	for _, curFormula := range formulas {
		sum += solveFormula(curFormula, solvePart2)
	}
	answer2 = sum

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
