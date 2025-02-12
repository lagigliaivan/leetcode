package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	var input = []int{0, 3, -1, 5, 12, 9, 18, 15, 24, 21, 30, 27}
	var expe = []int{-1, 0, 3, 5, 9, 12, 15, 18, 21, 24, 27, 30}

	result := bubbleSort(input)

	assert.Equal(t, expe, result)
}

func BenchmarkStringBuilderConcatenation(b *testing.B) {
	var input = []int{0, 3, -1, 5, 12, 9, 18, 15, 24, 21, 30, 27}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			bubbleSort(input)
		}
	}
}

//enchmarkStringBuilderConcatenation-8   	   17343	     68561 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	example.com/m	2.395s

// BenchmarkStringBuilderConcatenation-8   	   22484	     51017 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	example.com/m	1.953s
