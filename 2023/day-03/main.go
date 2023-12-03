package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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

func getAdjacentFieldsForLine(lineIndex int, columnIndex int) [][2]int {
	var adjacentFields [][2]int

	// top left
	adjacentFields = append(adjacentFields, [2]int{lineIndex - 1, columnIndex - 1})

	// top center
	adjacentFields = append(adjacentFields, [2]int{lineIndex - 1, columnIndex})

	// top right
	adjacentFields = append(adjacentFields, [2]int{lineIndex - 1, columnIndex + 1})

	// center left
	adjacentFields = append(adjacentFields, [2]int{lineIndex, columnIndex - 1})

	// center right
	adjacentFields = append(adjacentFields, [2]int{lineIndex, columnIndex + 1})

	// bottom left
	adjacentFields = append(adjacentFields, [2]int{lineIndex + 1, columnIndex - 1})

	// bottom center
	adjacentFields = append(adjacentFields, [2]int{lineIndex + 1, columnIndex})

	// bottom right
	adjacentFields = append(adjacentFields, [2]int{lineIndex + 1, columnIndex + 1})

	return adjacentFields
}

func readFullNumber(line string, colIndex int) string {
	var fullNumber string
	var startIndex = colIndex

	for startIndex > 0 && unicode.IsDigit(rune(line[startIndex-1])) {
		startIndex--
	}

	for startIndex < len(line) && unicode.IsDigit(rune(line[startIndex])) {
		fullNumber += string(line[startIndex])
		startIndex++
	}

	return fullNumber
}

func isPartOfSchematic(lines []string, lineIndex int, colIndex int) bool {
	if lineIndex < 0 || lineIndex >= len(lines) {
		return false
	}

	if colIndex < 0 || colIndex >= len(lines[lineIndex]) {
		return false
	}

	return true
}

func main() {
	partTwo()
}

func partOne() {
	lines := strings.Split(input, "\n")
	totalPartNumbers := 0

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				adjacentFields := getAdjacentFieldsForLine(lineIndex, charIndex)
				processedNumbers := make(map[string]bool)

				for _, adjacentField := range adjacentFields {
					if !isPartOfSchematic(lines, adjacentField[0], adjacentField[1]) {
						continue
					}

					if unicode.IsDigit(rune(lines[adjacentField[0]][adjacentField[1]])) {
						numberString := readFullNumber(lines[adjacentField[0]], adjacentField[1])
						_, exists := processedNumbers[numberString]

						if exists {
							continue
						} else {
							processedNumbers[numberString] = true
							number, _ := strconv.Atoi(numberString)
							totalPartNumbers += number
							continue
						}
					}
				}
			}
		}
	}

	fmt.Println("Total part numbers:", totalPartNumbers)
}

func partTwo() {
	lines := strings.Split(input, "\n")
	totalGearRatios := 0

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char == '*' {
				adjacentFields := getAdjacentFieldsForLine(lineIndex, charIndex)
				processedNumbers := make(map[string]bool)

				for _, adjacentField := range adjacentFields {
					if !isPartOfSchematic(lines, adjacentField[0], adjacentField[1]) {
						continue
					}

					if unicode.IsDigit(rune(lines[adjacentField[0]][adjacentField[1]])) {
						numberString := readFullNumber(lines[adjacentField[0]], adjacentField[1])
						_, numberAlreadyRead := processedNumbers[numberString]

						if numberAlreadyRead {
							continue
						} else {
							processedNumbers[numberString] = true

							if len(processedNumbers) == 2 {
								var product = 1

								for number := range processedNumbers {
									numberInt, _ := strconv.Atoi(number)
									product *= numberInt
								}

								totalGearRatios += product
							}

							continue
						}
					}
				}
			}
		}
	}

	fmt.Println("Total part numbers:", totalGearRatios)
}
