package main

import (
	_ "embed"
	"fmt"
	"snapsnapturtle/advent-of-code/util"
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

	fmt.Println("--- Day 6: Wait For It ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")

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

func partTwo(input string) int {
	lines := strings.Split(input, "\n")

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
