package day_01

import (
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"sort"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalSum := 0

	regex := regexp.MustCompile(`([0-9]+)\s+([0-9]+)`)

	leftNumbers := make([]int, len(lines))
	rightNumbers := make([]int, len(lines))

	for index, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		leftNumber, _ := strconv.Atoi(matches[0][1])
		rightNumber, _ := strconv.Atoi(matches[0][2])

		leftNumbers[index] = leftNumber
		rightNumbers[index] = rightNumber
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	for index := 0; index < len(leftNumbers); index++ {
		totalSum += Abs(leftNumbers[index] - rightNumbers[index])
	}

	return totalSum
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalSum := 0

	regex := regexp.MustCompile(`([0-9]+)\s+([0-9]+)`)

	leftNumbers := make([]int, len(lines))
	rightNumbers := make([]int, len(lines))

	for index, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)

		leftNumber, _ := strconv.Atoi(matches[0][1])
		rightNumber, _ := strconv.Atoi(matches[0][2])

		leftNumbers[index] = leftNumber
		rightNumbers[index] = rightNumber
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	for _, leftNumber := range leftNumbers {
		matchesInRightNumbers := 0

		for _, rightNumber := range rightNumbers {
			if leftNumber == rightNumber {
				matchesInRightNumbers++
			}
		}

		totalSum += leftNumber * matchesInRightNumbers
	}

	return totalSum
}
