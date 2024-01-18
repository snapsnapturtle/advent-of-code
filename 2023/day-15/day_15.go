package day_15

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func PartOne(input string) int {
	words := strings.Split(strings.TrimRight(input, "\n"), ",")
	total := 0

	for _, word := range words {
		total += hash(word)
	}

	return total
}

type Lens struct {
	Label       string
	FocalLength int
}

func PartTwo(input string) int {
	commands := strings.Split(strings.TrimRight(input, "\n"), ",")
	fmt.Println(len(commands))
	boxes := make(map[int][]Lens, 256)

	for _, word := range commands {
		regex := regexp.MustCompile(`(?P<Label>\w+)(?P<Operator>[=-])(?P<FocalLength>\d+)?`)
		matches := regex.FindStringSubmatch(word)

		lensLabel := matches[regex.SubexpIndex("Label")]
		operator := matches[regex.SubexpIndex("Operator")]
		focalLength, _ := strconv.Atoi(matches[regex.SubexpIndex("FocalLength")])

		boxNumber := hash(lensLabel)

		if operator == "-" {
			boxes[boxNumber] = slices.DeleteFunc(boxes[boxNumber], func(lens Lens) bool {
				return lens.Label == lensLabel
			})
		}

		if operator == "=" {
			_, ok := boxes[boxNumber]

			if !ok {
				boxes[boxNumber] = make([]Lens, 0)
			}

			existingIndex := slices.IndexFunc(boxes[boxNumber], func(lens Lens) bool {
				return lens.Label == lensLabel
			})

			lens := Lens{
				Label:       lensLabel,
				FocalLength: focalLength,
			}

			if existingIndex >= 0 {
				boxes[boxNumber][existingIndex] = lens
			} else {
				boxes[boxNumber] = append(boxes[boxNumber], lens)
			}
		}
	}

	return calculateFocusPower(boxes)
}

func hash(value string) int {
	current := 0

	for _, char := range value {
		current += int(char)
		current *= 17
		current = current % 256
	}

	return current
}

func calculateFocusPower(boxes map[int][]Lens) int {
	total := 0

	for boxIndex, lenses := range boxes {
		if len(lenses) == 0 {
			continue
		}

		for lensIndex, lens := range lenses {
			total += (boxIndex + 1) * (lensIndex + 1) * lens.FocalLength
		}
	}

	return total
}
