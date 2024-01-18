package day_03

import (
	_ "embed"
	"snapsnapturtle/advent-of-code/util"
	"strings"
	"unicode"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalPriorities := 0

	firstLower := int('a') - 1
	firstUpper := int('A') - 1

	for _, line := range lines {
		halfLength := len(line) / 2

		firstHalf := line[:halfLength]
		secondHalf := line[halfLength:]

		intersections := make(map[rune]bool)

		for i := 0; i < len(firstHalf); i++ {
			if strings.Contains(secondHalf, string(firstHalf[i])) {
				intersections[rune(firstHalf[i])] = true
			}
		}

		for letter := range intersections {
			if unicode.IsUpper(letter) {
				totalPriorities += int(letter) - firstUpper + 26
			} else {
				totalPriorities += int(letter) - firstLower
			}
		}
	}

	return totalPriorities
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)
	totalPriorities := 0

	firstLower := int('a') - 1
	firstUpper := int('A') - 1

	for index := 0; index < len(lines); index += 3 {
		occurrences := make(map[rune]int)

		for _, letter := range lines[index] {
			occurrences[letter] = 1
		}

		for _, letter := range lines[index+1] {
			if occurrences[letter] == 1 {
				occurrences[letter] = 2
			}
		}

		for _, letter := range lines[index+2] {
			if occurrences[letter] == 2 {
				occurrences[letter] = 3
			}
		}

		for letter, count := range occurrences {
			if count == 3 {
				if unicode.IsUpper(letter) {
					totalPriorities += int(letter) - firstUpper + 26
				} else {
					totalPriorities += int(letter) - firstLower
				}
			}
		}
	}

	return totalPriorities
}
