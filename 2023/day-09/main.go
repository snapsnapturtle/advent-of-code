package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
	"time"
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
	timeStart := time.Now()

	fmt.Println("--- Day 9: Mirage Maintenance ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

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

func partOne(input string) int {
	lines := strings.Split(input, "\n")

	extrapolatedNumber := 0

	for _, line := range lines {
		differenceRows := buildHistoryFromInput(line)

		for _, row := range differenceRows {
			extrapolatedNumber += row[len(row)-1]
		}
	}

	return extrapolatedNumber
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")

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
