package day_13

import (
	_ "embed"
	"strings"
)

func PartOne(input string) int {
	maps := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	totalNumber := 0

	for _, singleMap := range maps {
		rows := strings.Split(singleMap, "\n")

		if reflectionRows := findReflectionPoint(rows, 0); reflectionRows >= 0 {
			totalNumber += reflectionRows * 100
			continue
		}

		if reflectionCols := findReflectionPoint(transpose(rows), 0); reflectionCols >= 0 {
			totalNumber += reflectionCols
			continue
		}
	}

	return totalNumber
}

func PartTwo(input string) int {
	maps := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	totalNumber := 0

	for _, singleMap := range maps {
		rows := strings.Split(singleMap, "\n")

		if reflectionRows := findReflectionPoint(rows, 1); reflectionRows >= 0 {
			totalNumber += reflectionRows * 100
			continue
		}

		if reflectionCols := findReflectionPoint(transpose(rows), 1); reflectionCols >= 0 {
			totalNumber += reflectionCols
			continue
		}
	}

	return totalNumber
}

func transpose(rows []string) []string {
	output := make([]string, len(rows[0]))

	for i := 0; i < len(rows[0]); i++ {
		for j := 0; j < len(rows); j++ {
			output[i] += string(rows[j][i])
		}
	}

	return output
}

func findReflectionPoint(grid []string, allowedDifferences int) int {
	width := len(grid[0])
	height := len(grid)

	for mid := 0; mid < height-1; mid++ {
		differences := 0
		for col := 0; col < width; col++ {
			for offset := 0; ; offset++ {
				leftRow := mid - offset
				rightRow := mid + offset + 1

				if leftRow < 0 || rightRow >= height {
					break
				}

				if grid[leftRow][col] != grid[rightRow][col] {
					differences++
				}
			}
		}

		if differences == allowedDifferences {
			return mid + 1
		}
	}

	return -1
}
