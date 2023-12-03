package util

func IsPartOfGrid(grid []string, row int, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}

	if col < 0 || col >= len(grid[row]) {
		return false
	}

	return true
}
