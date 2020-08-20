package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	var m map[rune]bool = make(map[rune]bool)
	var sb strings.Builder
	var l string

	if len(s) <= 1 {
		return len(s)
	}

	for i, v := range s {
		sb.WriteRune(v)
		si := s[i+1:]
		for _, w := range si {
			if v == w || m[w] {
				break
			}
			sb.WriteRune(w)
			m[w] = true
		}
		if sbs := sb.String(); len(l) < len(sbs) {
			l = sbs
		}
		// reset
		m = make(map[rune]bool)
		sb.Reset()
	}
	return len(l)
}
