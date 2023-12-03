package main

import (
	_ "embed"
	"flag"
	"fmt"
	"snapsnapturtle/advent-of-code/util"
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

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running Part: ", part)

	if part != 1 {
		ans := partOne(input)
		fmt.Println("Output: ", ans)
	} else {
		ans := partTwo(input)
		fmt.Println("Output:", ans)
	}
}

func readFullNumber(line string, colIndex int) string {
	var fullNumber string
	var fullArray []string
	var startIndex = colIndex

	for startIndex > 0 && unicode.IsDigit(rune(line[startIndex-1])) {
		startIndex--
		fullArray = append([]string{string(line[startIndex])}, fullArray...)
	}

	startIndex = colIndex

	for startIndex < len(line) && unicode.IsDigit(rune(line[startIndex])) {
		fullNumber += string(line[startIndex])
		fullArray = append(fullArray, string(line[startIndex]))
		startIndex++
	}

	return strings.Join(fullArray, "")
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	totalPartNumbers := 0

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				adjacentFields := util.GetAdjacentFieldsForLine(lineIndex, charIndex)
				processedNumbers := make(map[string]bool)

				for _, adjacentField := range adjacentFields {
					if !util.IsPartOfGrid(lines, adjacentField[0], adjacentField[1]) {
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

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	totalGearRatios := 0

	for lineIndex, line := range lines {
		for charIndex, char := range line {
			if char == '*' {
				adjacentFields := util.GetAdjacentFieldsForLine(lineIndex, charIndex)
				processedNumbers := make(map[string]bool)

				for _, adjacentField := range adjacentFields {
					if !util.IsPartOfGrid(lines, adjacentField[0], adjacentField[1]) {
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
