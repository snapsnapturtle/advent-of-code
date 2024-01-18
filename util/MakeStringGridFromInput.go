package util

func MakeStringGridFromInput(input string) [][]string {
	var grid [][]string

	lines := ParseLinesFromInput(input)

	for _, line := range lines {
		var row []string
		for _, r := range line {
			row = append(row, string(r))
		}

		grid = append(grid, row)
	}

	return grid
}
