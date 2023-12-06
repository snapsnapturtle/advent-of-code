package util

import (
	"regexp"
	"strconv"
)

func ReadNumbersInString(input string) []int {
	var numbers []int

	numberRegex := regexp.MustCompile(`(?P<Number>\d+)`)
	numberMatches := numberRegex.FindAllString(input, -1)

	for _, numberMatch := range numberMatches {
		number, _ := strconv.Atoi(numberMatch)
		numbers = append(numbers, number)
	}

	return numbers
}
