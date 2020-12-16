package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 14
const masklength int = 36

func convertToBin(num int, length int) string {
	s := ""

	// num /= 2 divides num by 2 for each iteration and assigns the result to num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() casts numbers to strings
		s = strconv.Itoa(lsb) + s
	}

	// zeropad to length 36
	pad := strings.Repeat("0", length-len(s))
	return pad + s
}

type writeOp struct {
	addr int
	val  int
}

func parseToWriteOp(rule string) writeOp {
	parts := strings.Split(rule, " = ")
	val, _ := strconv.Atoi(parts[1])
	addrStr := strings.Split(strings.Split(parts[0], "[")[1], "]")[0]
	addr, _ := strconv.Atoi(addrStr)

	return writeOp{addr: addr, val: val}
}

func applyMask(value int, mask string) string {
	val := convertToBin(value, masklength)

	s := ""
	for i := range mask {
		if mask[i] == []byte("X")[0] {
			s += string(val[i])
		} else if mask[i] == []byte("1")[0] {
			if val[i] == []byte("0")[0] {
				s += "1"
			} else if val[i] == []byte("1")[0] {
				s += "1"
			}
		} else if mask[i] == []byte("0")[0] {
			if val[i] == []byte("0")[0] {
				s += "0"
			} else if val[i] == []byte("1")[0] {
				s += "0"
			}
		}
	}
	return s
}

func applyMaskAddr(addr int, mask string) string {
	val := convertToBin(addr, masklength)

	s := ""
	for i := range mask {
		if mask[i] == []byte("X")[0] {
			s += "X"
		} else if mask[i] == []byte("1")[0] {
			s += "1"
		} else if mask[i] == []byte("0")[0] {
			s += string(val[i])
		}
	}
	return s
}

func swapFloatingBits(addr string, swap string) string {
	s := ""
	cnt := 0
	for i := range addr {
		if addr[i] == []byte("X")[0] {
			s += string(swap[cnt])
			cnt++
			continue
		}
		s += string(addr[i])
	}
	return s
}

func main() {
	fmt.Printf("Day %d\n", day)

	var answer1 int
	var answer2 int
	inputItems := aoc2020_shared.LoadString("input14.txt", "\n")

	bitMask := strings.Split(inputItems[0], " = ")[1]
	rules := inputItems[1:]

	memory := make(map[int]string)
	for _, ruleStr := range rules {
		if strings.HasPrefix(ruleStr, "mask") {
			bitMask = strings.Split(ruleStr, " = ")[1]
			continue
		}
		writeOp := parseToWriteOp(ruleStr)
		binVal := applyMask(writeOp.val, bitMask)
		memory[writeOp.addr] = binVal
	}

	// sum up values in memory
	for _, binary := range memory {
		valDec, _ := strconv.ParseInt(binary, 2, 64)
		answer1 += int(valDec)
	}

	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	// reset
	bitMask = strings.Split(inputItems[0], " = ")[1]

	memory2 := make(map[int]int)
	for _, ruleStr := range rules {
		if strings.HasPrefix(ruleStr, "mask") {
			bitMask = strings.Split(ruleStr, " = ")[1]
			continue
		}
		writeOp := parseToWriteOp(ruleStr)

		addrMask := applyMaskAddr(writeOp.addr, bitMask)

		floatingBitCnt := strings.Count(addrMask, "X")
		perms := int(math.Pow(2, float64(floatingBitCnt)))

		// this is some disgusting code
		i := 0
		for i < perms {
			swap := convertToBin(i, floatingBitCnt)
			valDec, _ := strconv.ParseInt(swapFloatingBits(addrMask, swap), 2, 64)
			memory2[int(valDec)] = writeOp.val
			i++
		}
	}

	// sum up values in memory
	for _, dec := range memory2 {
		answer2 += int(dec)
	}

	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
