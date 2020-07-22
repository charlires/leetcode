package main

import (
	"math"
)

func reverse(x int) int {
	var minus bool
	if minus = x < 0; minus {
		x = -x
	}

	a := make([]int, 0)
	for x > 0 {
		a = append(a, x%10)
		x /= 10
	}

	var result int
	l := len(a)
	for i, v := range a {
		result += v * int(math.Pow10(l-(1+i)))
	}

	if minus {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
