package aoc2020_shared

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// LoadInt - load input as int
func LoadInt(fileName string, sep string) []int {
	lines := loadInput(fileName, sep)
	return convertToInt(lines)
}

// LoadString - load input as str
func LoadString(fileName string, sep string) []string {
	return loadInput(fileName, "\n")
}

// LoadWords - load input as str
func LoadWords(fileName string, sep string) [][]string {
	lines := loadInput(fileName, "\n")
	final := make([][]string, len(lines))
	for i, line := range lines {
		words := strings.Split(line, sep)
		final[i] = words
	}
	return final
}

func loadInput(fileName string, sep string) []string {
	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Could not fetch from", fileName)
	}
	bodyBytes, err := ioutil.ReadAll(fp)

	if err != nil {
		log.Fatal("Could not read from", fileName)
	}
	bodyString := string(bodyBytes)
	lines := strings.Split(bodyString, sep)
	return lines
}

func convertToInt(lines []string) []int {
	var inputArray = []int{}

	for _, entry := range lines {
		intValue, err := strconv.Atoi(strings.TrimSpace(entry))
		if err != nil {
			panic(err)
		}
		inputArray = append(inputArray, intValue)
	}
	return inputArray
}
