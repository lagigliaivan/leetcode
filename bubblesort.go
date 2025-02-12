package main

func bubbleSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-(i+1); j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}

	return input
}
