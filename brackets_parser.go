package main

import (
	"fmt"
	"math"
)

const (
	close = ']'
	open  = '['
)

func main() {
	swaps := minSwaps("][][")

	fmt.Printf("swaps: %d\n", swaps)
}

func minSwaps(s string) int {
	myString := s

	balance := findClosedToSwap(myString)
	if balance != 0 {
		balance = int(math.Abs(float64(balance)))
	}

	if balance <= 1 {
		return balance
	}

	return (balance / 2)
}

func findClosedToSwap(myString string) int {
	unbalanced := 0
	stop := len(myString)

	for i := 0; i < stop; i++ {
		if myString[i] == open {
			unbalanced++
		} else if myString[i] == close && unbalanced > 0 {
			unbalanced--
		} else if myString[i] == close {
			unbalanced++
		}
	}

	return unbalanced
}
