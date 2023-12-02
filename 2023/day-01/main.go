package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running: Part ", part)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if part == 1 {
		answer := partOne(file)
		fmt.Println("Output:", answer)
	} else {
		answer := partTwo(file)
		fmt.Println("Output:", answer)
	}
}

func partOne(file *os.File) int {
	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`([1-9])`)
		matches := r.FindAllString(line, -1)

		sumOfLine, _ := strconv.Atoi(matches[0] + matches[len(matches)-1])
		totalSum += sumOfLine
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalSum
}

var replacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func extractNumbersAndSpelledNumbers(line string) int {
	r := regexp.MustCompile(`([1-9])`)
	var adjustedLine = line

	for key, value := range replacements {
		adjustedLine = strings.ReplaceAll(adjustedLine, key, value)
	}

	matches := r.FindAllString(adjustedLine, -1)

	firstMatch := matches[0]
	lastMatch := matches[len(matches)-1]

	combinedNumber, _ := strconv.Atoi(firstMatch + lastMatch)

	return combinedNumber
}

func partTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		combinedNumber := extractNumbersAndSpelledNumbers(line)
		totalSum += combinedNumber
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalSum
}
