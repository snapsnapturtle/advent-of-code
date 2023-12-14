package util

import "strings"

func MakeGridFromInput(input string) [][]string {
	var grid [][]string

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		var row []string
		for _, r := range line {
			row = append(row, string(r))
		}

		grid = append(grid, row)
	}

	return grid
}
