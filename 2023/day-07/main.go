package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
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
	timeStart := time.Now()

	fmt.Println("--- Day 7: Camel Cards ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

type ScoredCard struct {
	Cards []string
	Score int
	Bid   int
}

func evaluateCards(cards []string, joker string) int {
	cardToAmountMap := make(map[string]int, len(cards))
	jokers := 0

	for _, card := range cards {
		if card == joker {
			jokers += 1
			continue
		}

		cardToAmountMap[card] += 1
	}

	cardAmounts := make([]int, 0, len(cards))

	for _, amount := range cardToAmountMap {
		cardAmounts = append(cardAmounts, amount)
	}

	sort.Slice(cardAmounts, func(i, j int) bool {
		return cardAmounts[i] > cardAmounts[j]
	})

	if len(cardAmounts) == 0 {
		cardAmounts = append(cardAmounts, 5)
	} else {
		cardAmounts[0] += jokers
	}

	totalScore := 0

	for index, amount := range cardAmounts {
		totalScore += amount * (5 - index)
	}

	return totalScore
}

func sortScoredCards(scores []ScoredCard, cardValues map[string]int) []ScoredCard {
	sort.Slice(scores, func(i, j int) bool {
		if scores[i].Score == scores[j].Score {
			for cardIndex := 0; cardIndex < len(scores[i].Cards); cardIndex++ {
				if cardValues[scores[i].Cards[cardIndex]] == cardValues[scores[j].Cards[cardIndex]] {
					continue
				}

				return cardValues[scores[i].Cards[cardIndex]] < cardValues[scores[j].Cards[cardIndex]]
			}
		}

		return scores[i].Score < scores[j].Score
	})

	return scores
}

func partOne(input string) int {
	lines := strings.Split(input, "\n")

	cardValues := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	var scores []ScoredCard

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")

		cards := strings.Split(parts[0], "")
		bid, _ := strconv.Atoi(parts[1])

		scores = append(scores, ScoredCard{
			Cards: cards,
			Score: evaluateCards(cards, ""),
			Bid:   bid,
		})
	}

	sortScoredCards(scores, cardValues)

	totalScore := 0

	for index, scoredCard := range scores {
		totalScore += scoredCard.Bid * (index + 1)
	}

	return totalScore
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")

	cardValues := map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	var scores []ScoredCard

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, " ")

		cards := strings.Split(parts[0], "")
		bid, _ := strconv.Atoi(parts[1])
		handScore := evaluateCards(cards, "J")

		scores = append(scores, ScoredCard{
			Cards: cards,
			Score: handScore,
			Bid:   bid,
		})
	}

	sortScoredCards(scores, cardValues)

	totalScore := 0

	for index, scoredCard := range scores {
		totalScore += scoredCard.Bid * (index + 1)
	}

	return totalScore
}
