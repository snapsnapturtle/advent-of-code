package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
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
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running Part: ", part)

	if part != 1 {
		ans := partOne(input)
		fmt.Println("Output: ", ans)
	} else {
		ans := partTwo(input)
		fmt.Println("Output:", ans)
	}
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")

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

func partTwo(input string) int {
	lines := strings.Split(input, "\n")

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
