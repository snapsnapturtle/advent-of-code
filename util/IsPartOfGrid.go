package util

func IsPartOfGrid(grid [][]string, position [2]int) bool {
	if position[0] < 0 || position[0] >= len(grid) {
		return false
	}

	if position[1] < 0 || position[1] >= len(grid[position[0]]) {
		return false
	}

	return true
}
