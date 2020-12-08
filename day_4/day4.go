package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 4

func nextEmptyLineIdx(lines []string, startIdx int) int {
	numLines := len(lines)
	for i := startIdx; i < numLines; i++ {
		if lines[i] == "" {
			return i
		}
	}

	return -1
}

var required = []string{
	"byr", //(Birth Year)
	"iyr", //(Issue Year)
	"eyr", //(Expiration Year)
	"hgt", //(Height)
	"hcl", //(Hair Color)
	"ecl", //(Eye Color)
	"pid", //(Passport ID)
	//"cid", //(Country ID)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isValid(cur string, part2 bool) bool {
	parts := strings.Split(cur, " ")

	// init all false map
	requiredMap := make(map[string]bool)
	for _, req := range required {
		requiredMap[req] = false
	}

	for _, p := range parts {
		curKey := strings.TrimSpace(strings.Split(p, ":")[0])
		if _, ok := requiredMap[curKey]; ok {
			requiredMap[curKey] = true
		}
	}

	for _, v := range requiredMap {
		if !v {
			return false
		}
	}

	// extra validations
	if part2 {
		for _, p := range parts {
			curKey := strings.TrimSpace(strings.Split(p, ":")[0])
			curValue := strings.TrimSpace(strings.Split(p, ":")[1])

			switch curKey {
			case "byr":
				if !isFourDigitBetween(curValue, 1920, 2002) {
					return false
				}
			case "iyr":
				if !isFourDigitBetween(curValue, 2010, 2020) {
					return false
				}
			case "eyr":
				if !isFourDigitBetween(curValue, 2020, 2030) {
					return false
				}
			case "hgt":
				unit := curValue[len(curValue)-2:]
				size, err := strconv.Atoi(curValue[:len(curValue)-2])
				if err != nil {
					return false
				}

				if unit == "cm" {
					if size < 150 || size > 193 {
						return false
					}
				} else if unit == "in" {
					if size < 59 || size > 76 {
						return false
					}
				} else {
					return false
				}
			case "hcl":
				if curValue[0] != []byte("#")[0] {
					return false
				}
				if len(curValue[1:]) != 6 {
					return false
				}
				_, err := strconv.ParseUint(curValue[1:], 16, 64)
				if err != nil {
					return false
				}
			case "ecl":
				if len(curValue) != 3 {
					return false
				}
				allowed := []string{
					"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
				}
				if !contains(allowed, curValue) {
					return false
				}
			case "pid":
				if len(curValue) != 9 {
					return false
				}
				stripped := strings.TrimLeft(curValue, "0")
				_, err := strconv.Atoi(stripped)
				if err != nil {
					return false
				}

			default:
				continue
			}
		}

	}

	return true
}

func isFourDigitBetween(curValue string, min int, max int) bool {
	if len(curValue) != 4 {
		return false
	}
	intValue, err := strconv.Atoi(curValue)
	if err != nil {
		return false
	}
	if intValue < min || intValue > max {
		return false
	}
	return true
}

func getValidPassports(lines []string, part2 bool) int {
	startIdx := 0
	endIdx := 0

	var abort = false
	var passport string
	var cnt = 0
	for {
		endIdx = nextEmptyLineIdx(lines, startIdx)
		if endIdx == -1 {
			passport = strings.Join(lines[startIdx:], " ")
			abort = true
		} else {
			passport = strings.Join(lines[startIdx:endIdx], " ")
		}

		if isValid(passport, part2) {
			cnt++
		}

		if abort {
			break
		}

		startIdx = endIdx + 1
	}

	return cnt
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	lines := aoc2020_shared.LoadString("input4.txt", "\n")

	answer1 = getValidPassports(lines, false)

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	answer2 = getValidPassports(lines, true)

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
