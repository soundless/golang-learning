/*
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring
 *
 * Medium (25.11%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 *
 * Example:
 *
 * Input: "babad"
 *
 * Output: "bab"
 *
 * Note: "aba" is also a valid answer.
 *
 *
 *
 * Example:
 *
 * Input: "cbbd"
 *
 * Output: "bb"
 *
 *
 */

package gogo

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	left, max := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= max/2 {
			break
		}

		b, e := i, i
		for e < len(s)-1 && s[e] == s[e+1] {
			e++
		}

		i = e + 1

		for e < len(s)-1 && b > 0 && s[b-1] == s[e+1] {
			b--
			e++
		}

		tmp := e - b + 1
		if tmp > max {
			max = tmp
			left = b
		}
	}

	return s[left : left+max]
}
