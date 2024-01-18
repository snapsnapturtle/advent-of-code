package day_06

import (
	_ "embed"
	"snapsnapturtle/advent-of-code/util"
	"strings"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)

	timeStrings := util.ReadNumbersInString(lines[0])
	distanceStrings := util.ReadNumbersInString(lines[1])
	attempts := make([]int, 0)

	for raceIndex := 0; raceIndex < len(timeStrings); raceIndex++ {
		ranceTime := timeStrings[raceIndex]
		distanceToBeat := distanceStrings[raceIndex]

		validAttempts := 0

		for timePressed := 0; timePressed <= ranceTime; timePressed++ {
			distance := timePressed * (ranceTime - timePressed)

			if distance > distanceToBeat {
				validAttempts++
			}
		}

		attempts = append(attempts, validAttempts)
	}

	product := 1

	for _, attempt := range attempts {
		product *= attempt
	}

	return product
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)

	times := util.ReadNumbersInString(strings.ReplaceAll(lines[0], " ", ""))
	distancesToBeat := util.ReadNumbersInString(strings.ReplaceAll(lines[1], " ", ""))

	attempts := make([]int, 0)

	for raceIndex := 0; raceIndex < len(times); raceIndex++ {
		time := times[raceIndex]
		distanceToBeat := distancesToBeat[raceIndex]

		validAttempts := 0

		for timePressed := 0; timePressed <= time; timePressed++ {
			distance := timePressed * (time - timePressed)

			if distance > distanceToBeat {
				validAttempts++
			}
		}

		attempts = append(attempts, validAttempts)
	}

	product := 1

	for _, attempt := range attempts {
		product *= attempt
	}

	return product
}
