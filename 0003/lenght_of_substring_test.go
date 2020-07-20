package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		// TODO: Add test cases.
		{
			arg:  "abcabcbb",
			want: 3,
		},
		{
			arg:  "bbbbb",
			want: 1,
		},
		{
			arg:  "pwwkew",
			want: 3,
		},
		{
			arg:  " ",
			want: 1,
		},
		{
			arg:  "aab",
			want: 2,
		},
		// {
		// 	arg:  "dvdf",
		// 	want: 3,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.arg); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
