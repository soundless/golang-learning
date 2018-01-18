/*
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list
 *
 * Medium (33.03%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '[1]\n1'
 *
 * Given a linked list, remove the nth node from the end of list and return its
 * head.
 *
 *
 * For example,
 *
 *
 * ⁠  Given linked list: 1->2->3->4->5, and n = 2.
 *
 * ⁠  After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 *
 * Note:
 * Given n will always be valid.
 * Try to do this in one pass.
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package gogo

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
