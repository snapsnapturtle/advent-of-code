package main

import (
	_ "embed"
	"errors"
	"fmt"
	"slices"
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

	fmt.Println("--- Day 10: Pipe Maze ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

const (
	Top    = iota
	Bottom = iota
	Left   = iota
	Right  = iota
)

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	loop := findLoop(lines)

	return len(loop) / 2
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	loop := findLoop(lines)

	return getEnclosedPoints(loop, len(lines[0]), len(lines))
}

func findFirstStartingPosition(lines []string) ([2]int, error) {
	for rowIndex, line := range lines {
		for colIndex, tile := range line {
			if tile == 'S' {
				return [2]int{rowIndex, colIndex}, nil
			}
		}
	}

	return [2]int{}, errors.New("no starting position")
}

func getNextPositionAndDirection(lines []string, currentPosition [2]int, currentDirection int) (newPosition [2]int, newDirection int, error error) {
	currentSymbol := lines[currentPosition[0]][currentPosition[1]]

	if currentSymbol == '-' {
		if currentDirection == Left {
			return [2]int{currentPosition[0], currentPosition[1] - 1}, Left, nil
		}

		if currentDirection == Right {
			return [2]int{currentPosition[0], currentPosition[1] + 1}, Right, nil
		}
	}

	if currentSymbol == '|' {
		if currentDirection == Top {
			return [2]int{currentPosition[0] - 1, currentPosition[1]}, Top, nil
		}

		if currentDirection == Bottom {
			return [2]int{currentPosition[0] + 1, currentPosition[1]}, Bottom, nil
		}
	}

	if currentSymbol == 'L' {
		if currentDirection == Bottom {
			return [2]int{currentPosition[0], currentPosition[1] + 1}, Right, nil
		}

		if currentDirection == Left {
			return [2]int{currentPosition[0] - 1, currentPosition[1]}, Top, nil
		}
	}

	if currentSymbol == 'J' {
		if currentDirection == Bottom {
			return [2]int{currentPosition[0], currentPosition[1] - 1}, Left, nil
		}

		if currentDirection == Right {
			return [2]int{currentPosition[0] - 1, currentPosition[1]}, Top, nil
		}
	}

	if currentSymbol == '7' {
		if currentDirection == Right {
			return [2]int{currentPosition[0] + 1, currentPosition[1]}, Bottom, nil
		}

		if currentDirection == Top {
			return [2]int{currentPosition[0], currentPosition[1] - 1}, Left, nil
		}
	}

	if currentSymbol == 'F' {
		if currentDirection == Top {
			return [2]int{currentPosition[0], currentPosition[1] + 1}, Right, nil
		}

		if currentDirection == Left {
			return [2]int{currentPosition[0] + 1, currentPosition[1]}, Bottom, nil
		}
	}

	return [2]int{}, Top, errors.New("not possible to go there")
}

func findLoop(lines []string) [][2]int {
	startingPosition, _ := findFirstStartingPosition(lines)

	for direction := 0; direction < 4; direction++ {
		var nextPosition [2]int
		nextDirection := direction
		loop := [][2]int{startingPosition}

		if direction == Top {
			nextPosition = [2]int{startingPosition[0] - 1, startingPosition[1]}
		}

		if direction == Bottom {
			nextPosition = [2]int{startingPosition[0] + 1, startingPosition[1]}
		}

		if direction == Left {
			nextPosition = [2]int{startingPosition[0], startingPosition[1] - 1}
		}

		if direction == Right {
			nextPosition = [2]int{startingPosition[0], startingPosition[1] + 1}
		}

		if !util.IsPartOfGrid(lines, nextPosition[0], nextPosition[1]) {
			continue
		}

		for {
			calculatedNextPosition, calculatedNextDirection, err := getNextPositionAndDirection(lines, nextPosition, nextDirection)

			if err != nil {
				fmt.Println(err)
				break
			}

			loop = append(loop, nextPosition)

			nextPosition = calculatedNextPosition
			nextDirection = calculatedNextDirection

			if lines[nextPosition[0]][nextPosition[1]] == 'S' {
				return loop
			}
		}
	}

	return [][2]int{}
}

func getEnclosedPoints(loop [][2]int, columns int, rows int) int {
	var enclosedPoints [][2]int

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			pointToCheck := [2]int{row, col}

			if slices.ContainsFunc(loop, func(position [2]int) bool {
				return position[0] == row && position[1] == col
			}) {
				continue
			}

			if util.IsPointWithinPolyline(loop, [2]int{row, col}) {
				enclosedPoints = append(enclosedPoints, pointToCheck)
			}
		}
	}

	drawLoop(loop, columns, rows)
	drawLoop(enclosedPoints, columns, rows)

	return len(enclosedPoints)
}

func drawLoop(markers [][2]int, columns int, rows int) {
	for row := 0; row < rows; row++ {
		line := ""

		for col := 0; col < columns; col++ {
			shouldMark := slices.ContainsFunc(markers, func(position [2]int) bool {
				return position[0] == row && position[1] == col
			})

			if shouldMark {
				line += "#"
			} else {
				line += "."
			}
		}
	}
}
