package main

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "even number of nodes",
			args: args{
				l1: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
				l2: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
			},
			want: &ListNode{4, &ListNode{4, &ListNode{4, nil}}},
		},
		{
			name: "odd number of nodes",
			args: args{
				l1: &ListNode{1, &ListNode{2, nil}},
				l2: &ListNode{3, &ListNode{2, &ListNode{1, nil}}},
			},
			want: &ListNode{4, &ListNode{4, &ListNode{1, nil}}},
		},
		{	
			name: "extra carry node",
			args: args{
				l1: &ListNode{1, &ListNode{2, &ListNode{9, nil}}},
				l2: &ListNode{2, &ListNode{3, &ListNode{4, nil}}},
			},
			want: &ListNode{3, &ListNode{5, &ListNode{3, &ListNode{1, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !equal(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(got, want *ListNode) bool {
	ln1, ln2 := got, want
	for ln1 != nil && ln2 != nil {
		if ln1.Val != ln2.Val {
			return false
		}
		ln1, ln2 = ln1.Next, ln2.Next
	}
	if ln1 != nil || ln2 != nil {
		return false
	}
	return true
}
