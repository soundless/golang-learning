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

/* func longestPalindrome(s string) string {
 *	if len(s) < 2 {
 *		return s
 *	}
 *
 *	left, max := 0, 1
 *	for i := 0; i < len(s); {
 *		if len(s)-i <= max/2 {
 *			break
 *		}
 *
 *		b, e := i, i
 *		for e < len(s)-1 && s[e] == s[e+1] {
 *			e++
 *		}
 *
 *		i = e + 1
 *
 *		for e < len(s)-1 && b > 0 && s[b-1] == s[e+1] {
 *			b--
 *			e++
 *		}
 *
 *		tmp := e - b + 1
 *		if tmp > max {
 *			max = tmp
 *			left = b
 *		}
 *	}
 *
 *	return s[left : left+max]
 *} 
 */

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    max := 0
    low := 0
    for i := 0; i < len(s) - 1; i++ {
        l1, r1 := expandAroundCenter(s, i, i)
        l2, r2 := expandAroundCenter(s, i, i + 1)
        if max < r1 - l1 - 1 {
            low = l1 + 1
            max = r1 - l1 - 1
        }
        if max < r2 - l2 - 1 {
            low = l2 + 1
            max = r2 - l2 - 1
        }
    }
    return s[low:low+max]
}

func expandAroundCenter(s string, left int, right int) (int, int) {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return left, right
}

