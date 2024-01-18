package day_02

import (
	"regexp"
	"snapsnapturtle/advent-of-code/util"
	"strconv"
	"strings"
)

func PartOne(input string) int {
	lines := util.ParseLinesFromInput(input)

	total := 0

	for _, line := range lines {
		re := regexp.MustCompile(`Game (\d+):\s`)
		matches := re.FindStringSubmatch(line)

		if matches == nil {
			continue
		}

		gameId, _ := strconv.Atoi(matches[1])
		game := strings.Split(strings.ReplaceAll(line, matches[0], ""), "; ")

		possibleDraw := true

		for _, drawString := range game {
			draws := make(map[string]int)
			countAndColorStrings := strings.Split(drawString, ", ")

			for _, countAndColorString := range countAndColorStrings {
				parts := strings.Split(countAndColorString, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]

				draws[color] = count
			}

			if draws["red"] > 12 || draws["green"] > 13 || draws["blue"] > 14 {
				possibleDraw = false
			}
		}

		if possibleDraw {
			total += gameId
		}
	}

	return total
}

func PartTwo(input string) int {
	lines := util.ParseLinesFromInput(input)

	total := 0

	for _, line := range lines {
		draws := strings.Split(line, "; ")
		minimums := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, drawString := range draws {
			countAndColorStrings := strings.Split(drawString, ", ")

			for _, countAndColorString := range countAndColorStrings {
				parts := strings.Split(countAndColorString, " ")
				count, _ := strconv.Atoi(parts[0])
				color := parts[1]

				minimums[color] = util.Max(minimums[color], count)
			}
		}

		total += minimums["red"] * minimums["green"] * minimums["blue"]
	}

	return total
}
