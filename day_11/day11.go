package main

import (
	"bytes"
	"fmt"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 11

func printSeatChart(seatChart [][]byte) {
	for _, row := range seatChart {
		fmt.Println(string(row))
	}
	fmt.Println("----------")
}

type coord struct {
	row int
	col int
}

func getNeightborIdx(seatChart [][]byte, row int, col int) []coord {
	// get 8 neighboring idx pairse

	// A B C
	// E X F
	// G H I
	neighborIdxs := []coord{
		coord{row - 1, col - 1},
		coord{row - 1, col},
		coord{row - 1, col + 1},
		coord{row, col - 1},
		coord{row, col + 1},
		coord{row + 1, col - 1},
		coord{row + 1, col},
		coord{row + 1, col + 1},
	}

	// filter for outOfBounds
	filteredIdxs := make([]coord, 0)

	for _, cur := range neighborIdxs {
		if cur.col < 0 || cur.col >= len(seatChart[0]) || cur.row < 0 || cur.row >= len(seatChart) {
			continue
		}
		filteredIdxs = append(filteredIdxs, cur)
	}
	return filteredIdxs
}

func cntNeighboringOccSeatsPart1(seatChart [][]byte, row int, col int) int {

	filteredIdxs := getNeightborIdx(seatChart, row, col)

	cnt := 0
	for _, coord := range filteredIdxs {
		seat := seatChart[coord.row][coord.col]
		if seat == []byte("#")[0] {
			cnt++
			continue
		}
	}
	return cnt
}

func cntNeighboringOccSeatsPart2(seatChart [][]byte, row int, col int) int {

	filteredIdxs := getNeightborIdx(seatChart, row, col)

	cnt := 0
	for _, coord := range filteredIdxs {
		seat := seatChart[coord.row][coord.col]
		if seat == []byte("#")[0] {
			cnt++
			continue
		}
		if seat == []byte(".")[0] {
			distRow := coord.row - row
			distCol := coord.col - col
			mult := 2
			for {
				nextRow := (row + mult*distRow)
				nextCol := (col + mult*distCol)
				if nextCol < 0 || nextCol >= len(seatChart[0]) || nextRow < 0 || nextRow >= len(seatChart) {
					break
				}
				seat = seatChart[nextRow][nextCol]
				if seat == []byte(".")[0] {
					mult++
					continue
				} else {
					if seat == []byte("#")[0] {
						cnt++
					}
					break
				}
			}
		}
	}
	return cnt
}

func countSeats(newSeatChart [][]byte) int {
	seatCnt := 0
	for i := range newSeatChart {
		seatCnt += bytes.Count(newSeatChart[i], []byte("#"))
	}
	return seatCnt
}

func checkAllEqual(seatChart [][]byte, newSeatChart [][]byte) bool {
	allEqual := true
	for i := range seatChart {
		if bytes.Compare(seatChart[i], newSeatChart[i]) != 0 {
			allEqual = false
			break
		}
	}
	return allEqual
}

func main() {
	fmt.Printf("Day %d\n", day)

	fmt.Printf("# = %d\n", []byte("#")[0])
	fmt.Printf("L = %d\n", []byte("L")[0])
	fmt.Printf(". = %d\n\n", []byte(".")[0])

	var answer1 int
	var answer2 int
	seatRows := aoc2020_shared.LoadString("input11.txt", "\n")
	rows := len(seatRows)
	seatChart := make([][]byte, rows)
	newSeatChart := make([][]byte, rows)
	for i, row := range seatRows {
		seatChart[i] = []byte(row)
		newSeatChart[i] = []byte(row)
	}
	cols := len(seatChart[0])

	rounds := 1
	for {
		for rowId := 0; rowId < rows; rowId++ {
			for colId := 0; colId < cols; colId++ {
				seat := seatChart[rowId][colId]
				cnt := cntNeighboringOccSeatsPart1(seatChart, rowId, colId)

				// fill seat
				if seat == []byte("L")[0] && cnt == 0 {
					newSeatChart[rowId][colId] = []byte("#")[0]
					// empty seat
				} else if seat == []byte("#")[0] && cnt >= 4 {
					newSeatChart[rowId][colId] = []byte("L")[0]
				}
			}
		}

		if checkAllEqual(seatChart, newSeatChart) {
			fmt.Printf("converged in round %d\n", rounds)
			break
		}

		// copy
		for i := range seatRows {
			_ = copy(seatChart[i], newSeatChart[i])
		}
		rounds++
	}

	// count seats

	answer1 = countSeats(newSeatChart)
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	// -------
	// part 2
	// -------

	// reset
	seatRows = aoc2020_shared.LoadString("input11.txt", "\n")
	for i, row := range seatRows {
		seatChart[i] = []byte(row)
		newSeatChart[i] = []byte(row)
	}

	rounds = 1
	for {
		for rowId := 0; rowId < rows; rowId++ {
			for colId := 0; colId < cols; colId++ {
				seat := seatChart[rowId][colId]
				cnt := cntNeighboringOccSeatsPart2(seatChart, rowId, colId)

				// fill seat
				if seat == []byte("L")[0] && cnt == 0 {
					newSeatChart[rowId][colId] = []byte("#")[0]
					// empty seat
				} else if seat == []byte("#")[0] && cnt >= 5 {
					newSeatChart[rowId][colId] = []byte("L")[0]
				}
			}
		}

		// check if is all equal
		if checkAllEqual(seatChart, newSeatChart) {
			fmt.Printf("converged in round %d\n", rounds)
			break
		}

		// copy
		for i := range seatRows {
			_ = copy(seatChart[i], newSeatChart[i])
		}
		//fmt.Println("After round ", rounds)
		//printSeatChart(newSeatChart)

		rounds++
	}

	answer2 = countSeats(newSeatChart)

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
