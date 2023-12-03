package util

func GetAdjacentFieldsForLine(lineIndex int, columnIndex int) [][2]int {
	var adjacentFields [][2]int

	// top left
	adjacentFields = append(adjacentFields, [2]int{lineIndex - 1, columnIndex - 1})

	// top center
	adjacentFields = append(adjacentFields, [2]int{lineIndex - 1, columnIndex})

	// top right
	adjacentFields = append(adjacentFields, [2]int{lineIndex - 1, columnIndex + 1})

	// center left
	adjacentFields = append(adjacentFields, [2]int{lineIndex, columnIndex - 1})

	// center right
	adjacentFields = append(adjacentFields, [2]int{lineIndex, columnIndex + 1})

	// bottom left
	adjacentFields = append(adjacentFields, [2]int{lineIndex + 1, columnIndex - 1})

	// bottom center
	adjacentFields = append(adjacentFields, [2]int{lineIndex + 1, columnIndex})

	// bottom right
	adjacentFields = append(adjacentFields, [2]int{lineIndex + 1, columnIndex + 1})

	return adjacentFields
}
