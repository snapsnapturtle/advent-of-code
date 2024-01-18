package day_02

import (
	_ "embed"
	"snapsnapturtle/advent-of-code/util"
	"strings"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)
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

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)
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
