package day_09

import (
	_ "embed"
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
)

func isEverythingZero(slice []int) bool {
	for _, number := range slice {
		if number != 0 {
			return false
		}
	}

	return true
}

func buildHistoryFromInput(line string) [][]int {
	regex := regexp.MustCompile(`(-?\d+)`)
	matches := regex.FindAllString(line, -1)

	var initialNumbers []int

	for _, match := range matches {
		number, _ := strconv.Atoi(match)
		initialNumbers = append(initialNumbers, number)
	}

	differenceRows := [][]int{initialNumbers}

	for !isEverythingZero(differenceRows[len(differenceRows)-1]) {
		latestRow := differenceRows[len(differenceRows)-1]
		currentDifferences := make([]int, 0)

		for i := 1; i < len(latestRow); i++ {
			previous := latestRow[i-1]
			current := latestRow[i]

			currentDifferences = append(currentDifferences, current-previous)
		}

		differenceRows = append(differenceRows, currentDifferences)
	}

	return differenceRows
}

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)

	extrapolatedNumber := 0

	for _, line := range lines {
		differenceRows := buildHistoryFromInput(line)

		for _, row := range differenceRows {
			extrapolatedNumber += row[len(row)-1]
		}
	}

	return extrapolatedNumber
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)

	extrapolatedNumbers := make([]int, 0)

	for _, line := range lines {
		differenceRows := buildHistoryFromInput(line)

		for i := len(differenceRows) - 1; i > 0; i-- {
			if i == len(differenceRows)-1 {
				differenceRows[i] = append([]int{0}, differenceRows[i]...)
			}

			previousRow := differenceRows[i]
			rowToCalculate := differenceRows[i-1]
			nextValue := rowToCalculate[0] - previousRow[0]

			differenceRows[i-1] = append([]int{nextValue}, rowToCalculate...)
		}

		extrapolatedNumbers = append(extrapolatedNumbers, differenceRows[0][0])
	}

	return util.SumOfSlice(extrapolatedNumbers)
}
