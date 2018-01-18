package gogo

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_19_removeNthFromEnd(t *testing.T) {
	var head *ListNode = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{5, nil}}}}}
	var expected *ListNode = &ListNode{Val: 1, Next: &ListNode{Val: 2,
		Next: &ListNode{Val: 3, Next: &ListNode{5, nil}}}}
	actual := removeNthFromEnd(head, 2)
	if cmp.Equal(expected, actual) != true {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
