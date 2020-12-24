package main

import (
	"fmt"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 19

type ruleBook map[string][]string

// func parseRule(rulez ruleBook, curRule string) ruleBook {
// 	return rulez
// }

func ruleResolvable(rulez ruleBook, cur string) bool {
	parts := strings.Split(cur, " | ")

	for _, p := range parts {
		nums := strings.Split(p, " ")
		for _, curNum := range nums {
			if _, ok := rulez[curNum]; !ok {
				return false
			}
		}
	}
	return true
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func resolveRule(rulez ruleBook, cur string) []string {
	parts := strings.Split(cur, " | ")
	resolved := make([]string, 0)

	buffer := ""
	for _, p := range parts {
		nums := strings.Split(p, " ")
		perms := make([][]string, 0)
		multiPermIdx := make([]int, 0)
		for i, curNum := range nums {
			curOps := rulez[curNum]
			if len(curOps) == 1 {
				buffer += curOps[0]
			} else {
				multiPermIdx = append(multiPermIdx, i)
				if len(buffer) != 0 {
					for i := range curOps {
						curOps[i] = buffer + curOps[i]
					}
				}
				perms = append(perms, curOps)
				buffer = ""
			}
		}

		// permute
		permIdx := 0
		for permIdx < len(perms)-1 {
			if contains(multiPermIdx, permIdx) {
				permutations := permute(perms[permIdx], perms[permIdx+1])
				resolved = append(resolved, permutations...)
			}
			permIdx++
		}

		if len(buffer) != 0 {
			if len(resolved) != 0 {
				for i := range resolved {
					resolved[i] = resolved[i] + buffer
				}
			} else {
				resolved = append(resolved, buffer)
			}

			buffer = ""
		}
	}

	return resolved
}

func permute(a []string, b []string) []string {
	res := make([]string, 0)
	for i := range a {
		for j := range b {
			res = append(res, a[i]+b[j])
		}
	}
	return res
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadString("input19.txt", "\n")

	ruleCnt := 0
	aIdx := 0
	bIdx := 0

	rawRules := make(map[string]string)
	for {
		line := inputItems[ruleCnt]
		if line == "" {
			break
		}
		if strings.Contains(line, "\"a\"") {
			aIdx = ruleCnt
		} else if strings.Contains(line, "\"b\"") {
			bIdx = ruleCnt
		} else {
			parts := strings.Split(line, ": ")
			rawRules[parts[0]] = parts[1]
		}
		ruleCnt++
	}

	fmt.Println(ruleCnt, aIdx, bIdx)

	rules := make(ruleBook, 0)
	aRule := strings.Split(inputItems[aIdx], ": ")[0]
	rules[aRule] = []string{"a"}

	bRule := strings.Split(inputItems[bIdx], ": ")[0]
	rules[bRule] = []string{"b"}

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	for len(rawRules) > 0 {
		for k, v := range rawRules {
			if ruleResolvable(rules, v) {
				rules[k] = resolveRule(rules, v)
				delete(rawRules, k)
				break
			}
		}
	}
	fmt.Println(rules)

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
