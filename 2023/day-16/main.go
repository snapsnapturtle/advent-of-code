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

	fmt.Println("--- Day 16: The Floor Will Be Lava ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func partOne(input string) int {
	grid := util.MakeGridFromInput(input)
	firstStep := NextStep{
		Position:  [2]int{0, -1},
		Direction: [2]int{0, 1},
	}

	activeFields := calculateActiveFields(grid, firstStep)
	return activeFields
}

func partTwo(input string) int {
	grid := util.MakeGridFromInput(input)

	var firstSteps []NextStep

	for row := 0; row < len(grid); row++ {
		firstSteps = append(firstSteps, NextStep{
			Position:  [2]int{row, -1},
			Direction: [2]int{0, 1},
		}, NextStep{
			Position:  [2]int{row, len(grid[row])},
			Direction: [2]int{0, -1},
		})
	}

	for col := 0; col < len(grid[0]); col++ {
		firstSteps = append(firstSteps, NextStep{
			Position:  [2]int{-1, col},
			Direction: [2]int{1, 0},
		}, NextStep{
			Position:  [2]int{len(grid), col},
			Direction: [2]int{-1, 0},
		})
	}

	best := 0

	for _, firstStep := range firstSteps {
		r := calculateActiveFields(grid, firstStep)

		best = max(best, r)
	}

	return best
}

type NextStep struct {
	Position  [2]int
	Direction [2]int
}

func calculateActiveFields(grid [][]string, initialStep NextStep) int {
	visited := map[NextStep]bool{}
	toVisit := []NextStep{initialStep}

	for {
		if len(toVisit) == 0 {
			break
		}

		currentStep := toVisit[:1][0]

		if _, ok := visited[currentStep]; ok {
			toVisit = toVisit[1:]
			continue
		}

		possible, nextSteps := getNextSteps(grid, currentStep.Position, currentStep.Direction)

		if possible {
			toVisit = append(toVisit, nextSteps...)
		}

		toVisit = toVisit[1:]
		visited[currentStep] = true

	}

	var visitedSteps []NextStep
	for step := range visited {
		visitedSteps = append(visitedSteps, step)
	}

	return countActiveFields(visitedSteps) - 1
}

func countActiveFields(positions []NextStep) int {
	uniquePositions := make(map[[2]int]bool)

	for _, step := range positions {
		uniquePositions[step.Position] = true
	}

	return len(uniquePositions)
}

func getNextSteps(grid [][]string, start [2]int, direction [2]int) (possible bool, next []NextStep) {
	nextPosition := [2]int{start[0] + direction[0], start[1] + direction[1]}

	if !util.IsPartOfGrid(grid, nextPosition) {
		return false, []NextStep{}
	}

	nextInstruction := grid[nextPosition[0]][nextPosition[1]]

	// symbol + direction -> new direction
	reflectionMap := map[string]map[[2]int][2]int{
		"/": {
			[2]int{0, 1}:  [2]int{-1, 0},
			[2]int{0, -1}: [2]int{1, 0},
			[2]int{1, 0}:  [2]int{0, -1},
			[2]int{-1, 0}: [2]int{0, 1},
		},
		"\\": {
			[2]int{0, 1}:  [2]int{1, 0},
			[2]int{0, -1}: [2]int{-1, 0},
			[2]int{1, 0}:  [2]int{0, 1},
			[2]int{-1, 0}: [2]int{0, -1},
		},
	}

	if newDirection, ok := reflectionMap[nextInstruction][direction]; ok {
		return true, []NextStep{
			{
				Position:  nextPosition,
				Direction: newDirection,
			},
		}
	}

	if nextInstruction == "|" && direction[1] != 0 {
		return true, []NextStep{
			{
				Position:  nextPosition,
				Direction: [2]int{1, 0},
			},
			{
				Position:  nextPosition,
				Direction: [2]int{-1, 0},
			},
		}
	}

	if nextInstruction == "-" && direction[0] != 0 {
		return true, []NextStep{
			{
				Position:  nextPosition,
				Direction: [2]int{0, 1},
			},
			{
				Position:  nextPosition,
				Direction: [2]int{0, -1},
			},
		}
	}

	return true, []NextStep{
		{
			Position:  nextPosition,
			Direction: direction,
		},
	}
}
