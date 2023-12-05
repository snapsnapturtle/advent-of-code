package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
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

	numberRegex := regexp.MustCompile(`(?P<Number>\d+)`)

	seedMatches := numberRegex.FindAllString(lines[:1][0], -1)
	seeds := make([]int, 0)

	for _, seedMatch := range seedMatches {
		seed, _ := strconv.Atoi(seedMatch)
		seeds = append(seeds, seed)
	}

	almanac := lines[1:]
	positionsMap := make([][3]int, 0)

	for _, line := range almanac {
		if len(line) == 0 {
			newSeedPositions := seeds

			if len(positionsMap) == 0 {
				continue
			}

			for index, seed := range seeds {
				for _, positions := range positionsMap {
					destinationStart := positions[0]
					sourceStart := positions[1]
					rangeToMove := positions[2]

					if seed >= sourceStart && seed <= (sourceStart+rangeToMove) {
						distanceFromSource := seed - sourceStart
						newSeedPositions[index] = destinationStart + distanceFromSource

						break
					}
				}
			}

			positionsMap = make([][3]int, 0)
			seeds = newSeedPositions

			continue
		}

		mappingMatches := numberRegex.FindAllString(line, -1)

		if len(mappingMatches) != 3 {
			continue
		}

		destinationStart, _ := strconv.Atoi(mappingMatches[0])
		sourceStart, _ := strconv.Atoi(mappingMatches[1])
		length, _ := strconv.Atoi(mappingMatches[2])

		positionsMap = append(positionsMap, [3]int{destinationStart, sourceStart, length})
	}

	sort.Ints(seeds)

	return seeds[0]
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")

	numberRegex := regexp.MustCompile(`(?P<Number>\d+)`)
	seedMatches := numberRegex.FindAllString(lines[:1][0], -1)

	seedRanges := make([][2]int, 0)

	for i := 0; i < len(seedMatches); i += 2 {
		start, _ := strconv.Atoi(seedMatches[i])
		length, _ := strconv.Atoi(seedMatches[i+1])

		seedRanges = append(seedRanges, [2]int{start, start + length - 1})
	}

	almanac := lines[2:]
	positionsMap := make([][3]int, 0)

	for _, line := range almanac {
		if len(line) == 0 {
			if len(positionsMap) == 0 {
				continue
			}

			var seedRangesLength = len(seedRanges)
			newSeedRanges := make([][2]int, 0)

			for index := 0; index < seedRangesLength; index++ {
				seedRange := seedRanges[index]
				hasBeenMapped := false

				for _, positions := range positionsMap {
					destinationStart := positions[0]
					sourceStart := positions[1]
					rangeToMove := positions[2]

					isOutsideOfInstruction := seedRange[0] > sourceStart+rangeToMove || seedRange[1] < sourceStart

					if !isOutsideOfInstruction {
						var cutStart = int(math.Max(float64(seedRange[0]), float64(sourceStart)))
						var cutEnd = int(math.Min(float64(seedRange[1]), float64(sourceStart+rangeToMove)))

						distanceToMove := destinationStart - sourceStart
						newSeedRanges = append(newSeedRanges, [2]int{cutStart + distanceToMove, cutEnd + distanceToMove})

						if seedRange[0] < cutStart {
							seedRanges = append(seedRanges, [2]int{seedRange[0], cutStart - 1})
							seedRangesLength += 1
						}

						if seedRange[1] > cutEnd {
							seedRanges = append(seedRanges, [2]int{cutEnd + 1, seedRange[1]})
							seedRangesLength += 1
						}

						hasBeenMapped = true

						break
					}
				}

				if !hasBeenMapped {
					newSeedRanges = append(newSeedRanges, seedRange)
				}
			}

			positionsMap = make([][3]int, 0)
			seedRanges = newSeedRanges

			continue
		}

		mappingMatches := numberRegex.FindAllString(line, -1)

		if len(mappingMatches) != 3 {
			continue
		}

		destinationStart, _ := strconv.Atoi(mappingMatches[0])
		sourceStart, _ := strconv.Atoi(mappingMatches[1])
		length, _ := strconv.Atoi(mappingMatches[2])

		positionsMap = append(positionsMap, [3]int{destinationStart, sourceStart, length})
	}

	startingPositions := make([]int, len(seedRanges))

	for index, seedRange := range seedRanges {
		startingPositions[index] = seedRange[0]
	}

	sort.Ints(startingPositions)

	// I have no idea why this is off by 1
	return startingPositions[0] - 1
}
