package main

import (
	"fmt"
)

func main() {
	nums := []int{0}
	j := canJump(nums)

	fmt.Printf("Can jump: %v\n", j)
}

var alreadyEvaluated map[int]bool

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	alreadyEvaluated = make(map[int]bool)

	r := buildTreeNode(0, nums, len(nums)-1, 0)

	return r
}

func buildTreeNode(index int, nums []int, lastIndex int, totalJumps int) bool {
	_, ok := alreadyEvaluated[index]
	if ok {
		return false
	}

	alreadyEvaluated[index] = true

	if totalJumps >= lastIndex || index >= lastIndex {
		return true
	}

	jumps := nums[index]
	if jumps == 0 {
		return false
	}

	if jumps >= lastIndex {
		return true
	}

	for j := jumps; j >= 1; j-- {
		exit := buildTreeNode(index+j, nums, lastIndex, totalJumps+j)
		if exit {
			return exit
		}
	}

	return false
}
