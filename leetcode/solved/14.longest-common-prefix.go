/*
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix
 *
 * Easy (31.30%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '[]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 */
package gogo

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	longest := strs[0]
	for i := 0; i < len(longest); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) > len(longest) || strs[j][i] != longest[i] {
				return longest[:i]
			}
		}
	}

	return longest
}
