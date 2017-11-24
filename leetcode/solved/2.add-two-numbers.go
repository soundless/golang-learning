/*
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers
 *
 * Medium (27.37%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order and each of their nodes
 * contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 *
 *
 * Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 * Output: 7 -> 0 -> 8
 */
/**
 * Definition for singly-linked list.
 */

package gogo

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode = new(ListNode)
	dummy := head
	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		tmp := v1 + v2 + carry
		carry = tmp / 10
		var new_node *ListNode = new(ListNode)
		new_node.Val = tmp % 10
		head.Next = new_node
		head = new_node
	}
	return dummy.Next
}
