package day_11

import (
	_ "embed"
	"math"
	"regexp"
	"snapsnapturtle/advent-of-code/util"
)

func PartOne(input string) int {
	universe := util.ParseLinesFromInput(input)

	return calculateShortestDistances(universe, 2)
}

func PartTwo(input string) int {
	universe := util.ParseLinesFromInput(input)

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
