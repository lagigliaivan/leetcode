package main


func binarySearch(input []int, start int, end int, number int) int {
	if start >= len(input) {
		return -1
	}

	distance := ((end - start) / 2)
	middle := 0

	if distance != 0 {
		middle = end - ((end - start) / 2)
	}

	if input[middle] == number {
		return middle
	}

	if input[middle] < number {
		return binarySearch(input, middle, end, number)
	}

	if input[middle] >= number {
		return binarySearch(input, start, middle, number)
	}

	return -1
}
