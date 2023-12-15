package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
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

	fmt.Println("--- Day 15: Lens Library ---")
	fmt.Println("Part 1:", partOne(input))
	fmt.Println("Part 2:", partTwo(input))
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

func partOne(input string) int {
	words := strings.Split(input, ",")
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

func partTwo(input string) int {
	commands := strings.Split(input, ",")
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
