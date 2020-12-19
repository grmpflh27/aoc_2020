package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 16

type attr struct {
	name    string
	allowed []int
}

func parseAttr(attrStr string) attr {
	allowed := make([]int, 0)
	parts := strings.Split(attrStr, ": ")
	name := parts[0]

	allowedRanges := strings.Split(parts[1], " or ")

	for _, r := range allowedRanges {
		min, _ := strconv.Atoi(strings.Split(r, "-")[0])
		max, _ := strconv.Atoi(strings.Split(r, "-")[1])

		for i := min; i <= max; i++ {
			allowed = append(allowed, i)
		}
	}

	return attr{
		name:    name,
		allowed: allowed,
	}
}

func merge(attrs []attr) []int {
	allowedMerged := make([]int, 0)

	for _, cur := range attrs {
		for _, number := range cur.allowed {
			if !contains(allowedMerged, number) {
				allowedMerged = append(allowedMerged, number)
			}
		}
	}
	sort.Ints(allowedMerged)
	return allowedMerged
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(s []int, exclude int) []int {
	tmp := make([]int, 0)

	for _, cur := range s {
		if cur == exclude {
			continue
		}
		tmp = append(tmp, cur)
	}
	return tmp
}

func getInvalidNumbers(lines []string, allowedMerged []int) []int {
	invalid := make([]int, 0)
	for _, line := range lines {
		vals := strings.Split(line, ",")

		for _, v := range vals {
			cur, _ := strconv.Atoi(v)
			if !contains(allowedMerged, cur) {
				invalid = append(invalid, cur)
			}
		}
	}
	return invalid
}

func filterInvalid(lines []string, allowedMerged []int) []string {

	valid := make([]string, 0)
	for _, line := range lines {
		vals := strings.Split(line, ",")

		allValid := true
		for _, v := range vals {
			cur, _ := strconv.Atoi(v)
			if !contains(allowedMerged, cur) {
				allValid = false
				break
			}
		}
		if allValid {
			valid = append(valid, line)
		}
	}

	fmt.Printf("%d out of %d are valid\n", len(valid), len(lines))
	return valid
}

func parseTickets(lines []string) [][]int {
	ticketCategories := make([][]int, len(strings.Split(lines[0], ",")))
	for _, line := range lines {
		vals := strings.Split(line, ",")

		for i, v := range vals {
			cur, _ := strconv.Atoi(v)
			ticketCategories[i] = append(ticketCategories[i], cur)
		}
	}
	return ticketCategories
}

// resolve by removing single Id columns iteratively
func resolve(idxsByClass map[string][]int) map[string]int {
	popped := make(map[string]int, 0)

	listed := make([]map[string][]int, 0)
	for k, v := range idxsByClass {
		curMap := map[string][]int{k: v}
		listed = append(listed, curMap)
	}

	curIdx := 0
	for len(listed) != 0 {
		curItem := listed[curIdx]
		for k, v := range curItem {
			if len(v) == 1 {
				popped[k] = v[0]
				// build new slice without k and v[0]
				tmp := make([]map[string][]int, len(listed)-1)
				cnt := 0
				for i := range listed {
					if i == curIdx {
						continue
					}
					for itemKey, itemIdxs := range listed[i] {
						curMap := map[string][]int{itemKey: remove(itemIdxs, v[0])}
						tmp[cnt] = curMap
						cnt++
					}
				}
				listed = tmp
				// start again from beginning
				curIdx = 0
				break
			}
			curIdx++
		}

	}
	return popped
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	lines := aoc2020_shared.LoadString("input16.txt", "\n")

	attrs := make([]attr, 0)
	curType := 1

	yourTixIdx := 0
	nearbyTixIdx := 0
	pyou := &yourTixIdx
	pnearby := &nearbyTixIdx

	for i, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "your ticket") {
			*pyou = i + 1
			curType++
			continue
		}
		if strings.HasPrefix(line, "nearby tickets") {
			*pnearby = i + 1
			curType++
			continue
		}

		if curType == 1 {
			fmt.Println(line)
			attrs = append(attrs, parseAttr(line))
		}
	}

	// merge
	allowedMerged := merge(attrs)
	invalidNumbers := getInvalidNumbers(lines[nearbyTixIdx:], allowedMerged)

	// sum
	for _, c := range invalidNumbers {
		answer1 += c
	}

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	valid := filterInvalid(lines[nearbyTixIdx:], allowedMerged)
	cols := parseTickets(valid)

	idxsByClass := make(map[string][]int, 0)

	// run over attrs
	for _, attr := range attrs {
		for i, curCol := range cols {
			thisCat := true
			for _, val := range curCol {
				if !contains(attr.allowed, val) {
					thisCat = false
					break
				}
			}
			if thisCat {
				idxsByClass[attr.name] = append(idxsByClass[attr.name], i)
			}
		}
	}

	// remove single entries
	fmt.Println(idxsByClass)

	// resolve
	resolved := resolve(idxsByClass)

	// parse your ticket
	parts := strings.Split(lines[yourTixIdx], ",")
	yourTicket := make([]int, len(parts))
	for i, cur := range parts {
		val, _ := strconv.Atoi(cur)
		yourTicket[i] = val
	}

	answer2 = 1
	for k, v := range resolved {
		if strings.HasPrefix(k, "departure") {
			fmt.Println(k, v, yourTicket[v])
			answer2 *= yourTicket[v]
		}
	}

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
