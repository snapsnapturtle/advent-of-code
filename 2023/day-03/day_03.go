package day_03

import (
	_ "embed"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
	"unicode"
)

func readFullNumber(line string, colIndex int) string {
	var fullArray []string
	var startIndex = colIndex

	for startIndex > 0 && unicode.IsDigit(rune(line[startIndex-1])) {
		startIndex--
		fullArray = append([]string{string(line[startIndex])}, fullArray...)
	}

	startIndex = colIndex

	for startIndex < len(line) && unicode.IsDigit(rune(line[startIndex])) {
		fullArray = append(fullArray, string(line[startIndex]))
		startIndex++
	}

	return strings.Join(fullArray, "")
}

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalPartNumbers := 0

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				adjacentFields := util.GetAdjacentFieldsForLine(lineIndex, charIndex)
				processedNumbers := make(map[string]bool)

				for _, adjacentField := range adjacentFields {
					if !util.IsPartOfLines(lines, adjacentField[0], adjacentField[1]) {
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

	return totalPartNumbers
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalGearRatios := 0

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char == '*' {
				adjacentFields := util.GetAdjacentFieldsForLine(lineIndex, charIndex)
				processedNumbers := make(map[string]bool)

				for _, adjacentField := range adjacentFields {
					if !util.IsPartOfLines(lines, adjacentField[0], adjacentField[1]) {
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

	return totalGearRatios
}
