package main

import "testing"

func TestX(t *testing.T) {
	var input = []int{-1, 0, 3, 5, 9, 12}

	index := binarySearch(input, 0, len(input)-1, 9)

	if input[index] != 9 {
		t.Errorf("Expected 9, got %d", input[index])
	}
}

func TestXx(t *testing.T) {
	var input = []int{-1, 0, 3, 5, 9, 12, 15, 18, 21, 24, 27, 30}

	index := binarySearch(input, 0, len(input)-1, -1)

	if input[index] != -1 {
		t.Errorf("Expected -1, got %d", input[index])
	}
}
