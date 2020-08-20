package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := make([]int, 0)

	for x > 0 {
		a = append(a, x%10)
		x /= 10
	}

	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		if a[i] != a[j] {
			return false
		}
	}
	return true
}
