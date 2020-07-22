package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			args: args{
				nums:   []int{-3, 4, 3, 5, 11},
				target: 0,
			},
			want: []int{0, 2},
		},
		{
			args: args{
				nums:   []int{2, 7, 11, 3, 4},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{0, 7, 4, 3, 0},
				target: 0,
			},
			want: []int{0, 4},
		},
		{
			args: args{
				nums:   []int{8, 7, 4, 3, 0},
				target: 5,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums=%+v, target=%d", tt.args.nums, tt.args.target), func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
