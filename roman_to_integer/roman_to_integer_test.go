package main

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		// TODO: Add test cases.
		{
			arg:  "III",
			want: 3,
		},
		{
			arg:  "IV",
			want: 4,
		},
		{
			arg:  "IX",
			want: 9,
		},
		{
			arg:  "LVIII",
			want: 58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := romanToInt(tt.arg); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
