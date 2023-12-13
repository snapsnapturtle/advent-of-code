package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
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

	fmt.Println("--- Day 11: Cosmic Expansion ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func partOne(input string) int {
	universe := strings.Split(input, "\n")

	return calculateShortestDistances(universe, 2)
}

func partTwo(input string) int {
	universe := strings.Split(input, "\n")

	return calculateShortestDistances(universe, 1_000_000)
}

func calculateShortestDistances(universe []string, emptySpaceValue int) int {
	galaxies := findGalaxies(universe)
	rowIds, colIds := findExpandingSpaces(universe)

	totalDistance := 0
	pairs := 0

	for index, galaxy := range galaxies {
		for i := index + 1; i < len(galaxies); i++ {
			pairs++
			distantGalaxy := galaxies[i]

			distanceX := math.Abs(float64(galaxy[0] - distantGalaxy[0]))
			distanceY := math.Abs(float64(galaxy[1] - distantGalaxy[1]))

			// check if connection crosses through rowIds
			for _, rowId := range rowIds {
				if (rowId > galaxy[0] && rowId < distantGalaxy[0]) || (rowId < galaxy[0] && rowId > distantGalaxy[0]) {
					totalDistance += emptySpaceValue - 1
				}
			}

			// check if connection crosses through colIds
			for _, colId := range colIds {
				if (colId > galaxy[1] && colId < distantGalaxy[1]) || (colId < galaxy[1] && colId > distantGalaxy[1]) {
					totalDistance += emptySpaceValue - 1
				}
			}

			totalDistance += int(distanceX + distanceY)
		}
	}

	return totalDistance
}

func findGalaxies(universe []string) [][2]int {
	var galaxies [][2]int

	for rowIndex, row := range universe {
		for colIndex, col := range row {
			if col == '#' {
				galaxies = append(galaxies, [2]int{rowIndex, colIndex})
			}
		}
	}

	return galaxies
}

func findExpandingSpaces(universe []string) (rowIndices []int, colIndices []int) {
	var rowIds []int
	var colIds []int

	galaxyRunes := make(map[int]bool)

	for rowIndex, row := range universe {
		galaxyRuneRegex := regexp.MustCompile(`(#)`)
		galaxyRuneMatches := galaxyRuneRegex.FindAllStringIndex(row, -1)

		if len(galaxyRuneMatches) > 0 {
			for _, match := range galaxyRuneMatches {
				galaxyRuneIndex := match[0]
				galaxyRunes[galaxyRuneIndex] = true
			}
		} else {
			rowIds = append(rowIds, rowIndex)
		}
	}

	for colIndex := range universe[0] {
		if _, ok := galaxyRunes[colIndex]; !ok {
			colIds = append(colIds, colIndex)
		}
	}

	return rowIds, colIds
}
