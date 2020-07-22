package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	var tests = []struct {
		arg  int
		want bool
	}{
		// TODO: Add test cases.
		{
			arg:  101,
			want: true,
		},
		{
			arg:  10,
			want: false,
		},
		{
			arg:  111,
			want: true,
		},
		{
			arg:  -111,
			want: false,
		},
		{
			arg:  200,
			want: false,
		},
		{
			arg:  32423,
			want: true,
		},
		{
			arg:  31423,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.arg), func(t *testing.T) {
			if got := isPalindrome(tt.arg); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
