package day_01

func PartOne(input string) int {
	floor := 0

	for _, r := range input {
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		}
	}

	return floor
}

func PartTwo(input string) int {
	floor := 0

	for i, r := range input {
		if r == '(' {
			floor++
		} else if r == ')' {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}
