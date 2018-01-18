/*
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses
 *
 * Easy (33.11%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '"["'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * The brackets must close in the correct order, "()" and "()[]{}" are all
 * valid but "(]" and "([)]" are not.
 *
 */

package gogo

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			size := len(stack)
			curr := stack[size-1]
			if (curr == '(' && s[i] != ')') ||
				(curr == '[' && s[i] != ']') ||
				(curr == '{' && s[i] != '}') {
				return false
			} else {
				stack = stack[:size-1]
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
