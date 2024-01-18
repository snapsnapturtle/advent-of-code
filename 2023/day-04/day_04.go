package day_04

import (
	_ "embed"
	"regexp"
	"slices"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
)

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

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)
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

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)
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
