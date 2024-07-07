package day_01

import (
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalSum := 0

	regex := regexp.MustCompile(`([1-9])`)

	for _, line := range lines {
		matches := regex.FindAllString(line, -1)
		sumOfLine, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		totalSum += sumOfLine
	}

	return totalSum
}

var replacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func extractNumbersAndSpelledNumbers(line string) int {
	r := regexp.MustCompile(`([1-9])`)
	var adjustedLine = line

	for key, value := range replacements {
		adjustedLine = strings.ReplaceAll(adjustedLine, key, value)
	}

	matches := r.FindAllString(adjustedLine, -1)

	firstMatch := matches[0]
	lastMatch := matches[len(matches)-1]

	combinedNumber, _ := strconv.Atoi(firstMatch + lastMatch)

	return combinedNumber
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalSum := 0

	for _, line := range lines {
		combinedNumber := extractNumbersAndSpelledNumbers(line)
		totalSum += combinedNumber
	}

	return totalSum
}
