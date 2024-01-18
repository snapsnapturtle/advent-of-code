package util

import "strings"

func ParseLinesFromInput(input string) []string {
	if input == "" {
		return []string{}
	}

	return strings.Split(strings.Trim(input, "\n"), "\n")
}
