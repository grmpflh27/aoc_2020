package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/grmpflh27/aoc_2020/aoc2020_shared"
)

const day int = 12

type Orientation string

const (
	West  Orientation = "West"
	East  Orientation = "East"
	North Orientation = "North"
	South Orientation = "South"
)

type Step struct {
	direction string
	amnt      int
}

func parseStep(step string) Step {
	dir := string(step[0])
	amnt, _ := strconv.Atoi(step[1:])
	return Step{dir, amnt}
}

func main() {
	fmt.Printf("Day %d\n", day)

	easting := 0
	northing := 0
	orientation := "E"
	curDeg := 90

	dirByDeg := map[int]string{
		90:  "E",
		180: "S",
		270: "W",
		0:   "N",
	}

	var answer1 int
	var answer2 int
	steps := aoc2020_shared.LoadString("input12.txt", "\n")

	for _, stepStr := range steps {
		step := parseStep(stepStr)
		switch step.direction {
		case "N":
			northing += step.amnt
		case "S":
			northing -= step.amnt
		case "E":
			easting += step.amnt
		case "W":
			easting -= step.amnt
		case "F":
			switch orientation {
			case "N":
				northing += step.amnt
			case "S":
				northing -= step.amnt
			case "E":
				easting += step.amnt
			case "W":
				easting -= step.amnt
			}
		case "R":
			curDeg += step.amnt
			if curDeg >= 360 {
				curDeg -= 360
			}
			orientation = dirByDeg[curDeg]
		case "L":
			curDeg -= step.amnt
			if curDeg < 0 {
				curDeg += 360
			}
			orientation = dirByDeg[curDeg]
		}

		fmt.Printf("[%d, %d]\n", easting, northing)
	}

	answer1 = int(math.Abs(float64(easting)) + math.Abs(float64(northing)))
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Println("=======================")

	wpEasting := 10
	wpNorthing := 1

	easting = 0
	northing = 0

	for _, stepStr := range steps {
		step := parseStep(stepStr)
		switch step.direction {
		case "N":
			wpNorthing += step.amnt
		case "S":
			wpNorthing -= step.amnt
		case "E":
			wpEasting += step.amnt
		case "W":
			wpEasting -= step.amnt
		case "F":
			deltaNorthing := step.amnt * wpNorthing
			deltaEasting := step.amnt * wpEasting
			northing += deltaNorthing
			easting += deltaEasting
		case "R":
			tmpEast := wpEasting
			tmpNorth := wpNorthing
			switch step.amnt {
			case 90:
				wpNorthing = -1 * tmpEast
				wpEasting = tmpNorth
			case 180:
				wpNorthing *= -1
				wpEasting *= -1
			case 270:
				wpNorthing = tmpEast
				wpEasting = -1 * tmpNorth
			}
		case "L":
			tmpEast := wpEasting
			tmpNorth := wpNorthing
			switch step.amnt {
			case 90:
				wpNorthing = tmpEast
				wpEasting = -1 * tmpNorth
			case 180:
				wpNorthing *= -1
				wpEasting *= -1
			case 270:
				wpNorthing = -1 * tmpEast
				wpEasting = tmpNorth
			}
		}

		fmt.Printf("[%d, %d] - wp [%d, %d]\n", easting, northing, wpEasting, wpNorthing)
	}

	answer2 = int(math.Abs(float64(easting)) + math.Abs(float64(northing)))
	fmt.Printf("Answer 2: %d\n", answer2)
	fmt.Println("=======================")

}
