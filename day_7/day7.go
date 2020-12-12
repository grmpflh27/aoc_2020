package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 7

type Bag struct {
	amount int
	color  string
}

func parseBags(bagRules []string) map[string][]Bag {
	// find
	containMap := make(map[string][]Bag)

	for _, curBagRule := range bagRules {
		bagName := strings.TrimSpace(strings.Split(curBagRule, "bags")[0])

		containStr := strings.TrimRight(strings.Split(curBagRule, " contain ")[1], ".")
		var containedBags []Bag
		if !strings.Contains(containStr, "no other") {
			containStrSplit := strings.Split(containStr, ",")
			for _, c := range containStrSplit {
				curWords := strings.Split(strings.TrimSpace(c), " ")

				amnt, _ := strconv.Atoi(curWords[0])
				curBag := Bag{amount: amnt, color: strings.Join(curWords[1:3], " ")}
				containedBags = append(containedBags, curBag)
			}
		}
		containMap[bagName] = containedBags

	}
	return containMap
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsBag(s []Bag, color string) bool {
	for _, a := range s {
		if a.color == color {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	bagRules := aoc2020_shared.LoadString("input7.txt", "\n")

	containMap := parseBags(bagRules)

	// get all where shiny gold is directly within
	var outermostBags []string
	var alreadyChecked []string

	curBagToSearch := "shiny gold"

	for {
		for k, v := range containMap {
			if containsBag(v, curBagToSearch) {
				if !contains(outermostBags, k) {
					outermostBags = append(outermostBags, k)
				}
			}
		}
		alreadyChecked = append(alreadyChecked, curBagToSearch)

		changed := false
		for _, cur := range outermostBags {
			if !contains(alreadyChecked, cur) {
				curBagToSearch = cur
				changed = true
			}
		}

		if !changed {
			break
		}

	}

	fmt.Println(outermostBags)
	answer1 = len(outermostBags)
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	// count contained bags
	var cnt = 0
	//

	bagCntMap := make(map[string]int)

	for k, v := range containMap {
		curSum := 0
		for _, c := range v {
			curSum += c.amount
		}
		bagCntMap[k] = curSum
	}

	fmt.Println(bagCntMap)

	// walk up
	curBag := "shiny gold"
	alreadyChecked2 := make([]Bag, 0)
	toCheck2 := make([]Bag, 0)

	abort := false
	for {
		toCheck2 = append(toCheck2, containMap[curBag]...)
		for _, bag := range toCheck2 {
			cnt += bag.amount * bagCntMap[bag.color]
			nextBags := containMap[bag.color]

			newOne := false
			for _, n := range nextBags {
				if !containsBag(alreadyChecked2, n.color) {
					toCheck2 = append(toCheck2, n)
					newOne = true
					break
				}
			}

			if !newOne {
				abort = true
				break
			}
		}

		if abort {
			break
		}
	}

	// for {

	// }

	answer2 = cnt
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
