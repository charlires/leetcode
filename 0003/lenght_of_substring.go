package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]bool)
	var substr string
	var sb strings.Builder
	for _, r := range s {
		if m[r] {
			temp := sb.String()
			if len(temp) > len(substr) {
				substr = temp
			}
			m = make(map[rune]bool)
			sb.Reset()
		}
		m[r] = true
		sb.WriteRune(r)
	}
	temp := sb.String()
	if len(temp) > len(substr) {
		substr = temp
	}
	return len(substr)
}
