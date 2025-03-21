package main

import (
	"fmt"
	"reflect"
)

// Problem Statement:

// You are given:

// numbers: A sorted slice of integers (may contain duplicates, but is sorted).
// lower: An integer representing the lower bound of the range (inclusive).
// upper: An integer representing the upper bound of the range (inclusive).
// Your function should identify the "gaps" within the range [lower, upper] that are not present in the numbers slice.

// A "gap" is defined as a contiguous sequence of integers within the range [lower, upper] that are missing from the numbers sl
// ice.  The function should return a slice of slices, where each inner slice represents a gap and contains exactly two integers: the start and end of the gap (inclusive).

/*{1,3,5,9,10,12}

lower=1
upper=5


3 => {0,0}
5 => {4,4}
9 => {6,8}
10 => {0,0}
12 => {11,11}
*/

func findGaps(numbers []int, lower int, upper int) [][]int {
	result := [][]int{}
	missing := []int{}

	//exceptions
	if lower > upper {
		return result
	}

	if len(numbers) == 0 || lower > numbers[len(numbers)-1] {
		return append(result, []int{lower, upper})
	}

	//normal flow
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < lower {
			continue
		}

		//if diff is 0 then no gap
		x := numbers[i] - lower - 1
		if numbers[i] == lower || (x == 0 && i != 0) {
			lower++
			continue
		}

		if i == 0 {
			missing = []int{lower, lower + x}
		} else {
			missing = []int{lower + 1, lower + x}
		}

		result = append(result, missing)
		lower = numbers[i]
	}

	//upper is still not reached
	if lower < upper {
		result = append(result, []int{lower + 1, upper})
	}

	return result
}

func main() {
	testCases := []struct {
		numbers []int
		lower   int
		upper   int
		expect  [][]int
	}{
		{[]int{1, 3, 5, 10, 12, 15}, 0, 20, [][]int{{0, 0}, {2, 2}, {4, 4}, {6, 9}, {11, 11}, {13, 14}, {16, 20}}},
		{[]int{3, 5, 8, 10}, 1, 15, [][]int{{1, 2}, {4, 4}, {6, 7}, {9, 9}, {11, 15}}},
		{[]int{}, 5, 10, [][]int{{5, 10}}},
		{[]int{5, 6, 7, 8, 9, 10}, 5, 10, [][]int{}},
		{[]int{1, 2, 3}, 5, 10, [][]int{{5, 10}}},
		{[]int{7, 8, 9}, 5, 10, [][]int{{5, 6}, {10, 10}}},
		{[]int{1, 2, 3}, 10, 5, [][]int{}},
	}

	for i, tc := range testCases {
		result := findGaps(tc.numbers, tc.lower, tc.upper)
		if !reflect.DeepEqual(result, tc.expect) {
			fmt.Printf("Test Case %d FAILED:\n", i+1)
			fmt.Printf("  Input: numbers=%v, lower=%d, upper=%d\n", tc.numbers, tc.lower, tc.upper)
			fmt.Printf("  Expected: %v\n", tc.expect)
			fmt.Printf("  Got:      %v\n", result)
		} else {
			fmt.Printf("Test Case %d PASSED\n", i+1)
		}
	}
}
