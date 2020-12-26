package main

import (
	"fmt"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 19

type ruleBook map[string][]string

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

func containsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getPerms(a []string, b []string) []string {
	cnt := 0
	perms := make([]string, len(a)*len(b))
	for i := range a {
		for j := range b {
			perms[cnt] = a[i] + b[j]
			cnt++
		}
	}
	return perms
}

func resolveRule(rulez ruleBook, cur string) []string {
	parts := strings.Split(cur, " | ")
	resolved := make([]string, 0)

	// get all ops
	for _, p := range parts {
		nums := strings.Split(p, " ")
		ops := make([][]string, 0)

		size := 1
		for _, curNum := range nums {
			ops = append(ops, rulez[curNum])
			size *= len(rulez[curNum])
		}

		cnt := 0
		curPerms := make([]string, size)
		for cnt < len(ops)-1 {
			if len(ops[cnt]) > 1 && len(ops[cnt+1]) > 1 || cnt == 0 {
				for i, perm := range getPerms(ops[cnt], ops[cnt+1]) {
					curPerms[i] += perm
				}
			} else {
				for i := range curPerms {
					curPerms[i] += ops[cnt+1][0]
				}
			}

			cnt++
		}
		resolved = append(resolved, curPerms...)
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

	rules := make(ruleBook, 0)
	aRule := strings.Split(inputItems[aIdx], ": ")[0]
	rules[aRule] = []string{"a"}

	bRule := strings.Split(inputItems[bIdx], ": ")[0]
	rules[bRule] = []string{"b"}

	for len(rawRules) > 0 {
		for k, v := range rawRules {
			if ruleResolvable(rules, v) {
				rules[k] = resolveRule(rules, v)
				delete(rawRules, k)
				break
			}
		}
	}
	fmt.Println(rules["0"])

	ruleCnt++
	for ruleCnt < len(inputItems) {
		curLine := inputItems[ruleCnt]
		if containsStr(rules["0"], curLine) {
			answer1++
		}
		ruleCnt++
	}
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
