package util

func SumOfSlice(slice []int) int {
	var sum int

	for _, number := range slice {
		sum += number
	}

	return sum
}
