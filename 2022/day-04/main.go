package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running Part: ", part)

	if part == 1 {
		ans := partOne(input)
		fmt.Println("Output: ", ans)
	} else {
		ans := partTwo(input)
		fmt.Println("Output:", ans)
	}
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")

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

func partTwo(input string) int {
	lines := strings.Split(input, "\n")

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
