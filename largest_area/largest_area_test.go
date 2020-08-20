package main

import (
	"fmt"
	"testing"
)

func Test_largestArea(t *testing.T) {
	var tests = []struct {
		samples [][]int32
		want    int32
	}{
		{[][]int32{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0},
		{[][]int32{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 1},
		{[][]int32{{0, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1},
		{[][]int32{{1, 1, 0}, {1, 1, 0}, {0, 0, 0}}, 2},
		{[][]int32{{0, 0, 0}, {0, 1, 1}, {0, 1, 1}}, 2},
		{[][]int32{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d", tt.samples, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := largestArea(tt.samples)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
