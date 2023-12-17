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

	fmt.Println("--- Day 14: Parabolic Reflector Dish ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func partOne(input string) int {
	initialBoard := util.MakeStringGridFromInput(input)
	tiltedBoard := moveRocks(initialBoard, -1, 0)

	return calculateWeightForNorthBeam(tiltedBoard)
}

type CycleAndBoard struct {
	Cycle int
	Board [][]string
}

func partTwo(input string) int {
	board := util.MakeStringGridFromInput(input)
	memory := make(map[string]CycleAndBoard)

	cycles := 0
	cyclesToComplete := 1_000_000_000
	repetitionCycles := -1

	for cycles < cyclesToComplete {
		cycles++

		board = moveRocks(board, -1, 0)
		board = moveRocks(board, 0, -1)
		board = moveRocks(board, 1, 0)
		board = moveRocks(board, 0, 1)

		serialized := serialize(board)

		if value, ok := memory[serialized]; ok && repetitionCycles < 0 {
			repetitionCycles = cycles - value.Cycle
			cycles = cyclesToComplete - (cyclesToComplete-cycles)%repetitionCycles
		}

		memory[serialized] = CycleAndBoard{
			Cycle: cycles,
			Board: board,
		}
	}

	return calculateWeightForNorthBeam(board)
}

func serialize(board [][]string) string {
	output := ""

	for _, row := range board {
		output += strings.Join(row, "")
	}

	return output
}

func moveRocks(board [][]string, verticalDirection int, horizontalDirection int) [][]string {
	newBoard := board
	currentRowIndex := 0

	if verticalDirection > 0 {
		currentRowIndex = len(board) - 1
	}

	for currentRowIndex >= 0 && currentRowIndex < len(board) {
		currentColIndex := 0

		if horizontalDirection > 0 {
			currentColIndex = len(board[currentRowIndex]) - 1
		}

		for currentColIndex >= 0 && currentColIndex < (len(board[currentRowIndex])) {
			if newBoard[currentRowIndex][currentColIndex] == "O" {
				if verticalDirection != 0 {
					newRow := currentRowIndex + verticalDirection

					for newRow >= 0 && newRow < len(board) {
						if newBoard[newRow][currentColIndex] != "." {
							newBoard[newRow-verticalDirection][currentColIndex] = "O"

							if newRow-verticalDirection != currentRowIndex {
								newBoard[currentRowIndex][currentColIndex] = "."
							}

							break
						}

						if newRow == 0 {
							newBoard[newRow][currentColIndex] = "O"
							newBoard[currentRowIndex][currentColIndex] = "."
						}

						if newRow == len(board)-1 {
							newBoard[newRow][currentColIndex] = "O"
							newBoard[currentRowIndex][currentColIndex] = "."
						}

						newRow += verticalDirection
					}
				}

				if horizontalDirection != 0 {
					newCol := currentColIndex + horizontalDirection

					for newCol >= 0 && newCol < len(board[currentRowIndex]) {
						if newBoard[currentRowIndex][newCol] != "." {
							newBoard[currentRowIndex][newCol-horizontalDirection] = "O"

							if newCol-horizontalDirection != currentColIndex {
								newBoard[currentRowIndex][currentColIndex] = "."
							}

							break
						}

						if newCol == 0 {
							newBoard[currentRowIndex][newCol] = "O"
							newBoard[currentRowIndex][currentColIndex] = "."
						}

						if newCol == len(board[currentRowIndex])-1 {
							newBoard[currentRowIndex][newCol] = "O"
							newBoard[currentRowIndex][currentColIndex] = "."
						}

						newCol += horizontalDirection
					}
				}
			}

			if horizontalDirection == 0 {
				currentColIndex += 1
			}

			currentColIndex -= horizontalDirection
		}

		if verticalDirection == 0 {
			currentRowIndex += 1
		}

		currentRowIndex -= verticalDirection
	}

	return newBoard
}

func calculateWeightForNorthBeam(board [][]string) int {
	totalWeight := 0

	for index, row := range board {
		for _, field := range row {
			if field == "O" {
				totalWeight += len(board) - index
			}
		}
	}

	return totalWeight
}
