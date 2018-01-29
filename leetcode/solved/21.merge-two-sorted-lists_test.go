package gogo

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func Test_21_mergeTwoLists(t *testing.T) {
	var l1 *ListNode = &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	var l2 *ListNode = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	var expected *ListNode = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	actual := mergeTwoLists(l1, l2)
	if cmp.Equal(expected, actual) != true {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
