package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
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

	if part == 1 {
		ans := partOne(input)
		fmt.Println("Output: ", ans)
	} else {
		ans := partTwo(input)
		fmt.Println("Output:", ans)
	}
}

func getMatchingNumbersAndGameId(line string) (int, int) {
	game := strings.Split(line, "|")

	gameDescription := game[0]
	gameNumbers := game[1]

	numberRegex := regexp.MustCompile(`(?P<Number>\d+)`)

	gameDescriptionMatches := numberRegex.FindAllString(gameDescription, -1)
	gameIdString, winningNumbers := gameDescriptionMatches[0], gameDescriptionMatches[1:]
	gameIdNumber, _ := strconv.Atoi(gameIdString)

	drawnNumbers := numberRegex.FindAllString(gameNumbers, -1)

	totalMatchingNumbers := 0
	for _, drawnNumber := range drawnNumbers {
		if slices.Contains(winningNumbers, drawnNumber) {
			totalMatchingNumbers += 1
		}
	}

	return totalMatchingNumbers, gameIdNumber
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")
	totalWinnings := 0

	for _, line := range lines {
		scoreForCard := 0
		matchingNumbers, _ := getMatchingNumbersAndGameId(line)

		for i := 1; i <= matchingNumbers; i++ {
			if scoreForCard == 0 {
				scoreForCard = 1
			} else {
				scoreForCard = scoreForCard * 2
			}
		}

		totalWinnings += scoreForCard
	}

	return totalWinnings
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	cardResults := make(map[int]int)
	scorecardsCount := make(map[int]int)

	for _, line := range lines {
		matchingNumbers, gameId := getMatchingNumbersAndGameId(line)

		cardResults[gameId] = matchingNumbers
		scorecardsCount[gameId] = 1
	}

	for currentIndex := 1; currentIndex <= len(scorecardsCount); currentIndex++ {
		scoreForCard := cardResults[currentIndex]
		scoreCardsOfType := scorecardsCount[currentIndex]

		for i := 1; i <= scoreForCard; i++ {
			scorecardsCount[currentIndex+i] += 1 * scoreCardsOfType
		}
	}

	totalScratchCards := 0

	for _, count := range scorecardsCount {
		totalScratchCards += count
	}

	return totalScratchCards
}
