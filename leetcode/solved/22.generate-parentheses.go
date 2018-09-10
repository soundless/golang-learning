/*
 * [22] Generate Parentheses
 *
 * https://leetcode.com/problems/generate-parentheses
 *
 * Medium (43.91%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '3'
 *
 *

package gogo

 * Given n pairs of parentheses, write a function to generate all combinations
 * of well-formed parentheses.
 *
 *
 *
 * For example, given n = 3, a solution set is:
 *
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
*/

package gogo

func generateParentheses(n int) []string {
	result := make([]string, 0)
	dfs(&result, "", n, n)
	return result
}

func dfs(result *[]string, temp string, left int, right int) {
	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*result = append(*result, temp)
		return
	}

	if left > 0 {
		dfs(result, temp+"(", left-1, right)
	}

	if right > 0 {
		dfs(result, temp+")", left, right-1)
	}
}
