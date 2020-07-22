package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{
		// TODO: Add test cases.
		{
			arg:  123,
			want: 321,
		},
		{
			arg:  -123,
			want: -321,
		},
		{
			arg:  120,
			want: 21,
		},
		{
			arg:  1534236469,
			want: 0,
		},
		{
			arg:  -1534236469,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("reverse(%d), want %v", tt.arg, tt.want), func(t *testing.T) {
			if got := reverse(tt.arg); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
