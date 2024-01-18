package day_01

import (
	_ "embed"
	"snapsnapturtle/advent-of-code/util"
	"sort"
	"strconv"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)

	maxCalories := 0
	currentElfCalories := 0

	for _, line := range lines {
		if len(line) == 0 {
			currentElfCalories = 0
			continue
		}

		calories, _ := strconv.Atoi(line)

		currentElfCalories += calories

		if currentElfCalories > maxCalories {
			maxCalories = currentElfCalories
		}
	}

	return maxCalories
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)

	var calorieCounts []int
	currentElfCalories := 0

	for _, line := range lines {
		if len(line) == 0 {
			calorieCounts = append(calorieCounts, currentElfCalories)
			currentElfCalories = 0

			continue
		}

		calories, _ := strconv.Atoi(line)

		currentElfCalories += calories
	}

	calorieCounts = append(calorieCounts, currentElfCalories)
	currentElfCalories = 0

	sort.Ints(calorieCounts)
	sort.Sort(sort.Reverse(sort.IntSlice(calorieCounts)))

	return calorieCounts[0] + calorieCounts[1] + calorieCounts[2]
}
