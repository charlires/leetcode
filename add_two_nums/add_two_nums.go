package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int // 0
	var val int   // 0
	result := &ListNode{}
	current := result
	for ; l1 != nil && l2 != nil; l1, l2, current = l1.Next, l2.Next, current.Next {
		val = l1.Val + l2.Val + current.Val
		carry = val / 10
		current.Val = val % 10
		if l1.Next != nil || l2.Next != nil || carry > 0 {
			current.Next = &ListNode{Val: carry}
			switch {
			case l1.Next == nil:
				l1.Next = &ListNode{}
			case l2.Next == nil:
				l2.Next = &ListNode{}
			}
		}
	}
	return result
}
