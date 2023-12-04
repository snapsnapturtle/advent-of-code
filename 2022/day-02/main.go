package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	totalScore := 0

	results := map[string]int{
		"AX": 3 + 1,
		"AY": 6 + 2,
		"AZ": 0 + 3,
		"BX": 0 + 1,
		"BY": 3 + 2,
		"BZ": 6 + 3,
		"CX": 6 + 1,
		"CY": 0 + 2,
		"CZ": 3 + 3,
	}

	for _, line := range lines {
		game := strings.Split(line, " ")
		score := results[game[0]+game[1]]

		totalScore += score
	}

	return totalScore
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	totalScore := 0

	results := map[string]int{
		"AX": 0 + 3,
		"AY": 3 + 1,
		"AZ": 6 + 2,
		"BX": 0 + 1,
		"BY": 3 + 2,
		"BZ": 6 + 3,
		"CX": 0 + 2,
		"CY": 3 + 3,
		"CZ": 6 + 1,
	}

	for _, line := range lines {
		game := strings.Split(line, " ")
		score := results[game[0]+game[1]]

		totalScore += score
	}

	return totalScore
}
