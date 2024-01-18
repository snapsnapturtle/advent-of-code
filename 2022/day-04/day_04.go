package day_04

import (
	_ "embed"
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)

	var overlappingRanges int

	for _, line := range lines {
		regex := regexp.MustCompile(`(\d+)`)
		matches := regex.FindAllString(line, -1)

		rangeStartA, _ := strconv.Atoi(matches[0])
		rangeEndA, _ := strconv.Atoi(matches[1])
		rangeStartB, _ := strconv.Atoi(matches[2])
		rangeEndB, _ := strconv.Atoi(matches[3])

		if (rangeStartA >= rangeStartB && rangeEndA <= rangeEndB) || (rangeStartB >= rangeStartA && rangeEndB <= rangeEndA) {
			overlappingRanges++
		}
	}

	return overlappingRanges
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)

	var overlappingRanges int

	for _, line := range lines {
		regex := regexp.MustCompile(`(\d+)`)
		matches := regex.FindAllString(line, -1)

		rangeStartA, _ := strconv.Atoi(matches[0])
		rangeEndA, _ := strconv.Atoi(matches[1])
		rangeStartB, _ := strconv.Atoi(matches[2])
		rangeEndB, _ := strconv.Atoi(matches[3])

		if (rangeStartA >= rangeStartB && rangeStartA <= rangeEndB) ||
			(rangeEndA >= rangeStartB && rangeEndA <= rangeEndB) ||
			(rangeStartB >= rangeStartA && rangeStartB <= rangeEndA) ||
			(rangeEndB >= rangeStartA && rangeEndB <= rangeEndA) {
			overlappingRanges++
		}
	}

	return overlappingRanges
}
