package util

import "strings"

func ParseLinesFromInput(input string) []string {
	if input == "" {
		return []string{}
	}

	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}
