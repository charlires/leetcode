package main

func romanToInt(s string) int {
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int
	var last int

	for i, r := range s {
		num := m[r]
		if last < num && last != 0 {
			res += -last * 2
		}
		res += num
		last = num
	}
	return res
}
