package gogo

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	var l1 *ListNode = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{3, nil}}}
	var l2 *ListNode = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{4, nil}}}

	var result *ListNode = addTwoNumbers(l1, l2)
	if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
		t.Errorf("Failed: %v -> %v -> %v", result.Val, result.Next.Val, result.Next.Next.Val)
	}
}
