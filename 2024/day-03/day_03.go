package day_01

import (
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
)

func PartOne(input string) int {
	parsedInput := strings.Join(util.ParseLinesFromInput(input), "")
	totalSum := 0

	regex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := regex.FindAllStringSubmatch(parsedInput, -1)

	for _, match := range matches {
		leftNumber, _ := strconv.Atoi(match[1])
		rightNumber, _ := strconv.Atoi(match[2])

		totalSum += leftNumber * rightNumber
	}

	return totalSum
}

func PartTwo(input string) int {
	parsedInput := strings.Join(util.ParseLinesFromInput(input), "")
	totalSum := 0
	isMultiplicationEnabled := true

	regex := regexp.MustCompile(`(do\(\)|don't\(\))|mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

	matches := regex.FindAllStringSubmatch(parsedInput, -1)

	for _, match := range matches {
		if match[1] == "don't()" {
			isMultiplicationEnabled = false
			continue
		}

		if match[1] == "do()" {
			isMultiplicationEnabled = true
			continue
		}

		if isMultiplicationEnabled {
			leftNumber, _ := strconv.Atoi(match[2])
			rightNumber, _ := strconv.Atoi(match[3])

			totalSum += leftNumber * rightNumber
		}
	}

	return totalSum
}
