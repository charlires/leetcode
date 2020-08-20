package main

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		// TODO: Add test cases.
		{
			arg:  "bbbbbb",
			want: 1,
		},
		{
			arg:  "bvbacdb",
			want: 5,
		},
		{
			arg:  "abcabcbb",
			want: 3,
		},
		{
			arg:  " ",
			want: 1,
		},
		{
			arg:  "au",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s=%d", tt.arg, tt.want), func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.arg); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
