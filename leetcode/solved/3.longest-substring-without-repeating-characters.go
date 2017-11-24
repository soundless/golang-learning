/*
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode.com/problems/longest-substring-without-repeating-characters
 *
 * Medium (24.14%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '"abcabcbb"'
 *
 * Given a string, find the length of the longest substring without repeating
 * characters.
 *
 * Examples:
 *
 * Given "abcabcbb", the answer is "abc", which the length is 3.
 *
 * Given "bbbbb", the answer is "b", with the length of 1.
 *
 * Given "pwwkew", the answer is "wke", with the length of 3. Note that the
 * answer must be a substring, "pwke" is a subsequence and not a substring.
 */

package gogo

func lengthOfLongestSubstring(s string) int {
	var i, max int = 0, 0
	hash := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		hash[s[j]]++
		for hash[s[j]] == 2 && i < j {
			hash[s[i]]--
			i++
		}

		if max < j-i+1 {
			max = j - i + 1
		}
	}
	return max
}
