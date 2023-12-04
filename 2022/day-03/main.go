package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running Part:", part)

	if part == 1 {
		ans := partOne(input)
		fmt.Println("Output:", ans)
	} else {
		ans := partTwo(input)
		fmt.Println("Output:", ans)
	}
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")
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

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
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
