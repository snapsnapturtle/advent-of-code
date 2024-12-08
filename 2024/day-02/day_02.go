package day_01

import (
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
)

func PartOne(input string) int {
	reports := util.ParseLinesFromInput(input)
	safeReports := 0

	for _, report := range reports {
		levels := strings.Fields(report)
		isValidReport := true

		previousLevelValue, _ := strconv.Atoi(levels[0])
		previousLevelDifference := 0

		for i, level := range levels {
			if i == 0 {
				continue
			}

			levelValue, _ := strconv.Atoi(level)

			absoluteDifference := util.Abs(levelValue - previousLevelValue)

			if absoluteDifference < 1 || absoluteDifference > 3 {
				isValidReport = false
				break
			}

			difference := levelValue - previousLevelValue

			if (previousLevelDifference > 0 && difference < 0) || previousLevelDifference < 0 && difference > 0 {
				isValidReport = false
				break
			}

			previousLevelValue = levelValue
			previousLevelDifference = difference
		}

		if isValidReport {
			safeReports++
		}
	}

	return safeReports
}

func PartTwo(input string) int {
	// todo: implement
	return 0
}
