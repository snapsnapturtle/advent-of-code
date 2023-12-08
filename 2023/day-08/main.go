package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strings"
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

type Instruction struct {
	Left  string
	Right string
}

func getNextWaypoint(instruction string, currentWaypoint string, waypoints map[string]Instruction) (string, error) {
	var nextWaypoint string

	if instruction == "L" {
		nextWaypoint = waypoints[currentWaypoint].Left
	} else {
		nextWaypoint = waypoints[currentWaypoint].Right
	}

	if currentWaypoint == nextWaypoint {
		return "", errors.New("Recursion for " + currentWaypoint + " and " + nextWaypoint)
	}

	return nextWaypoint, nil
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// Not a math person here, so I took it this from the internet
// Find The Least Common Multiple (leastCommonMultiple) via greatestCommonDivisor
func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	walkingInstructions := strings.Split(lines[:1][0], "")

	waypoints := make(map[string]Instruction, len(lines)-2)

	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}

		navigationRegex := regexp.MustCompile(`(\w{3})`)
		matches := navigationRegex.FindAllString(line, -1)

		waypoints[matches[0]] = Instruction{Left: matches[1], Right: matches[2]}
	}

	steps := 0
	nextWaypoint := "AAA"

	for nextWaypoint != "ZZZ" {
		instructionIndex := steps % len(walkingInstructions)
		calculatedNextWaypoint, err := getNextWaypoint(walkingInstructions[instructionIndex], nextWaypoint, waypoints)

		if err != nil {
			fmt.Println(err)
			break
		}

		nextWaypoint = calculatedNextWaypoint
		steps++
	}

	return steps
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	walkingInstructions := strings.Split(lines[:1][0], "")

	waypoints := make(map[string]Instruction, len(lines)-2)

	for _, line := range lines[2:] {
		navigationRegex := regexp.MustCompile(`(\w{3})`)
		matches := navigationRegex.FindAllString(line, -1)

		waypoints[matches[0]] = Instruction{Left: matches[1], Right: matches[2]}
	}

	var stepsToEnd []int

	for waypoint := range waypoints {
		if waypoint[2] == 'A' {
			next := waypoint
			steps := 0

			for next[2] != 'Z' {
				instructionIndex := steps % len(walkingInstructions)

				newNext, _ := getNextWaypoint(walkingInstructions[instructionIndex], next, waypoints)

				next = newNext
				steps++
			}

			stepsToEnd = append(stepsToEnd, steps)
		}
	}

	totalSteps := leastCommonMultiple(stepsToEnd[0], stepsToEnd[1], stepsToEnd[2:]...)

	return totalSteps
}
