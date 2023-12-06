package main

import (
	_ "embed"
	"flag"
	"fmt"
	"snapsnapturtle/advent-of-code/util"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running Part:", part)

	if part == 1 {
		ans := partOne(input)
		fmt.Println("Output:", ans)
	} else {
		ans := partTwo(input)
		fmt.Println("Output:", ans)
	}
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")

	timeStrings := util.ReadNumbersInString(lines[0])
	distanceStrings := util.ReadNumbersInString(lines[1])
	attempts := make([]int, 0)

	for raceIndex := 0; raceIndex < len(timeStrings); raceIndex++ {
		time := timeStrings[raceIndex]
		distanceToBeat := distanceStrings[raceIndex]

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
